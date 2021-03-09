package loader

import (
	"context"
	"errors"
	"log"
	"os"
	"path"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LoadPersonnel ...
 func LoadPersonnel(fname string) error {
	var personnel []jsontypes.Person

	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// read the personnel json data
	personnel, err = jsontypes.LoadPersonnel(fname)
	if err != nil {
		return err
	}

	// interface slice to contain bson marshalled personnel records
	dox := make([]interface{},0)

	// marshall all the documents into 'dox'
	for _,v := range personnel {
		// marshall the element into a byte slice
		b , err := bson.Marshal(v)
		if err != nil {
			return err
		}
		// append to array of documents
		dox = append(dox,b)
	}

	// get the personnel collection
	coll := client.Database("retrosheet").Collection("personnel")

	// insert the documents
	res, err  := coll.InsertMany(ctx,dox)

	// insertion failed
	if err != nil {
		return err
	}

	// some other fail, auth error can do this
	if res == nil {
		return errors.New("no result : authorization?")
	}
		
	defer client.Disconnect(ctx)

	return nil
 }

 // LoadTeams loads the teams.json file
 func LoadTeams(fname string) error {
    var teams []jsontypes.Team

	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// read the team json data
	teams, err = jsontypes.LoadTeams(fname)
	if err != nil {
		return err
	}

	// byte slice to contain bson marshalled personnel records
	dox := make([]interface{},0)

	// marshall all the documents into 'dox'
	for _,v := range teams {
		// marshall the element into a byte slice
		b , err := bson.Marshal(v)
		if err != nil {
			return err
		}
		// append to array of documents
		dox = append(dox,b)
	}

	// get the personnel collection
	coll := client.Database("retrosheet").Collection("teams")

	// insert the documents
	res, err  := coll.InsertMany(ctx,dox)

	// insertion failed
	if err != nil {
		return err
	}

	// some other fail, auth error can do this
	if res == nil {
		return errors.New("no result : authorization?")
	}

	defer client.Disconnect(ctx)

	return nil
 }

// LoadGameLog loads a single gamelog file
// gamelog files are organized by year
// there are many gamelog files each containing multiple games
func LoadGameLog(fname string) error {
	var games []jsontypes.Game

	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

    // get the game data from the json file
	games, err = jsontypes.LoadGamelog(fname)
	if err != nil {
		return err
	}

	// byte slice to contain bson marshalled personnel records
	dox := make([]interface{},0)

	// marshall all the documents into 'dox'
	for _,v := range games {
		// marshall the element into a byte slice
		b , err := bson.Marshal(v)
		if err != nil {
			return err
		}
		// append to array of documents
		dox = append(dox,b)
	}

	// get the personnel collection
	coll := client.Database("retrosheet").Collection("games")

	// insert the documents
	res, err  := coll.InsertMany(ctx,dox)

	// insertion failed
	if err != nil {
		return err
	}

	// some other fail, auth error can do this
	if res == nil {
		return errors.New("no result : authorization?")
	}


	defer client.Disconnect(ctx)

	return nil
 }

// LoadGames loads ALL gamelog files from the specifried directory
// gamelog files are organized by year
// there are many gamelog files each containing multiple games
// Note : as of the date of this file, there are over 200,000 
// games in all the gamelogs, so this function can take several seconds
// to execute
func LoadGames(dirname string) error {
	var games []jsontypes.Game
	var gamelog []jsontypes.Game

	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create a slice of ALL games
	// get the game data from all the files in the specified directory
	// iterate over the directory containing the game files
	games = make([]jsontypes.Game,0)
	dirs, err := os.ReadDir(dirname)
	for _,dir := range dirs {
		fname := dir.Name()
		// ignore non-json files
		if path.Ext(fname) != ".json" {
			continue
		}
		
		// fs.Direntry Name returns only final element of path
		fname = dirname + "/" + dir.Name()

		// get the game data from the json file
		gamelog, err = jsontypes.LoadGamelog(fname)
		if err != nil {
			return err
		}

		// merge it to the total games
		games = append(games,gamelog...)
	}


	// interface slice to contain bson marshalled personnel records
	dox := make([]interface{},0)

	// marshall all the documents into 'dox'
	for _,v := range games {
		// marshall the element into a byte slice
		b , err := bson.Marshal(v)
		if err != nil {
			return err
		}
		// append to array of documents
		dox = append(dox,b)
	}

	// get the personnel collection
	coll := client.Database("retrosheet").Collection("games")

	// insert the documents
	res, err  := coll.InsertMany(ctx,dox)

	// insertion failed
	if err != nil {
		return err
	}

	// some other fail, auth error can do this
	if res == nil {
		return errors.New("no result : authorization?")
	}
	
	defer client.Disconnect(ctx)

	return nil
}

// DropRetrosheet drop the retrosheet database
// in prep for repopulation
func DropRetrosheet() error {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// drop retrosheet
	err = client.Database("retrosheet").Drop(ctx)
	if err != nil {
		return err
	}
			
	defer client.Disconnect(ctx)

	return nil
}

// PopulateRetrosheet is used to repopulate the complete retrosheet database
// it drops the existing db and reloads personnel, teams and games
// it requires the file and directory names for the data files
func PopulateRetrosheet() error {
	var err error
	var personnel = os.Getenv("RETROSHEET") + "/personnel.json"
	var teams = os.Getenv("RETROSHEET") + "/teams.json"
	var games = os.Getenv("RETROSHEET") + "/games/json/"

	// delete the database
	err = DropRetrosheet()
	if err != nil {
		return err
	}

	// load the personnel  data
	err = LoadPersonnel(personnel)
	if err != nil {
		return err
	}

	// load the team data
	err = LoadTeams(teams)
	if err != nil {
		return err
	}

	// load the game data
	err = LoadGames(games)
	if err != nil {
		return err
	}

	return nil
}
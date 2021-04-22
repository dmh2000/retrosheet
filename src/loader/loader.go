package loader

import (
	"context"
	"errors"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LoadPersonnel ...
 func LoadPersonnel(uri string, fname string) error {
	var personnel []jsontypes.Person

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// read the personnel json data
	personnel, err = jsontypes.ReadPersonnel(fname)
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

	if len(dox) == 0 {
		// no personnel data
		return errors.New("No personnel data found, check path name")
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
 func LoadTeams(uri, fname string) error {
    var teams []jsontypes.Team

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err 
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
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

	if len(dox) == 0 {
		// no team data
		return errors.New("No team data found, check path name")
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
func LoadGameLog(uri string, fname string) error {
	var games []jsontypes.Game

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

    // get the game data from the json file
	games, err = jsontypes.ReadGamelog(fname)
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

	if len(dox) == 0 {
		// no game data
		return errors.New("No game data found, check pathname")
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
func LoadGames(uri, dirname string) error {
	var games []jsontypes.Game

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// create a slice of ALL games
	// get the game data from all the files in the specified directory
	// iterate over the directory containing the game files
	games = make([]jsontypes.Game,0)
	games, err = jsontypes.ReadGames(dirname)
	if err != nil {
		return err
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
	if len(dox) == 0 {
		// no game data
		return errors.New("no game data found, check pathname")
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
func DropRetrosheet(uri string) error {
	
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
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
func PopulateRetrosheet(dirname string, mongodb_uri string) error {
	var err error
	var personnel = dirname + "/personnel.json"
	var teams = dirname + "/teams.json"
	var games = dirname + "/gamelogs/json/"

	// delete the database
	err = DropRetrosheet(mongodb_uri)
	if err != nil {
		return err
	}

	// load the personnel  data
	err = LoadPersonnel(mongodb_uri,personnel)
	if err != nil {
		return err
	}

	// load the team data
	err = LoadTeams(mongodb_uri,teams)
	if err != nil {
		return err
	}

	// load the game data
	err = LoadGames(mongodb_uri, games)
	if err != nil {
		return err
	}

	return nil
}

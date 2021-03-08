package loader

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadall() {}

// wrap 
type mongoOp func(client *mongo.Client) error
 
 // LoadPersonnel ...
 func LoadPersonnel(fname string) error {
	var personnel []jsontypes.Person

	uri := "mongodb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	// return an error until the write to mongodb works
	if len(personnel) > 0 {
		return errors.New("no mongdb yet")
	}
		
	defer client.Disconnect(ctx)

	return nil
 }

 // LoadTeams loads the teams.json file
 func LoadTeams(fname string) error {
    var teams []jsontypes.Team

	uri := "mongodb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	// return an error until the write to mongodb works
	if len(teams) > 0 {
		return errors.New("no mongdb yet")
	}

	defer client.Disconnect(ctx)

	return nil
 }

// LoadGame loads a single gamelog file
// gamelog files are organized by year
// there are many gamelog files each containing multiple games
func LoadGame(fname string) error {
	var games []jsontypes.Game

	uri := "mongodb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	// return an error until the write to mongodb works
	if len(games) > 0 {
		fmt.Println(len(games))
		return errors.New("no mongdb yet")
	}

	defer client.Disconnect(ctx)

	return nil
 }

// LoadGames loads all gamelog files from the specifried directory
// gamelog files are organized by year
// there are many gamelog files each containing multiple games
// Note : as of the date of this file, there are over 200,000 
// games in all the gamelogs, so this function can take several seconds
// to execute
func LoadGames(dirname string) error {
	var games []jsontypes.Game
	var gamelog []jsontypes.Game

	uri := "mongodb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

	// return an error until the write to mongodb works
	if len(games) > 0 {
		fmt.Println(len(games))
		return errors.New("no mongdb yet")
	}
	
	defer client.Disconnect(ctx)

	return nil
}
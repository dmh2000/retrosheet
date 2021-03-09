package loader

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type game struct {
	Home string 
	Hscore int
	Visitor string
	Vscore int
}

func TestConnection(t *testing.T) {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		t.Error(err)
	}

	var g game

	g.Home = "yankees"
	g.Visitor = "padres"
	g.Hscore = 1
	g.Vscore = 0

	b, err := bson.Marshal(&g)

	// create a collection
	coll := client.Database("retrosheet").Collection("test")

	// insert a document
	res, err  := coll.InsertOne(ctx,b)

	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("res == nil")
	}


	defer client.Disconnect(ctx)
}

func TestPersonnel(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/personnel.json"

	err := LoadPersonnel(fname)

	if err != nil {
		t.Error(err)
	}
}


func TestTeams(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/teams.json"

	err := LoadTeams(fname)

	if err != nil {
		t.Error(err)
	}
}

// test loading a single gamelog file into mongodb
func TestGamelog(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/games/json/gl2010.json"

	err := LoadGameLog(fname)

	if err != nil {
		t.Error(err)
	}
}

// test loading ALL gamelogs from a single directory into mongodb
func TestGames(t *testing.T) {
	var dirname = os.Getenv("RETROSHEET") + "/games/json/"

	err := LoadGames(dirname)

	if err != nil {
		t.Error(err)
	}
}

func TestPopulate(t *testing.T) {
	err := PopulateRetrosheet()
	if err != nil {
		t.Error(err)
	}
}
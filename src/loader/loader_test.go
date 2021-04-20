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

var test_mongodb_uri = os.Getenv("RETROSHEET_MONGO")
var test_data_path = os.Getenv("RETROSHEET_DATA")

func TestConnection(t *testing.T) {
	uri := test_mongodb_uri
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
	fname := test_data_path + "/personnel.json"
	uri := test_mongodb_uri

	err := LoadPersonnel(uri, fname)

	if err != nil {
		t.Error(err)
	}
}


func TestTeams(t *testing.T) {
	fname := test_data_path + "/teams.json"
	uri := test_mongodb_uri

	err := LoadTeams(uri,fname)

	if err != nil {
		t.Error(err)
	}
}

// test loading a single gamelog file into mongodb
func TestGamelog(t *testing.T) {
	fname := test_data_path + "/games/json/gl2010.json"
	uri := test_mongodb_uri

	err := LoadGameLog(uri, fname)

	if err != nil {
		t.Error(err)
	}
}

// test loading ALL gamelogs from a single directory into mongodb
func TestGames(t *testing.T) {
	dirname := test_data_path + "/games/json/"
	uri := test_mongodb_uri

	err := LoadGames(uri, dirname)

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
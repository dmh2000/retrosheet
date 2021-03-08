package loader

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func TestConnection(t *testing.T) {
	uri := "mongodb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	cancel()
	if err != nil {
		t.Error(err)
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

	err := LoadGame(fname)

	if err != nil {
		t.Error(err)
	}
}

// test loading all gamelogs from a single directory into mongodb
func TestGames(t *testing.T) {
	var dirname = os.Getenv("RETROSHEET") + "/games/json/"

	err := LoadGames(dirname)

	if err != nil {
		t.Error(err)
	}
}

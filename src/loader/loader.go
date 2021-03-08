package loader

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadall() {}

// wrap 
type mongoOp func(client *mongo.Client) error
 
 // LoadPersonnel ...
 func LoadPersonnel() ([]jsontypes.Person, error) {
	uri := "mondb://dmh2000@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println(ctx)
	
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	return nil,nil
 }

 // LoadTeams ...
 func LoadTeams() ([]jsontypes.Team, error) {
	uri := "mondb://dmh2000@localhost:27017"
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

	defer client.Disconnect(ctx)

	return nil,nil
 }

  // LoadGames ...
  func LoadGames() ([]jsontypes.Game, error) {
	uri := "mondb://dmh2000@localhost:27017"
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

	defer client.Disconnect(ctx)

	return nil,nil
 }
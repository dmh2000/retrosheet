package loader

import (
	"context"
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
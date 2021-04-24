package query

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongodbUri() string {
	return os.Getenv("RETROSHEET_MONGO")
}

func getJsonDataPath() string {
	return os.Getenv("RETROSHEET_DATA")
}

type QueryParams struct {
	uri string				// mongdb server uri
	database string			// selected database
	collection string		// selected collection
	filter interface{}		// bson.D filter
	projection interface{}	// bson.D projection
}

type Decoder interface {
	// decode results after query is executed
	// lets it handle specific types
	decode(cursor *mongo.Cursor,ctx *context.Context) error
}

func QueryDb(qp QueryParams,q Decoder) error {
	var err error
	var client *mongo.Client
	var cursor *mongo.Cursor	
	var opts   *options.FindOptions
	client, err = mongo.NewClient(options.Client().ApplyURI(qp.uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	// get the personnel collection
	coll := client.Database(qp.database).Collection(qp.collection)

	// find documents based on filter and options
	if qp.projection != nil {
		opts = options.Find().SetProjection(qp.projection)
	}
	cursor, err  = coll.Find(ctx, qp.filter,opts)
	defer cursor.Close(ctx)

	// Find failed
	if err != nil {
		return  err
	}
	// some other fail, auth error can do this
	if cursor == nil {
		return errors.New("no result : authorization?")
	}
	
	// decode all matching records
	err = q.decode(cursor,&ctx)
	if err == nil {
		return err
	}
	return nil
}

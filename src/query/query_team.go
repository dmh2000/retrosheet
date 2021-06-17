package query

import (
	"context"
	"errors"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTeamByKey full inline
// this version gets records from the retrosheet:teams database:collection
// based on a single key:value filter. This implementation performs all the
// execution inline with no abstraction of the database access
func GetTeamByKey(uri string, key string, value string) ([]jsontypes.Team, error) {
	var err error
	var client *mongo.Client
	
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	// get the personnel collection
	coll := client.Database("retrosheet").Collection("teams")

	// find documents that have the key "abbr" with the specified value
	cursor, err  := coll.Find(ctx, bson.D{{Key:key, Value:value}} )
	if err != nil {
		return nil,err
	}
	
	defer cursor.Close(ctx)

	// insertion failed
	if err != nil {
		return nil, err
	}

	// some other fail, auth error can do this
	if cursor == nil {
		return nil, errors.New("no result : authorization?")
	}
	
	// decode all matching records
	var teams []jsontypes.Team
	for cursor.Next(ctx) {
		var t jsontypes.Team

		// decode into a team
		err = cursor.Decode(&t)
		if err != nil {
			return nil, err
		}

		// add to list
		teams = append(teams,t)
	}
	return teams, nil
}

// =================================================================
// The following version uses the refactored database query 
// handler QueryDb (in query_util.go)
// this hander requires an instance of the Decoder interface
// and an instance of QueryParams

// TeamQuery : a struct that gets the results of a query
type TeamQuery struct {
	teams []jsontypes.Team
}

// decode implementation of the interface query_util:Query
// given a cursor and a context, this function extracts the 
// team data into the instance of TeamQuery
func (t *TeamQuery) decode(cursor *mongo.Cursor,ctx *context.Context) error {
	var err error
	// decode all matching records
	for cursor.Next(*ctx) {
		var team jsontypes.Team

		// decode into a team
		err = cursor.Decode(&team)
		if err != nil {
			return err
		}

		// add to list
		t.teams = append(t.teams,team)
	}
	return nil
}
	
// QueryTeam performs the same query as GetTeamByKey above,
// however it uses the 'queryDB' function that handles the database 
// connection and access. The only things this function needs to provide
// are the database parameters and an object that 
// implements the query_util:Query interface
func QueryTeam(uri string, database string,filter interface{}, projection interface{}) ([]jsontypes.Team,error) {
	var tq TeamQuery
	var err error
	var params QueryParams = QueryParams{
		uri : uri,
		database: database,
		collection: "teams",
		filter: filter,
		projection: projection,
	}
	err = QueryDb(params,&tq)
	if err != nil {
		return nil,err
	}

	return tq.teams,nil
}
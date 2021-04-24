package query

import (
	"context"

	"github.com/dmh2000/retrosheet/src/jsontypes"
	"go.mongodb.org/mongo-driver/mongo"
)

// =================================================================
// The following version uses the refactored database query
// handler QueryDb (in query_util.go)
// this hander requires an instance of the Decoder interface
// and an instance of QueryParams

// PersonQuery : a struct that gets the results of a query
type GameQuery struct {
	games []jsontypes.Game
}

// decode implementation of the interface query_util:Query
// given a cursor and a context, this function extracts the 
// personnel data into the instance of PersonQuery
func (p *GameQuery) decode(cursor *mongo.Cursor,ctx *context.Context) error {
	var err error
	// decode all matching records
	for cursor.Next(*ctx) {
		var game jsontypes.Game

		// decode into a team
		err = cursor.Decode(&game)
		if err != nil {
			return err
		}

		// add to list
		p.games = append(p.games,game)
	}
	return nil
}
	
// QueryPerson performs the same query as GetPersonByKey above,
// however it uses the 'queryDB' function that handles the database 
// connection and access. The only things this function needs to provide
// are the database parameters and an object that 
// implements the query_util:Query interface
func QueryGames(uri string, database string,filter interface{}, projection interface{}) ([]jsontypes.Game,error) {
	var q GameQuery
	var err error
	var params QueryParams = QueryParams{
		uri : uri,
		database: database,
		collection: "games",
		filter: filter,
		projection: projection,
	}
	err = QueryDb(params,&q)
	if err != nil {
		return nil,err
	}

	return q.games,nil
}
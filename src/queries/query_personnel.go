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
type PersonQuery struct {
	people []jsontypes.Person
}

// decode implementation of the interface query_util:Query
// given a cursor and a context, this function extracts the 
// personnel data into the instance of PersonQuery
func (p *PersonQuery) decode(cursor *mongo.Cursor,ctx *context.Context) error {
	var err error
	// decode all matching records
	for cursor.Next(*ctx) {
		var person jsontypes.Person

		// decode into a team
		err = cursor.Decode(&person)
		if err != nil {
			return err
		}

		// add to list
		p.people = append(p.people,person)
	}
	return nil
}
	
// QueryPerson performs the same query as GetPersonByKey above,
// however it uses the 'queryDB' function that handles the database 
// connection and access. The only things this function needs to provide
// are the database parameters and an object that 
// implements the query_util:Query interface
func QueryPersonnel(uri string, database string,filter interface{}, projection interface{}) ([]jsontypes.Person,error) {
	var tq PersonQuery
	var err error
	var params QueryParams = QueryParams{
		uri : uri,
		database: database,
		collection: "personnel",
		filter: filter,
		projection: projection,
	}
	err = QueryDb(params,&tq)
	if err != nil {
		return nil,err
	}

	return tq.people,nil
}
package api

import graphql "github.com/graph-gophers/graphql-go"

/**
ROOT RESOLVER
*/
type rootResolver struct{}

const schemaDef = `
schema {
	query:Query
}

type Query {
	teamByAbbr(abbr:String!):Team
	teamByName(name:String!):Team
}

type Team {
	abbr: String!
	league: String!,
	city: String!,
	nickname: String!,
	firstYear: String!,
	lastYear: String!
}
`

// return the parsed schema
func RootSchema() *graphql.Schema {
	return graphql.MustParseSchema(schemaDef, &rootResolver{})
}

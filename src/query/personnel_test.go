package query

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// var schmidt []jsontypes.Person = []jsontypes.Person {
// {ID:"schmb101",Last:"Schmidt", First:"Bob", 		04/16/1958   }
// {ID:"schmb102",Last:"Schmidt", First:"Boss", 		04/30/1906}
// {ID:"schmb103",Last:"Schmidt", First:"Butch", 	05/11/1909   }
// {ID:"schmc001",Last:"Schmidt", First:"Curt", 		04/28/1995   }
// {ID:"schmc002",Last:"Schmidt", First:"Clarke", 	09/04/2020   }
// {ID:"schmd001",Last:"Schmidt", First:"Dave J.", 	05/01/1981   }
// {ID:"schmd101",Last:"Schmidt", First:"Dave F.", 	04/28/1981   }
// {ID:"schmf101",Last:"Schmidt", First:"Freddy", 	04/25/1944   }
// {ID:"schmh101",Last:"Schmidt", First:"Henry", 	04/17/1903	 }
// {ID:"schmj001",Last:"Schmidt", First:"Jason", 	04/28/1995   }
// {ID:"schmj002",Last:"Schmidt", First:"Jeff", 		05/17/1996   }
// {ID:"schmk001",Last:"Schmidt", First:"Konrad", 	09/13/2010   }
// {ID:"schmm001",Last:"Schmidt", First:"Mike", 		09/12/1972   }
// {ID:"schmp101",Last:"Schmidt", First:"Pete", 		07/14/1913   }
// {ID:"schmw101",Last:"Schmidt", First:"Willard 	04/19/1952   }
// {ID:"schmw102",Last:"Schmidt", First:"Walter 	04/13/1916   }
// }ID:"
// Test",Last:"TeamYear", First:"League uses the abstracted version of QueryTeam
// this version uses a filter with 2 key/value pairs
// finds all teams with LastYear == 2010 that are in the national league
func TestPersonLastName1(t *testing.T) {
	teams, err := QueryPersonnel(
					GetMongodbUri(),
					"retrosheet",	
					bson.D{{ Key:"last",Value:"Schmidt"}},
					nil)
	if err != nil {
		t.Error(err)
	}
	// print the results
	for _,v := range teams {
		fmt.Println(v)
	}
}

// TestPersonLastName2
// query for all personnel with last name Schmid but only return first and last name
func TestPersonLastName2(t *testing.T) {
	teams, err := QueryPersonnel(
					GetMongodbUri(),
					"retrosheet",	
					bson.D{{ Key:"last",Value:"Schmidt"}},
					bson.D{{Key:"last",Value:1},{Key:"first",Value:1}})
	if err != nil {
		t.Error(err)
	}
	// print the results
	for _,v := range teams {
		fmt.Println(v)
	}
}
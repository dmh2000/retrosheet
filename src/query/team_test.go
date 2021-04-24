package query

import (
	"fmt"
	"testing"

	"github.com/dmh2000/retrosheet/src/jsontypes"

	"go.mongodb.org/mongo-driver/bson"
)

// ================================
// this test covers 'query_team.go'
// ================================

// TestTeamByAbbr
// this version uses the full inline function GetTeamByKey
func TestTeamByAbbr(t *testing.T) {
	teams,err := GetTeamByKey(GetMongodbUri(),"abbr","ATL")
	if err != nil {
		t.Error(err)
	}
	if len(teams) == 0 {
		t.Error("no record found")
	}
	if len(teams) > 1 {
		t.Error("multiple records found", teams)
	}

	// should only be 1 entry
	var v jsontypes.Team = teams[0]
	var u jsontypes.Team

	u = jsontypes.Team{Abbr:"ATL",League:"NL",City:"Atlanta",Nickname:"Braves",  FirstYear:"1966",LastYear:"2010"}		

	if v != u {
		t.Error("record mismatch",v,u)
	}

	fmt.Println(v)
}

// expected data from team by city test
var stlouis []jsontypes.Team = []jsontypes.Team {
	{Abbr:"SL1",League:"NA",City:"St. Louis",Nickname:"Red Stockings",  FirstYear:"1875",LastYear:"1875"},
	{Abbr:"SL2",League:"NA",City:"St. Louis",Nickname:"Brown Stockings",FirstYear:"1875",LastYear:"1875"},
	{Abbr:"SL3",League:"NL",City:"St. Louis",Nickname:"Brown Stockings",FirstYear:"1876",LastYear:"1877"},
	{Abbr:"SL4",League:"AA",City:"St. Louis",Nickname:"Cardinals",      FirstYear:"1882",LastYear:"1891"},
	{Abbr:"SLU",League:"UA",City:"St. Louis",Nickname:"Maroons",        FirstYear:"1884",LastYear:"1884"},
	{Abbr:"SL5",League:"NL",City:"St. Louis",Nickname:"Maroons",        FirstYear:"1885",LastYear:"1886"},
	{Abbr:"SLN",League:"NL",City:"St. Louis",Nickname:"Cardinals",      FirstYear:"1892",LastYear:"2010"},
}

// TestTeamByCity1 uses the abstracted version of QueryTeam
// see 
// 		query_util.go : Decoder and QueryDb
//      query_team.go : QueryTeam
func TestTeamByCity1(t *testing.T) {
	teams, err := QueryTeam(
					GetMongodbUri(),
					"retrosheet",	
					bson.D{{ Key:"city",Value:"St. Louis"}},
				nil)
	if err != nil {
		t.Error(err)
	}
	// print the results
	for i,v := range teams {
		if v != stlouis[i] {
			t.Error(v, " should be ", stlouis[i])
		}
		fmt.Println(v)
	}
}

// TestTeamByCity2 uses the full inline version to query the database
func TestTeamByCity2(t *testing.T) {
	teams,err := GetTeamByKey(GetMongodbUri(),"city", "St. Louis")
	if err != nil {
		t.Error(err)
	}
	if len(teams) == 0 {
		t.Error("no record found")
	}

	if len(teams) != len(stlouis) {
		t.Error("wrong number of teams, should be 7")
	}
	for i,v := range teams {
		if v != stlouis[i] {
			t.Error(v, " should be ", stlouis[i])
		}
		fmt.Println(v)
	}
}

func TestTeamByLastYear(t *testing.T) {
	teams,err := GetTeamByKey(GetMongodbUri(),"lastyear", "2010")
	if err != nil {
		t.Error(err)
	}
	if len(teams) == 0 {
		t.Error("no record found")
	}

	if len(teams) != 30 {
		t.Error(len(teams), " should be ", 30)
	}
	count := 0
	for _,v := range teams {
		fmt.Println(v)
		count++
	}
	fmt.Println(count)
}

var year_league []jsontypes.Team = []jsontypes.Team {
	{Abbr:"CHN",League:"NL",City:"Chicago"			,Nickname: "Cubs" 		  ,FirstYear: "1876",LastYear: "2010"},
	{Abbr:"PHI",League:"NL",City:"Philadelphia"     ,Nickname: "Phillies"	  ,FirstYear: "1883",LastYear: "2010"},
	{Abbr:"PIT",League:"NL",City:"Pittsburgh"       ,Nickname: "Pirates"	  ,FirstYear: "1887",LastYear: "2010"},
	{Abbr:"CIN",League:"NL",City:"Cincinnati"       ,Nickname: "Reds" 		  ,FirstYear: "1890",LastYear: "2010"},
	{Abbr:"SLN",League:"NL",City:"St. Louis"        ,Nickname: "Cardinals" 	  ,FirstYear: "1892",LastYear: "2010"},
	{Abbr:"LAN",League:"NL",City:"Los Angeles"      ,Nickname: "Dodgers" 	  ,FirstYear: "1958",LastYear: "2010"},
	{Abbr:"SFN",League:"NL",City:"San Francisco"    ,Nickname: "Giants" 	  ,FirstYear: "1958",LastYear: "2010"},
	{Abbr:"HOU",League:"NL",City:"Houston"          ,Nickname: "Colts" 		  ,FirstYear: "1962",LastYear: "2010"},
	{Abbr:"NYN",League:"NL",City:"New York"         ,Nickname: "Mets" 		  ,FirstYear: "1962",LastYear: "2010"},
	{Abbr:"ATL",League:"NL",City:"Atlanta"          ,Nickname: "Braves" 	  ,FirstYear: "1966",LastYear: "2010"},
	{Abbr:"SDN",League:"NL",City:"San Diego"        ,Nickname: "Padres" 	  ,FirstYear: "1969",LastYear: "2010"},
	{Abbr:"FLO",League:"NL",City:"Florida"          ,Nickname: "Marlins" 	  ,FirstYear: "1993",LastYear: "2010"},
	{Abbr:"COL",League:"NL",City:"Colorado"         ,Nickname: "Rockies" 	  ,FirstYear: "1993",LastYear: "2010"},
	{Abbr:"ARI",League:"NL",City:"Arizona"          ,Nickname: "Diamondbacks" ,FirstYear: "1998",LastYear: "2010"},
	{Abbr:"MIL",League:"NL",City:"Milwaukee"        ,Nickname: "Brewers" 	  ,FirstYear: "1998",LastYear: "2010"},
	{Abbr:"WAS",League:"NL",City:"Washington"       ,Nickname: "Nationals" 	  ,FirstYear: "2005",LastYear: "2010"},
}

// TestTeamYearLeague uses the abstracted version of QueryTeam
// this version uses a filter with 2 key/value pairs
// finds all teams with LastYear == 2010 that are in the national league
func TestTeamYearLeague(t *testing.T) {
	teams, err := QueryTeam(
					GetMongodbUri(),
					"retrosheet",	
					bson.D{{ Key:"lastyear",Value:"2010"}, { Key:"league",Value:"NL"}},
					nil)
	if err != nil {
		t.Error(err)
	}
	// print the results
	for i,v := range teams {
		if v != year_league[i] {
			t.Error(v,"should be",year_league[i])
		}
		fmt.Println(v)
	}
}
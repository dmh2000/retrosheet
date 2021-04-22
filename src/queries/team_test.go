package query

import (
	"fmt"
	"testing"

	"github.com/dmh2000/retrosheet/src/jsontypes"
)

func TestTeamByAbbr(t *testing.T) {
	teams,err := getTeamByKey("mongodb://localhost:27017","abbr","ATL")
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

func TestTeamByCity(t *testing.T) {
	teams,err := getTeamByKey("mongodb://localhost:27017","city", "St. Louis")
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
	teams,err := getTeamByKey("mongodb://localhost:27017","lastyear", "2010")
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
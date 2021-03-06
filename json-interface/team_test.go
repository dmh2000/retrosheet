package json_interface

import (
	"fmt"
	"reflect"
	"testing"
)

// TestTeamFile ...
func TestTeamFile(t *testing.T) {
	var fname = "../../data/teams.json"
    var teams []team

	teams = make([]team,0)

	teams,err := loadTeams(fname)
	if err != nil {
		t.Error("load failed",err)
	}
	if teams == nil {
		t.Error("load failed (2)")
	}

}

func TestTeamStruct(t *testing.T) {
	var fname = "../../data/teams.json"
    var teams []team

	teams = make([]team,0)

	teams,err := loadTeams(fname)

	if err != nil {
		t.Error("load failed",err)
	}
	if teams == nil {
		t.Error("load failed (2)")
	}

	// use reflect to print the struct values
	v := reflect.ValueOf(teams[0])
	u := v.Type()

	for i:=0;i<v.NumField();i++ {
		fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	}
}
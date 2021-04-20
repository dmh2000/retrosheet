package jsontypes

import (
	"os"
	"testing"
)

// TestTeamFile ...
func TestTeamFile(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/teams.json"
    var teams []Team

	teams = make([]Team,0)

	teams,err := LoadTeams(fname)
	if err != nil {
		t.Error("load failed",err)
	}
	if teams == nil {
		t.Error("load failed (2)")
	}

}

func TestTeamStruct(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/teams.json"
    var teams []Team

	teams = make([]Team,0)

	teams,err := LoadTeams(fname)

	if err != nil {
		t.Error("load failed",err)
	}
	if teams == nil {
		t.Error("load failed (2)")
	}

	// // uncomment this to print and verify the fields of this instance are correct
	// // use reflect to print the struct values
	// v := reflect.ValueOf(teams[0])
	// u := v.Type()

	// for i:=0;i<v.NumField();i++ {
	// 	fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	// }
}
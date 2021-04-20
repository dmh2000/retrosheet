package jsontypes

import (
	"fmt"
	"os"
	"testing"
)

var fpath string = os.Getenv("RETROSHEET") + "/gamelogs/json" 
// TestGameFile ...
func TestGameFile(t *testing.T) {
	var fname = fpath + "/gl2013.json"
    var games []Game

	games = make([]Game,0)

	games,err := LoadGamelog(fname)
	if err != nil {
		t.Error("load failed",err)
	}
	if games == nil {
		t.Error("load failed (2)")
	}

}

func TestGameDir(t *testing.T) {
	var dirname = fpath
    var games []Game

	games,err := LoadGames(dirname)
	if err != nil {
		t.Error("load failed",err)
	}
	if games == nil {
		t.Error("load failed (2)")
	}
	if len(games) == 0 {
		t.Error("no games found")
	}

	// print number of games
	fmt.Printf("number of games %v\n",len(games))
}

func TestGameStruct(t *testing.T) {
	var fname = fpath + "/gl2013.json"
    var games []Game

	games = make([]Game,0)

	games,err := LoadGamelog(fname)

	// print number of games
	fmt.Printf("number of games %v\n",len(games))

	if err != nil {
		t.Error("load failed",err)
	}
	if games == nil {
		t.Error("load failed (2)")
	}

	// uncomment this to print and verify the fields of this instance are correct
	// // use reflect to print the struct values
	// v := reflect.ValueOf(games[0])
	// u := v.Type()

	// for i:=0;i<v.NumField();i++ {
	// 	fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	// }
}
package jsontypes

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// TestGameFile ...
func TestGameFile(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/games/json/gl2013.json"
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
	var dirname = os.Getenv("RETROSHEET") + "/games/json/"
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
}

func TestGameStruct(t *testing.T) {
	var fname = os.Getenv("RETROSHEET") + "/games/json/gl2013.json"
    var games []Game

	games = make([]Game,0)

	games,err := LoadGamelog(fname)

	if err != nil {
		t.Error("load failed",err)
	}
	if games == nil {
		t.Error("load failed (2)")
	}

	// use reflect to print the struct values
	v := reflect.ValueOf(games[0])
	u := v.Type()

	for i:=0;i<v.NumField();i++ {
		fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	}
}
package jsontypes

import (
	"fmt"
	"reflect"
	"testing"
)

// TestGameFile ...
func TestGameFile(t *testing.T) {
	var fname = "../../data/games/json/gl2013.json"
    var games []Game

	games = make([]Game,0)

	games,err := LoadGames(fname)
	if err != nil {
		t.Error("load failed",err)
	}
	if games == nil {
		t.Error("load failed (2)")
	}

}

func TestGameStruct(t *testing.T) {
	var fname = "../../data/games/json/gl2013.json"
    var games []Game

	games = make([]Game,0)

	games,err := LoadGames(fname)

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
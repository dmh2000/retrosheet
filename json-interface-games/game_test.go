package main

import (
	"fmt"
	"testing"
)

// TestGameFile ...
func TestGameFile(t *testing.T) {
	var fname = "gl2020.json"
    var games []game

	games = make([]game,0)

	games,err := loadGames(fname)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(games[0])
}
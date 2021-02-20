package main

import "fmt"

func main() {
	// only used to generate the json to go struct name mapping
	// uncomment to use
//	makeGameStruct()

	var fname = "/home/dmh2000/projects/baseball/data/games/json/1871_2020/gl2013.json"
    var games []game

	games = make([]game,0)

	games,err := loadGames(fname)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(games[0])
}

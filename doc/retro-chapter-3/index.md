---
title: "So You Like Go and Baseball - Chapter 3"
date: 2020-05-18
slug: "/retro-chapter-3"
---

# Chapter 3 - Reading the JSON Data

A library for reading the json data files into Go objects.

The general method used in each of Games, Teams and Personnel is to read the json files created in chapter 2 and use json.Unmarshal them into arrays of Go structs. In addition the code for each type contains the relevant struct definition and the associated json annotations. The annotations do two things:

- json:_name_ provides a mapping from the field names in the json file to the names in the Go struct. In this case I tried to keep the names the same but for consistency it is helpful to have this annotation.
- json:'omitempty' is especially important because the raw data has some missing items and this annotation will populate the affected field with the default value for the type.

## Gamelogs

Each gamelog in data/games/json/glXXXX.json contains multiple games. Typically a season but could also be playoffs, depending on the files used.

### game.go

```go
// the type definition for the data associated with a single game
type Game struct

// ReadGameLog reads a single json gamelog file and returns a slice
// containing the data for all the games in that single gamelog
func ReadGamelog(fname string) ([]Game, error)

// ReadGames reads all gamelog files in a single directory and
// then uses LogGameLog to read and store the gamelog data.
// The output is a slice of 'Game' that contains all the game data
// from the specified directory
func ReadGames(dirname string) ([]Game, error)
```

## Personnel

There is a single json file containing data for all the personnel that participated in MLB baseball. This includes players, managers and umpires. The

### personnel.go

```go
// the type definition of a Person object in the MLB context
type Person struct

// ReadPersonnel reads the retrosheet personnel json file and returns a slice of
// Person's
func ReadPersonnel(fname string) ([]Person, error)
```

## Teams

There is a single json file containing data for all the teams that participated in MLB baseball.

### teams.go

```go
// the type definition of a Team object
type Team struct

// ReadTeams reads team data from the specified file and
// returns a slice of Team data or an error
func ReadTeams(fname string) ([]Team, error)
```

## Directory Structure

The _jsontypes_ directory (under _src_) contains the code and tests for each of the three types of data

<pre>
jsontypes
│   ├── game.go            // functions to load game logs
│   ├── game_test.go       // test game.go
│   ├── personnel.go       // functions to load personnel data
│   ├── personnel_test.go  // test personnel.go
│   ├── team.go            // functions to load team data
│   ├── team_test.go       // test team.go
│   ├── test.sh            // shell command to run the tests
│   └── README.md        
</pre>

---
title: "So You Like Go and Baseball - Chapter 3"
date: 2020-05-18
slug: "/retro-chapter-3"
---

# Chapter 3 - Reading the JSON Data (jsontypes)

A library named "jsontypes" for reading the json data files into Go objects.

The general method used in each of Games, Teams and Personnel is to read the json files created in chapter 2 and use json.Unmarshal them into arrays of Go structs. In addition the code for each type contains the relevant struct definition and the associated json annotations. The annotations do two things:

- json:_name_ provides a mapping from the field names in the json file to the names in the Go struct. In this case I tried to keep the names the same but for consistency it is helpful to have this annotation.
- json:'omitempty' is especially important because the raw data has some missing items and this annotation will populate the affected field with the default value for the type.

## Data Location

For convenience and portability the code access the csv and json data using an environment variable.
The file _env.sh_ exports the appropriate strings. You will need to configure it for your system.

```bash
# env.sh
# source this file to update settings

# WARNING
# treat this file as a template
# do not add it to version control if it contains any
# private information such as usernames, passwords, addresses or other authentication data

# set the required environment variables

# base path of RETROSHEET project code
export RETROSHEET="${HOME}/projects/baseball/retrosheet"

# url to mongodb server
# default local instance
export RETROSHEET_MONGO="mongodb://localhost:27017"

# example for mongodb atlas
# export RETROSHEET_MONGO="mongodb+srv://<username>:<password>@cluster0.<cluster id>.mongodb.net/<database name>?retryWrites=true&w=majority"

# based path to retrosheet data files
export RETROSHEET_DATA="${HOME}/projects/baseball/data"
```

## Code

[Code is here](https://github.com/dmh2000/go_baseball_with_retrosheet/tree/main/src/jsontypes)

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

### Gamelogs

Each gamelog in data/gamelogs/json/glXXXX.json contains multiple games. Typically a season but could also be playoffs, depending on the files used.

#### game.go

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

### Personnel

There is a single json file containing data for all the personnel that participated in MLB baseball. This includes players, managers and umpires. The

#### personnel.go

```go
// the type definition of a Person object in the MLB context
type Person struct

// ReadPersonnel reads the retrosheet personnel json file and returns a slice of
// Person's
func ReadPersonnel(fname string) ([]Person, error)
```

### Teams

There is a single json file containing data for all the teams that participated in MLB baseball.

#### teams.go

```go
// the type definition of a Team object
type Team struct

// ReadTeams reads team data from the specified file and
// returns a slice of Team data or an error
func ReadTeams(fname string) ([]Team, error)
```

## Runtime Environment

The code and scripts in this article were all developed and run on Linux. The Node.js/JavaScript and Go code will run on windows without change. However any shell scripts would have to be modified for Windows cmd.exe of Powershell.

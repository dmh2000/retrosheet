- csv-transform : using json csv-parse to read original retrosheet csv files and transform to JSON
- json-interface : convert the csv data spec to a golang interface for use by rtfs
- loader-go : golang code to read the json files

# Contents

## ./doc

## ./csv-transform-games

Code that transforms the gamelog csv files from retrosheet.org into json files
These files contain detailed information about specific games by season and year.
(Game Logs)[https://www.retrosheet.org/gamelogs]

## ./csv-transform-teams

Code that transforms the teams csv file from retrosheet.org into json files
This file contains information about the teams that appear in the gamelogs.
(Teams)[https://www.retrosheet.org/TeamIDs.htm]

## ./csv-transform-ids

Code that transforms the csv file with data about personnel into json files.
(Personnel)[https://www.retrosheet.org/retroID.htm]

## ./json-interface

A package that has functions to unmarshall the json files for teams, personnel and game logs. Also includes test code for each type of data.

- game.go :
  - func loadGames(fname string) ([]game, error)
  - returns a slice of game logs or an error
- personnel.go :
  - func loadIDs(fname string) ([]id, error)
  - returns a slice of personnel data or an error
- team.go :
  - func loadTeams(fname string) ([]team, error)
  - returns a slice of team data or an error

## loader-go

A program that takes the gamelog json files and loads them into a local mongodb database.

## temporary

- misc
- json-interface-games
- json-interface-ids
- loader

# RETROSHEET

The code in this repo my own creation and is not supported or endorsed by
retrosheet.org. The following are the license terms for use of the
retrosheet data.

## RETROSHEET DATA

The information used here was obtained free of
charge from and is copyrighted by Retrosheet. Interested
parties may contact Retrosheet at "www.retrosheet.org".

## RETROSHEET LICENSE

Recipients of Retrosheet data are free to make any desired use of
the information, including (but not limited to) selling it,
giving it away, or producing a commercial product based upon the
data. Retrosheet has one requirement for any such transfer of
data or product development, which is that the following
statement must appear prominently:

     The information used here was obtained free of
     charge from and is copyrighted by Retrosheet.  Interested
     parties may contact Retrosheet at "www.retrosheet.org".

Retrosheet makes no guarantees of accuracy for the information
that is supplied. Much effort is expended to make our website
as correct as possible, but Retrosheet shall not be held
responsible for any consequences arising from the use the
material presented here. All information is subject to corrections
as additional data are received. We are grateful to anyone who
discovers discrepancies and we appreciate learning of the details.

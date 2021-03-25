#!/bin/bash

# exit when any command fails
set -e

# check that script is in the right place
D=$(basename `pwd`)
if [ "$D" != "csv-transform" ]; then
    echo Must be in directory csv-transform
    exit 1
fi

# Game logs
cd csv-transform-games 
echo "------" $(basename `pwd`) "------"
node index.js ~/projects/baseball/data/gamelogs/csv ~/projects/baseball/data/gamelogs/json
cd ..

# Teams
cd csv-transform-teams
echo "------" $(basename `pwd`) "------"
node index.js ~/projects/baseball/data/teamabr.csv >~/projects/baseball/data/teams.json
cd ..

# Personnel
cd csv-transform-personnel
echo "------" $(basename `pwd`) "------"
node index.js ~/projects/baseball/data/retroid.csv >~/projects/baseball/data/personnel.json

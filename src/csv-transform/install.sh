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
npm install
cd ..

# Teams
cd csv-transform-teams
echo "------" $(basename `pwd`) "------"
npm install
cd ..

# Personnel
cd csv-transform-personnel
echo "------" $(basename `pwd`) "------"
npm install

exit
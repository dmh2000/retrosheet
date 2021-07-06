#!/bin/bash

# quit on any error
set -e 

# Set environment variables
echo Be sure the RETROSHEET environment variables are set correctly
echo $RETROSHEET_DATA
echo $RETROSHEET_MONGO
# source ./env.sh

# ================================
# delete existing processed data
# ================================
pushd $RETROSHEET_DATA 
rm -fv *.json
rm -fv gamelogs/json/*.json
popd


pushd src/csv-transform
# ================================
# Install node modules 
# ================================
echo NPM INSTALL
./install.sh 
# ================================
# Transform csv to json
# ================================
echo TRANFORMING CSV FILES
./transform.sh
popd

# ================================
# test jsontype library
# ================================
pushd src/jsontypes
echo TESTING 
go test -v .
popd

# ================================
# populate the database
# ================================
echo Populating mongdb retrosheet databse
echo use "./retrosheet -p to populate the database"
# ./retrosheet -p

# ================================
# test sample queries
# ================================
echo TESTING QUERIES
echo use "./retrosheet -q to test the queries"
# ./retrosheet -q

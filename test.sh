#!/bin/bash

# quit on any error
set -e 

# Set environment variables
source ./env.sh

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

pushd src/jsontypes
# ================================
# test jsontype library
# ================================
echo TESTING 
go test -v .
popd

pushd src/loader
# ================================
# test loader
# ================================
# populates the database
# ================================
echo TESTING LOADER + populating mongdb retrosheet databse
go test -v .
popd

pushd src/queries
# ================================
# test sample queries
# ================================
echo TESTING QUERIES
go test -v .
popd
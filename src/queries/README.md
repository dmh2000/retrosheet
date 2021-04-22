# Queries

This directory contains various sample queries and the associated test code.

Each file, query_team, query_personnel, query_game contain various preconfigured queries meant to be examples that could be expanded upon.

The query tests use Retrosheet data as of 2021/4/21.

IMPORTANT: the keys in Mongodb records are all lower case while the matching field names in the Go structs are Capitalized.

example for Team query by city

- mongodb key is "city"
- jsontypes.Team field name is City

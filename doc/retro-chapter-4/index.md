---
title: "So You Like Go and Baseball - Chapter 4"
date: 2020-05-18
slug: "/retro-chapter-4"
---

# Chapter 4 - Populating the Database (loader)

This step populates a Mongodb database with the games, teams and personnel data. It uses the jsontypes library to access the local data and upload it using the Golang mongodb API.

For this step you need to have an active mongodb instance with access permissions to create and access a database.
It could be local or remote (such as Mongodb Atlas Cloud). Mongodb Atlas is a good choice for quick setup without having to install mongodb locally. It has a free tier that is fine for this application. However, because access is over the network, some of the data uploads will take a lot longer than if you have a local instance.

You can find how to use either approach at [Mongodb.com](mongodb.com).

Note: the Mongodb Atlas free tier has a 512MB limit and depending on which set of games you load it may overflow that limit. When the limit is reached it just stops saving the additional games. If you use the recommended data and scripts in this repo it may overflow so you only get most of the games but the later dates will not be stored. One way to fix this is to remove gamelog json files in the raw data that are in years you aren't concerned with, and repopulate the database. If you have a local instance of mongodb you probably won't encounter this issue.

## env.sh

To make things easy and portable all the tests access the mongodb instance via an environment variable.

To connect to the mongodb instance you will need a connection string or URI. The file _env.sh_ contains examples:

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

## Runtime Environment

The code and scripts in this article were all developed and run on Linux. The Node.js/JavaScript and Go code will run on windows without change. However any shell scripts would have to be modified for Windows cmd.exe of Powershell.

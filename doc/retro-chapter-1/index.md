---
title: "So You Like Go and Baseball"
date: 2020-05-18
slug: "/retro-chapter-1"
---

# Table of Contents

- [Chapter 1 - Overview](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-1/index.md)
- [Chapter 2 - Transform CSV](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-2/index.md)
- [Chapter 3 - Read JSON](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-3/index.md)
- [Chapter 4 - Upload to Mongodb](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-4/index.md)
- [Chapter 5 - Simple Queries](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-5/index.md)
- [Chapter 6 - API Microservice](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-6/index.md)

# So You Like Go and Baseball

## Introduction

### Runtime Environment

The code and scripts in this article were all developed and run on Linux. The Node.js/JavaScript and Go code will run on windows without change. However any shell scripts would have to be modified for Windows cmd.exe of Powershell.

Have you ever wondered what let sportscasters pull up stats
on all sorts of things during a game? Did you ever want
to have something like that for free? And do you prefer
programming in Go?

This series is intended to help you put together a package
that will let me (and you) build a database of baseball data and
access it with simple queries. All the code will be on [github at](https://github.com/dmh2000/retrosheet).

## Chapter 1 - Retrosheet

There are various sources for baseball data that can be accessed online. A good list of the main ones is discussed at [sabr.org](https://sabr.org/how-to/statistical-databases-and-websites). Some of those have data that is already cooked and you perform queries online, but you aren't able to download the actual databases.
Some have API's available. Another source that I have used is [mysportsfeed.com](https://mysportsfeeds.com). It is excellent and provides live game data as well as archives via an API. It isn't free but if you want some live data its worth a look. I wrote a phone app that utilized this service to show game stats, using Dart and Flutter.

What I ended up wanting to do is use one of these sources that let me host the database and provide an API for using it. Not because I want to complete with the online services, but instead as an exercise in using Go to write a service and maybe a web front end for making queries. A learning experience.

One of the sources mentioned at [sabr.org](https://sabr.org/how-to/statistical-databases-and-websites) that fit the bill is (retrosheet.org)[https://www.retrosheet.org/]. It has some pretty comprehensive data about games, personnel and teams. Best of all their data is free to download and use with the proper attribution. So I started this project using their data as the basis. Just to be clear, I am not affiliated with retrosheet.org, I am just using their data and putting a wrapper around it.

Retrosheet is put together by volunteers and its a lot of work. You may find that some of the data is incomplete. I'm sure they would welcome any help updating the data.

I suggest you peruse (retrosheet.org)[https://www.retrosheet.org/] to get a feel for what it provides. The folks there put a lot of work into making this data available. I thank them for that!

<pre>
     The information used here was obtained free of
     charge from and is copyrighted by Retrosheet.  Interested
     parties may contact Retrosheet at "www.retrosheet.org".
</pre>

### The Data

The data from Retrosheet I am using includes a database of all professional games from 1871 to the most recent completed season. That's the big one. It also has a couple of tables for personnel and teams. I am using those three to begin with. They also
have a database of play-by-play data for pro teams from 1920 to 2020. To keep it simple at first I am skipping that one. But the methodology I am using to incorporate the data is applicable to that one too.

Here's the files you should download if you want to follow along.

- [Game Logs](https://www.retrosheet.org/gamelogs/index.html)
  - The gamelogs include scores, player info and other stuff. They are organized by season and can be downloaded piecemeal. I downloaded (the whole thing)[https://www.retrosheet.org/gamelogs/gl1871_2020.zip] but my code doesn't require all the years if you don't need them. Just download the ones you prefer. In a later chapter there is an option to upload the gamelog data to Mongodb Cloud. The free tier of Mongodb has a 512MB limit and this gamelog set exceeds that limit, so you would have to delete about 2/3 of the csv data before proceeding.
  - Each gamelog is a csv file with a row for each game. A [description of the columns is here](https://www.retrosheet.org/gamelogs/glfields.txt). There is a gamelog file for each year.
- [Team Data](https://www.retrosheet.org/TeamIDs.htm)
  - This data has just a few fields but is useful for queries because you can use it as keys for finding data in the game logs. There is only one file .
  - The [Team Data File](https://www.retrosheet.org/TEAMABR.TXT) is here.
- [Personnel Data](https://www.retrosheet.org/retroID.htm)
  - This data has just a few fields but again gives you keys that can be used for lookups in the game logs. There is only one file.
  - This data doesn't have a separate download, you just need to copy the data from the web page itself.

To use the data with the code I put together, you should create a directory and put the team and personnel files in it. Then in that directory make a subdirectory (name it csv?) and unzip the gamelogs there. The code will want the paths to the two individual files plus the path to the directory containing the gamelog csv files.

### The Code

I tried to put together a step by step process that will end up with something usable. I am keeping each step as a separate chapter. And the chapters are short, focusing on one step in the process. Here's the table of contents going forward.

- Chapter 2 - CSV Transform to JSON
  - transform the retrosheet csv files to JSON
  - oops this one uses node.js instead of Go
- Chapter 3 - Reading the JSON Data
  - a Go package that has functions to read the JSON files and returns the data
  - ok from here on out its Go
- Chapter 4 - Populating the Database
  - a Go package that has functions to populate a Mongodb database with collections for team, personnel and game data.
- Chapter 5 - Simple Queries
  - a Go package to execute some simple queries using a cli.
- Chapter 6 and more
  - I promise to complete chapters 1-5. After that I envision creating a web app as a front end for queries to the Mongodb database. That's going to be a lot of work so no promises for 6 on. Wow is coming out with a big patch that might take up a bunch of my time.

#### My Coding Approach

I am a fan of not optimizing my code. Read that again. I am not against optimizing, I just prefer to avoid it unless I need it. I don't try to optimize for speed (prematurely). Nothing in chapters 2-4 have any speed requirement. Chapter 5 (queries) might need a little more focus since its the beginning of the user interface.

That also applies to making my code concise. You'll see that I have some repetition across some functions. I'd rather have each function simple and easy to understand and kind of standalone. At least for a quasi-tutorial like this. I refactor for DRY when I need to. In this project I chose not to. The drawback of this approach is that it can lead to cut-and-paste errors or diverging snippets that should be the same. I ran into those problems a couple of times.

I use TDD with the Go native test framework. Since this isn't production stuff I probably don't test as comprehensively as I could. I find the Go native test framework along with Visual Studio Code provides a way to execute individual tests so conveniently that it makes incremental development extremely easy.

#### The Stack

Here's what I use for my stack:

- Visual Studio Code
  - running on Windows with remote to my Linux box.
- Linux (Mint)
  - all my actual code is resident on the Linux box. All execution is on Linux. Should work on Windows or Mac.
- Chapters 2 - 5
  - Node.js
  - Go
  - Mongodb
- Chapters 6 ...?
  - React
  - Graphql
  - Nginx
  - Docker

## Data Downloads

- make a directory to store the downloaded files
- if you name this directory 'data' and create it adjacent to the cloned retrosheet repo then you can use some scripts later on. See the directory structure below.
- Download from Retrosheet.org
  - [Team Abbreviations](https://www.retrosheet.org/TEAMABR.TXT)
    - [the reference for this file is at ](https://www.retrosheet.org/TeamIDs.htm)
    - this file contains abbreviations and data about teams
    - this file can be downloaded and renamed teamabr.csv without any editing
  - [Retrosheet IDs](https://www.retrosheet.org/retroID.htm)
    - [the reference for this file is at ](https://www.retrosheet.org/retroID.htm)
    - this contains player id, name and dates
    - this file downloads with some text at the beginning that needs to be removed
    - copy the file to retroid.csv, edit it and remove everything down to the first player id
  - [Game Logs](https://www.retrosheet.org/gamelogs/gl1871_2020.zip)
    - [the reference for this file is at](https://www.retrosheet.org/gamelogs/glfields.txt)
    - download and unzip to 'retrosheet'
    - you can download various subsets of game logs
    - download the reference glfields.txt if you want a local copy

This is the directory structure I used. Its not required but there are some scripts set up for this layout.

  <pre>
.
├── data
│   |── glfields.txt
│   |── teamabr.txt
│   |── teamabr.csv
│   |── teamabr.json (converted)
│   |── retroid.csv
│   |── retroid.txt
│   |── retroid.json (converted)
│   └── gamelogs
│       ├── csv (downloaded gamelogs)
│       └── json (converted gamelogs)
└── retrosheet
    ├── doc
    │   ├── retro-chapter-1
    │   ├── retro-chapter-2
    │   ├── retro-chapter-3
    │   ├── retro-chapter-4
    │   └── retro-chapter-5
    └── src
        ├── csv-transform
        │   ├── csv-transform-games
        │   ├── csv-transform-personnel
        │   └── csv-transform-teams
        ├── jsontypes
        ├── loader
        └── queries

  </pre>

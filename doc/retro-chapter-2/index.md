---
title: "So You Like Go and Baseball - Chapter 2"
date: 2020-05-18
slug: "/retro-chapter-2"
---

# Table of Contents

- [Chapter 1 - Overview](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-1/index.md)
- [Chapter 2 - Transform CSV](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-2/index.md)
- [Chapter 3 - Read JSON](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-3/index.md)
- [Chapter 4 - Populate the Database](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-4/index.md)
- [Chapter 5 - Simple Queries](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-5/index.md)
- [Chapter 6 - API Microservice](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-6/index.md)
  
# Chapter 2 - CSV Transform to JSON

The code for this step is in the repo at retrosheet/src/csv-transform. Yes I know I said this was about Go but in this case node/JavaScript was a bit simpler to get working right. I promise the rest of the work is in Go.

You will need to install node.js version 10 or later for this step.
This is the directory layout for retrosheet/csv-transform after the 'npm install' are been done. (see below)

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

[Code is here](https://github.com/dmh2000/retrosheet/tree/main/src/csv-transform)

<pre>
csv-transform
├── csv-transform-games
│   ├── csvFields.js
│   ├── gameLog.js
│   ├── glfields.js
│   ├── index.js
│   ├── LICENSE
│   ├── node_modules
│   ├── package.json
│   ├── package-lock.json
│   └── README.md
├── csv-transform-personnel
│   ├── index.js
│   ├── node_modules
│   ├── package.json
│   ├── package-lock.json
│   └── README.md
├── csv-transform-teams
│   ├── index.js
│   ├── node_modules
│   ├── package.json
│   ├── package-lock.json
│   └── README.md
├── install.sh
└── transform.sh
</pre>

There code in each of the three directories is pretty simple. Each program reads the appropriate input .csv file, uses the node.js package 'csv-parse' to read the csv lines into JavaScript Objects, which are then converted to JSON using JSON.stringify. Then concatenated into a single JSON file and written out to the destination directory.

In each script there is a definition of the column headings used by 'csv-parse'. These column headings become the field names in the output JSON files.

## Scripts

If you used the directory layout from chapter 1 then you can execute the following script to transform the downloaded files to json. Typically you only need to do this process once unless you download new files, such as more gamelogs. If you rerun the commands with the same paths it will overwrite any matching files in the destination.

- If you haven't run the script before and haven't done "npm install" in any of the sub directories, run this one. It does "npm install" before running the scripts
- $ ./install.sh
- If you have run 'install.sh" before, or already did npm install in the sub directories, run this script.
- $ ./transform.sh

## Individual

### Gamelogs

- this produces one file for each baseball season in the downloaded gamelogs. Each file has all the games for that season.
- cd to csv-transform/csv-transform-games
- npm install
- node index.js [directory with retrosheet gamelog .TXT files] [output directory]
- "$ node index.js ~/projects/baseball/data/gamelogs/csv ~/projects/baseball/data/gamelogs/json"

### Teams

- the output file is named 'teams.json'
- cd to csv-transform/csv-transform-teams
- npm install
- node index.js [path to data/teamabr.csv] >[path to data/teamabr.json]
- "$ node index.js ~/projects/baseball/data/teamabr.csv >~/projects/baseball/data/teams.json"

### Personnel

- The output file is named 'personnel.json'
- cd to csv-transform/csv-transform-personnel
- npm install
- node index.js [path to data/retroid.csv] >[path to data/retroid.json]
- "$ node index.js ~/projects/baseball/data/retroid.csv >~/projects/baseball/data/personnel.json"

## Runtime Environment

The code and scripts in this article were all developed and run on Linux. The Node.js/JavaScript and Go code will run on windows without change. However any shell scripts would have to be modified for Windows cmd.exe of Powershell.

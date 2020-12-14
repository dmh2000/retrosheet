# retrosheet data conversion

The node files in this repo can be used to convert retrosheet baseball gamelogs from https://www.retrosheet.org/gamelogs/index.html.

The content of the gamelogs is specified in https://www.retrosheet.org/gamelogs/glfields.txt

The javascript executable files that are currently working are :

1. gameToJson.js : decode one game log file to json
   - execute this one with an input file path and an output filepath to convert a retrosheet gamelog GLyyyy.TXT to glyyyy.json.
   - node gameToJson.js ../data/csv/GL2010.TXT ../data/json/gl2010.json

1. gamesToJson.js : decode all GLxxxx.TXT files in a specified directory.
   - execute this one with an input directory path containing GLxxxx.TXT files and an output directory path to convert all retrosheet gamelog GLyyyy.TXT to glyyyy.json.
   - node gamesToJson \<directory with gamelog .TXT files\> \<output directory\>

3. csvFields.js : execute this one to get a one line output that has all the field labels in comma delimited format. this is useful it you
   want to prepend it to one of the GLyyy.TXT file as its label header so you can read it as a csv file with Excel

- node csvFields.js

4. src/gameLog.js : a library file that supports parsing a gamelog file. used by gameToJson.js and gamesToJson

5. src/glfields.js : a library file that contains the headers for each field in one game line in a GLyyyy.TXT file

The javascript executable files that are not working or are not user friendly are:

1. loader.js : loads a game log into a mongodb database
2. query.js : performs a simple query for gamelog data on a mongodb database
3. collections.js : contains listCollections(url,dbname) which is used to list the available collections in a mongodb database
4. index.js : used to test different executables under development

NOTE : if you use gameToJson.js to read a game file and you get an exception of the form 'line length overrun ...' it means something
in the gamelog input file is not right or is different from the
standard as described in the glfields.txt file from retrosheet.

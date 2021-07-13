# DMH2000's Retrosheet Baseball Data Utilities

## Documentation

Usage docs are in ./doc

1. [Chapter 1 - Overview](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-1/index.md)

- Information about [Retrosheet](https://www.retrosheet.org/)
- Data downloads

2. [Chapter 2 - Transform CSV](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-2/index.md)

- Transforming Retrosheet data to JSON files
- uses node.js instead of Go (sorry)

3. [Chapter 3 - Read JSON](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-3/index.md)

- A Go package that reads the JSON files and converts them to Go objects
- Includes relevant tests
- Used in Chapter 4 for uploading to Mongodb

4. [Chapter 4 - Upload to Mongodb](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-4/index.md)

- A Go package that uses the package from Chapter 3 to read the JSON data and upload it to mongodb documents.
- Includes relevant tests

5. [Chapter 5 - Simple Queries](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-5/index.md)

- A Go package that provides some simple test queries to the Mongodb database.

6. [Chapter 6 - API and React Client](https://github.com/dmh2000/retrosheet/blob/main/doc/retro-chapter-6/index.md)

- An API for making queries with a sample React client

## RETROSHEET

The code in this repo my own creation and is not part of, supported by or endorsed by
retrosheet.org. The following are the (very generous) license terms for use of the
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

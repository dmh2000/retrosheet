const fs = require("fs");

const MongoClient = require("mongodb").MongoClient;

const glfields = require("./src/glfields");
const collections = require("./src/collections");
const query = require("./src/query");
const loader = require("./src/loader");

const url = "mongodb://192.168.5.3:27017";
const dbname = "baseball";
const fname = __dirname + "/data/GL2010.TXT";
// glfields.readFile(__dirname + "/data/GL2010.TXT");

// read the record and get JSON back (synchronous)
const game = glfields.readGameRecord(fname);

// save it as a json file (synchronous)
fs.writeFileSync(__dirname + "/gl2010x.json", JSON.stringify(game, null, 2));
/*
collections
  .listCollections(url, dbname)
  .then((a) => {
    console.log(a);
  })
  .then(() => {
    return query.find(url, dbname, "xyz", {}, {});
  })
  .then((q) => {
    console.log(q);
  })
  .catch((err) => {
    console.log(err);
  });
*/

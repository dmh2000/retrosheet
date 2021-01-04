const MongoClient = require("mongodb").MongoClient;
const fs = require("fs");
const path = require("path");

/**
 *
 * @param {*} url
 * @param {*} dbname
 * @param {*} query  query string
 * @param {*} fields
 * @return Promise
 */
function load(url, dbname, collectionName, documents) {
  const client = new MongoClient(url, { useUnifiedTopology: true });
  let db;

  return (
    client
      .connect()
      .then((client) => {
        // select the database
        db = client.db(dbname);
        // get the collection
        return db.collection(collectionName);
      })
      .then((collection) => {
        // insert the documents
        return collection.insertMany(documents);
      })
      // finish
      .then((result) => {
        client.close();
        return Promise.resolve(result);
      })
      .catch((err) => {
        return Promise.reject(err);
      })
  );
}

// ========================================
// TEST
// ========================================
const url = "mongodb://192.168.5.3:27017";
// const url = "mongodb://localhost:27017";
const dbname = "baseball";
const data = "/home/dmh2000/projects/baseball/data/games/json";
const fname = "gl2010.json";

const datapath = path.join(data, fname);
console.log(datapath);

const gameJson = fs.readFileSync(path.join(data, "gl2010.json"), { encoding: "UTF-8" });
const games = JSON.parse(gameJson);
console.log(games.length);

load(url, dbname, "games", games)
  .then((result) => {
    console.log(result);
  })
  .catch((error) => console.log(error));

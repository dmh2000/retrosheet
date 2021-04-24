const MongoClient = require("mongodb").MongoClient;
const fs = require("fs");
const glfields = require("./glfields");
const loader = require("./loader");

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

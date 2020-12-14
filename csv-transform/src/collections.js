const MongoClient = require('mongodb').MongoClient;

/**
 *
 * @param {*} url
 * @param {*} dbname
 * @return Promise
 */
function listCollections(url, dbname) {
  const client = new MongoClient(url, { useUnifiedTopology: true });
  let db;
  let list = [];

  return client
    .connect()
    .then((client) => {
      // select the database
      db = client.db(dbname);
      // get list of all collections
      cursor = db.listCollections();
      // forward it as a promise
      return cursor.toArray();
    })
    .then((a) => {
      // toArray promise resolves to array of collection objects
      client.close();
      return Promise.resolve(a.map((v) => v.name));
    })
    .catch((err) => {
      return Promise.reject(err);
    });
}

module.exports = {
  listCollections,
};

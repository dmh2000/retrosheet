const MongoClient = require('mongodb').MongoClient;

/**
 *
 * @param {*} url
 * @param {*} dbname
 * @param {*} query  query string
 * @param {*} fields
 * @return Promise
 */
function find(url, dbname, collectionName, query, options) {
  const client = new MongoClient(url, { useUnifiedTopology: true });
  let db;

  return client
    .connect()
    .then((client) => {
      // select the database
      db = client.db(dbname);
      // get the collection
      return db.collection(collectionName);
    })
    .then((collection) => {
      return collection.find(query, options);
    })
    .then((cursor) => {
      return cursor.toArray();
    })
    .then((a) => {
      client.close();
      return Promise.resolve(a);
    })
    .catch((err) => {
      return Promise.reject(err);
    });
}

// ========================================
// TEST
// ========================================
const url = 'mongodb://192.168.5.3:27017';
const dbname = 'baseball';

find(url, dbname, 'games', { date: '20100405' }, { projection: { date: 1, visitorTeam: 1, homeTeam: 1 } })
  .then((a) => {
    console.log(a);
  })
  .catch((err) => {
    console.log(err);
  });

module.exports = {
  find,
};

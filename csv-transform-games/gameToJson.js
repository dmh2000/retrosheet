const gameLogToJson = require('./src/gameLog');

if (process.argv.length < 4) {
  console.log('gameToJson <.../GLyyyy.TXT> <.../glyyyy.json>');
  process.exit(1);
}

const infile = process.argv[2];
const outfile = process.argv[3];

gameLogToJson(infile, outfile);

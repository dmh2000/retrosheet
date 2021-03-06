#!/usr/bin/env node
const fs = require("fs");
const gameLogToJson = require("./gameLog");

if (process.argv.length < 4) {
  console.log(
    "gamesToJson <directory with retrosheet gamelog .TXT files> <output directory>"
  );
  process.exit(1);
}

const indir = fs.opendirSync(process.argv[2]);
const outdir = fs.opendirSync(process.argv[3]);
for (;;) {
  const d = indir.readSync();
  if (d == null) {
    break;
  }
  const m = d.name.match(/GL(\d\d\d\d)\.TXT/);
  if (m != null) {
    console.log(`${indir.path}/${d.name}`);
    gameLogToJson(`${indir.path}/${d.name}`, `${outdir.path}/gl${m[1]}.json`);
  }
}

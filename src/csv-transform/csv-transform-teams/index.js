const fs = require("fs");
const parse = require("csv-parse/lib/sync");

const idFields = [
  "Abbr",
  "League",
  "City",
  "Nickname",
  "FirstYear",
  "LastYear",
];

function TeamsToJson(infile) {
  // read the file
  const data = fs.readFileSync(infile, "UTF-8");

  // parse into json
  const records = parse(data, {
    columns: idFields,
    skip_empty_lines: true,
  });

  // parse
  let json = records.map(JSON.stringify).join(",\n");
  json = "[\n" + json + "\n]\n";

  // output
  console.log(json);
}

if (process.argv.length < 3) {
  console.error("node index.js <source> >stdout");
  process.exit(1);
}
TeamsToJson(process.argv[2]);

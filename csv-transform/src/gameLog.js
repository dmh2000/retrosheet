const fs = require("fs");
const parse = require("csv-parse/lib/sync");
const gsf = require("./glfields");

function gameLogToJson(infile, outfile) {
  // read the file
  const data = fs.readFileSync(infile, "UTF-8");

  // parse into json
  const records = parse(data, {
    columns: gsf.glfields,
    skip_empty_lines: true,
    // cast to turn string numbers to literal ints
    // use for golang. javascript doesn't need it but it doesn't hurt
    cast: function (value, context) {
      if (gsf.gltypes[context.index] === "number") {
        return parseInt(value);
      } else {
        return value;
      }
    },
  });

  // parse
  let json = records.map(JSON.stringify).join(",\n");
  json = "[\n" + json + "\n]\n";

  // output
  fs.writeFileSync(outfile, json);
}

module.exports = gameLogToJson;

const fs = require("fs");
const parse = require("csv-parse/lib/sync");

const idFields = [
  "ID",
  "Last",
  "First",
  "PlayDebut",
  "MgrDebut",
  "CoachDebut",
  "UmpDebut",
];

function IDsToJson(infile, outfile) {
  // read the file
  const data = fs.readFileSync(infile, "UTF-8");

  // parse into json
  const records = parse(data, {
    columns: idFields,
    skip_empty_lines: true,
    // // cast to turn string numbers to literal ints
    // // use for golang. javascript doesn't need it but it doesn't hurt
    // cast: function (value, context) {
    //   if (context.index >= gsf.gltypes.length) {
    //     // undefined, treat as a string
    //     return value;
    //   } else if (gsf.gltypes[context.index] === "number") {
    //     // convert to int type
    //     return parseInt(value);
    //   } else {
    //     // leave anything else as a string
    //     return value;
    //   }
    // },
  });

  // parse
  let json = records.map(JSON.stringify).join(",\n");
  json = "[\n" + json + "\n]\n";

  // output
  fs.writeFileSync(outfile, json);
}

console.log(process.argv[2]);
IDsToJson(process.argv[2], "ids.json");

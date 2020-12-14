const glf = require('./src/glfields');

// turn the list of fields into a single line to prepend to game csv file to read in excel or d3
function csvFields() {
  let s = '';
  let comma = '';
  for (label of glf.glfields) {
    s += `${comma}${label}`;
    comma = ',';
  }
  console.log(s);
}

csvFields();

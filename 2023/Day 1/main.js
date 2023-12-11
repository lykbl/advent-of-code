const fs = require('fs');
const lines = fs.readFileSync('./input.txt', 'utf-8').split('\n');
let calibrationValues = [];

for (let line of lines) {
  let first, last, number;
  for (let character of line) {
    number = parseInt(character)
    if (!isNaN(number)) {
      if (first === undefined) {
        first = number;
        last = number;
      } else {
        last = number;
      }
    }
  }
  calibrationValues.push([first, last]);
}

const documentCalibration = calibrationValues.reduce(
  (acc, curr) => acc + parseInt(`${curr[0]}${curr[1]}`),
  0
);
console.log(documentCalibration);
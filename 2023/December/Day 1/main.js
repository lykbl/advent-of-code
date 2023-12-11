const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const lines = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8').split('\n');
let calibrationValues = [];

const  NUMBERS = {
  'one': 1,
  'two': 2,
  'three': 3,
  'four': 4,
  'five': 5,
  'six': 6,
  'seven': 7,
  'eight': 8,
  'nine': 9,
}

function detectNumber(line, leftToRight = true) {
  const [start, end, increment] = leftToRight
    ? [0, line.length, 1]
    : [line.length - 1, -1, -1]
  ;
  let numberLetters = '';

  for (let i = start; i !== end; i += increment) {
    let character = line[i];
    if (!isNaN(parseInt(character))) {
      return parseInt(character);
    }

    numberLetters = leftToRight ? `${numberLetters}${character}` : `${character}${numberLetters}`;
    if (!canBeNumber(numberLetters)) {
      while (!canBeNumber(numberLetters) && numberLetters.length > 0) {
        numberLetters = leftToRight ? numberLetters.slice(1) : numberLetters.slice(0, -1);
      }
    }

    if (NUMBERS[numberLetters]) {
      return NUMBERS[numberLetters]
    }
  }

  return null;
}

function canBeNumber(characters) {
  return Object.keys(NUMBERS).filter((possibleNumber) => possibleNumber.includes(characters)).length > 0;
}

for (let line of lines) {
  let first, last;

  first = detectNumber(line, true);
  last = detectNumber(line, false);

  calibrationValues.push([first, last ?? first]);
}

const documentCalibration = calibrationValues.reduce(
  (acc, curr) => acc + parseInt(`${curr[0]}${curr[1]}`),
  0
);

console.log(documentCalibration);
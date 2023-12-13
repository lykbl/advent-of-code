const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const isPartCharacter = (char) => char !== undefined && isNaN(parseInt(char)) && char !== '.' && char !== '\n';

const isGearCharacter = (char) => char === '*';
function analyzeNumber(input, firstDigitIndex, lineLength) {
  const currentLine = Math.floor(firstDigitIndex / lineLength);
  const lineStart = currentLine * lineLength
  const lineEnd = lineStart + lineLength;
  let middleCursor = firstDigitIndex - 1 > lineStart ? firstDigitIndex - 1 : firstDigitIndex;
  let lastDigitIndex = firstDigitIndex;
  let topCursor, bottomCursor;
  let isPartNumber = false;
  let value = '';

  while (middleCursor >= lineStart && middleCursor <= lineEnd) {
    topCursor = middleCursor - lineLength;
    bottomCursor = middleCursor + lineLength;
    isPartNumber = isPartCharacter(input[topCursor]) || isPartCharacter(input[bottomCursor]) || isPartNumber;

    if (isNaN(parseInt(input[middleCursor]))) {
      isPartNumber = isPartCharacter(input[middleCursor]) || isPartNumber;
      if (middleCursor > firstDigitIndex) {
        break;
      }
    } else {
      value += input[middleCursor];
      lastDigitIndex = middleCursor;
    }

    middleCursor += 1;
  }

  return {
    value: parseInt(value),
    firstDigitIndex,
    lastDigitIndex,
    isPartNumber,
  }
}

function detectGears(input, firstDigitIndex, lineLength) {
  const currentLine = Math.floor(firstDigitIndex / lineLength);
  const lineStart = currentLine * lineLength
  const lineEnd = lineStart + lineLength;
  let middleCursor = firstDigitIndex - 1 > lineStart ? firstDigitIndex - 1 : firstDigitIndex;
  let lastDigitIndex = firstDigitIndex;
  let topCursor, bottomCursor;
  let value = '';
  let gearLocations = [];

  while (middleCursor >= lineStart && middleCursor <= lineEnd) {
    topCursor = middleCursor - lineLength;
    bottomCursor = middleCursor + lineLength;
    for (let cursor of [topCursor, bottomCursor, middleCursor]) {
      if (isGearCharacter(input[cursor])) {
        gearLocations.push(cursor);
      }
    }

    if (isNaN(parseInt(input[middleCursor]))) {
      //better way to break out?
      if (middleCursor > firstDigitIndex) {
        break;
      }
    } else {
      value += input[middleCursor];
      lastDigitIndex = middleCursor;
    }

    middleCursor += 1;
  }

  return {
    value: parseInt(value),
    gearLocations,
  };
}

const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8');
const lineLength = input.indexOf('\n') + 1;
const partNumbers = [];
const gearsMap = {};
for (let i = 0; i < input.length; i++) {
  const char = input[i];
  if (!isNaN(parseInt(char))) {
    const numberAnalysis = analyzeNumber(input, i, lineLength);
    if (numberAnalysis.isPartNumber) {
      partNumbers.push(numberAnalysis);
    }
    const gearAnalysis = detectGears(input, i, lineLength);
    for (let gearLocation of gearAnalysis.gearLocations) {
      gearsMap[gearLocation] === undefined ? gearsMap[gearLocation] = [gearAnalysis.value] : gearsMap[gearLocation].push(gearAnalysis.value);
    }

    i = numberAnalysis.lastDigitIndex + 1;
  }
}

console.log(gearsMap)
let totalGearRatio = 0;
for (let gearLocation in gearsMap) {
  if (gearsMap[gearLocation].length === 2) {
    totalGearRatio += gearsMap[gearLocation][0] * gearsMap[gearLocation][1];
  }
}
console.log(totalGearRatio)
// console.log(partNumbers.map(partNumber => partNumber.value))
// console.log(partNumbers.reduce((acc, curr) => acc + curr.value, 0));
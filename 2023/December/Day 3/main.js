const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const isPartCharacter = (char) => char !== undefined && isNaN(parseInt(char)) && char !== '.' && char !== '\n';
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

const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8');
const lineLength = input.indexOf('\n') + 1;
const partNumbers = [];
for (let i = 0; i < input.length; i++) {
  const char = input[i];
  if (!isNaN(parseInt(char))) {
    const numberAnalysis = analyzeNumber(input, i, lineLength);
    if (numberAnalysis.isPartNumber) {
      partNumbers.push(numberAnalysis);
    }

    i = numberAnalysis.lastDigitIndex + 1;
  }
}

console.log(partNumbers.map(partNumber => partNumber.value))
console.log(partNumbers.reduce((acc, curr) => acc + curr.value, 0));
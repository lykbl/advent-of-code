const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
const lines = input.split('\n');

let result = 0;
let result2 = 0;
for (const line of lines) {
  const histories = [line.split(' ').map(Number)];

  let allZeroes = false
  let sequenceI = 0;
  let numberI = histories[sequenceI].length - 1;
  histories[sequenceI + 1] = Array.from({ length: numberI - 1 })//.from({ length: numberI - 1 });
  while (!allZeroes) {
    if (numberI === 0) {
      numberI = histories[++sequenceI].length - 1;
      histories[sequenceI + 1] = Array.from({ length: numberI });
      allZeroes = true;
    }

    histories[sequenceI + 1][numberI - 1] = histories[sequenceI][numberI] - histories[sequenceI][numberI - 1];
    allZeroes = histories[sequenceI + 1].every(n => n === 0); // :(
    numberI--;
  }

  let [leftAdded, rightAdded] = [0, 0];
  for (sequenceI = histories.length - 2; sequenceI >= 0; sequenceI--) {
    const [leftMostValue, rightMostValue] = [histories[sequenceI][0], histories[sequenceI][histories[sequenceI].length - 1]];
    rightAdded += rightMostValue;
    leftAdded = leftMostValue - leftAdded;
  }

  result += rightAdded;
  result2 += leftAdded;
}

console.log(result)
console.log(result2)
// 1789635132 p1
// 913 p2
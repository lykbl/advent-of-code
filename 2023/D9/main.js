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
  let y = 0;
  let i = histories[y].length - 1;
  histories[y + 1] = Array.from({ length: i - 1 });
  while (!allZeroes) {
    if (i === 0) {
      y += 1;
      i = histories[y].length - 1;
      histories[y + 1] = Array.from({ length: i });
    }

    const diff = histories[y][i] - histories[y][i - 1];
    histories[y + 1][i - 1] = diff;
    allZeroes = histories[y + 1].every(n => n === 0);

    i--;
  }

  for (y = histories.length - 1; y >= 0; y--) {
    if (y === histories.length - 1 ) {
      histories[y].push(0)
    } else {
      const rightMostValue = histories[y][histories[y].length - 1];
      const belowValue = histories[y + 1][histories[y + 1].length - 1];
      histories[y].push(rightMostValue + belowValue);
    }
  }
  const rightAdded = histories[0][histories[0].length - 1];

  for (y = histories.length - 1; y >= 0; y--) {
    if (y === histories.length - 1) {
      histories[y].unshift(0)
    } else {
      const leftMostValue = histories[y][0];
      const belowValue = histories[y + 1][0];
      histories[y].unshift(leftMostValue - belowValue);
    }
  }
  const leftAdded = histories[0][0];

  result += rightAdded;
  result2 += leftAdded;
}

console.log(result)
console.log(result2)
// 1789635132 p1
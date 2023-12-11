const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
function isGamePossible(setsLine, constraints) {
  let cubesCount = '';
  let color = '';

  for (let char of setsLine) {
    if (char === ' ') {
      continue;
    } else if (!isNaN(parseInt(char))) {
      cubesCount += char;
    } else if (char === ',' || char === ';') {
      if (constraints[color] < parseInt(cubesCount)) {
        return false;
      }
      cubesCount = '';
      color = '';
    } else {
      color += char;
    }
  }

  return true;
}

const gameLines = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8').split('\n');
const constraints = {
  'red': 12,
  'green': 13,
  'blue': 14,
};
let result = 0;
const impossibleGames = [];
for (let gameLine of gameLines) {
  const [gameLabel, setsLine] = gameLine.split(': ');
  const gameId = new RegExp(/Game ([\d]+)/).exec(gameLabel)[1];
  if (isGamePossible(setsLine, constraints)) {
    result += parseInt(gameId);
  } else {
    impossibleGames.push(gameId);
  }
}

console.log(impossibleGames)
console.log(result)
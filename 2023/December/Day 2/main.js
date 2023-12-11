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

function fewestCubesNeeded(setsLine) {
  let cubesCount = '';
  let color = '';
  const fewestCubesNeeded = {
    'red': 0,
    'green': 0,
    'blue': 0,
  };

  for (let i = 0; i < setsLine.length; i++) {
    const char = setsLine[i];
    if (char === ' ') {
      continue;
    } else if (!isNaN(parseInt(char))) {
      cubesCount += char;
    } else if (char === ',' || char === ';') {
      fewestCubesNeeded[color] = Math.max(fewestCubesNeeded[color], parseInt(cubesCount));
      cubesCount = '';
      color = '';
    } else {
      color += char;
      if (i === setsLine.length - 1) {
        fewestCubesNeeded[color] = Math.max(fewestCubesNeeded[color], parseInt(cubesCount));
      }
    }
  }

  let power = 1;
  for (let count in fewestCubesNeeded) {
    power *= fewestCubesNeeded[count];
  }

  return power;
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

const gameStats = [];
for (let gameLine of gameLines) {
  const [gameLabel, setsLine] = gameLine.split(': ');
  const gameId = new RegExp(/Game ([\d]+)/).exec(gameLabel)[1];
  const fewestCubes = fewestCubesNeeded(setsLine);
  gameStats.push({
    gameId,
    fewestCubes,
  });
}

const totalPower = gameStats.reduce((acc, curr) => acc + curr.fewestCubes, 0);

console.log(impossibleGames)
console.log(result)
console.log(gameStats)
console.log(totalPower)
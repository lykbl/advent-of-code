const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const isDigit = (character) => !isNaN(parseInt(character))

const games = [];
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8');
let luckyNumbers, gameResult, numberAcc;
let gameState = 0;

for (let i = 0; i < input.length; i++) {
  let character = input[i];
  if (character === ':') {
    gameState = 1;
    luckyNumbers = {};
    gameResult = 0;
    numberAcc = '';
  }

  if (gameState < 1) {
    continue;
  }

  if (isDigit(character)) {
    numberAcc += character;
  }

  if (character === '|') {
    gameState = 2;
  }

  if ((character === ' ' || character === '\n') && numberAcc.length) {
    if (gameState === 1) {
      luckyNumbers[numberAcc] = true;
    } else if (gameState === 2 && luckyNumbers[numberAcc]) {
      gameResult = gameResult === 0 ? 1 : gameResult * 2;
    }
    numberAcc = '';
  }
  if (character === '\n') {
    numberAcc = '';
    gameState = 0;
    luckyNumbers = {};
    games.push(gameResult);
  }
}

console.log(games)
console.log(games.reduce((acc, curr) => acc + curr, 0));
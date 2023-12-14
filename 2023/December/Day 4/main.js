const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const isDigit = (character) => !isNaN(parseInt(character))

const games = [];
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf-8');
let luckyNumbers, gameResult, numberAcc, copiedCardOffset, copiedCardIndex, currentCardCopies;
let gameState = 0;
let gameIndex = -1;

Array.prototype.incrementIndex = function(index, value) {
  this[index] === undefined ? this[index] = value : this[index] += value;
}

for (let i = 0; i < input.length; i++) {
  let character = input[i];
  if (character === ':') {
    [gameState, luckyNumbers, numberAcc, copiedCardOffset, gameIndex] = [1, {}, '', 0, ++gameIndex];
    games.incrementIndex(gameIndex, 1);
    // gameResult = 0;
  }

  if (gameState < 1) {
    continue;
  }

  if (character === '|') {
    gameState = 2;
  }

  if (isDigit(character)) {
    numberAcc += character;
  }

  if ([' ', '\n'].includes(character) && numberAcc.length) {
    if (gameState === 1) {
      luckyNumbers[numberAcc] = true;
    }

    if (gameState === 2 && luckyNumbers[numberAcc]) {
      copiedCardOffset += 1;
      copiedCardIndex = copiedCardOffset + gameIndex;
      currentCardCopies = games[gameIndex];
      games.incrementIndex(copiedCardIndex, currentCardCopies)
      // gameResult = gameResult === 0 ? 1 : gameResult * 2;
    }

    numberAcc = '';
    gameState = character === '\n' ? 0 : gameState;
  }
}

console.log(games)
console.log(games.reduce((acc, curr) => acc + curr, 0));
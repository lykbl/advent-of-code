const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
let rows = input.split('\n').map(row => row.split(''));
function moveBoulders(rows, direction = 'north') {
  const movingType = ['north', 'south'].includes(direction) ? 'vertical' : 'horizontal';
  const movingDirection = ['north', 'west'].includes(direction) ? 'negative' : 'positive';
  const columns = rows[0];

  const traversalStart = movingType === 'vertical' ? {
    firstLevel: rows, secondLevel: columns
  } : {
    firstLevel: columns, secondLevel: rows
  };
  const movingDimensionStart  = movingDirection === 'negative' ? traversalStart.firstLevel.length - 1 : 0;
  const movingDimensionEnd    = movingDirection === 'negative' ? 0 : traversalStart.firstLevel.length - 1;
  const movingDimensionChange = movingDirection === 'negative' ? -1 : 1;
  const movingCharEnd         = traversalStart.secondLevel.length - 1;
  const bouldersToMove   = Object.fromEntries(Array.from({ length: movingCharEnd + 1 }, (_, i) => [i, []]));

  let movingDimensionIndex = movingDimensionStart;
  while (movingDirection === 'negative' ? movingDimensionIndex >= movingDimensionEnd : movingDimensionIndex <= movingDimensionEnd) {
    let movingCharIndex = 0;
    while(movingCharIndex <= movingCharEnd) {
      const movingChar = movingType === 'vertical' ? rows[movingDimensionIndex][movingCharIndex] : rows[movingCharIndex][movingDimensionIndex];
      if (movingChar === 'O') {
        bouldersToMove[movingCharIndex].push(movingCharIndex);
      } else if (movingChar === '#') {
        bouldersToMove[movingCharIndex] = [];
      } else if ((movingChar === '.' || movingDimensionIndex === movingDimensionEnd) && bouldersToMove[movingCharIndex].length > 0) {
        const dotIndex = movingDimensionIndex + bouldersToMove[movingCharIndex].length * -movingDimensionChange;
        if (movingType === 'vertical') {
          [rows[movingDimensionIndex][movingCharIndex], rows[dotIndex][movingCharIndex]] = [rows[dotIndex][movingCharIndex], rows[movingDimensionIndex][movingCharIndex]];
        } else {
          [rows[movingCharIndex][movingDimensionIndex], rows[movingCharIndex][dotIndex]] = [rows[movingCharIndex][dotIndex], rows[movingCharIndex][movingDimensionIndex]];
        }
      }
      movingCharIndex++;
    }
    movingDimensionIndex += movingDimensionChange
  }

  return rows;
}
function calculatePlatformWeight(rows) {
  let weight = 0;
  for (let row = rows.length - 1; row >= 0; row--) {
    for (let col = 0; col < rows[row].length; col++) {
      if (rows[row][col] === 'O') {
        weight += rows.length - row;
      }
    }
  }

  return weight;
}

const movedPlatform = moveBoulders(rows, 'north');
console.log(calculatePlatformWeight(movedPlatform))

const cyclesCount = 1000000000;
const memory = new Map();
for (let i = 0; i < cyclesCount; i += 1) {
  const rowsHash = rows.map(row => row.join('')).join('');
  if (memory.has(rowsHash)) {
    const cycleSize = memory.size - memory.get(rowsHash);
    let remainingIterations = (cyclesCount - i) % cycleSize;
    i = cyclesCount - remainingIterations;
  }
  for (let direction of ['north', 'west', 'south', 'east']) {
    rows = moveBoulders(rows, direction);
    memory.set(rowsHash, i);
  }
}
console.log(calculatePlatformWeight(rows));
//107951
//95736
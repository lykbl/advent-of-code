const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const image = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

const isGalaxy = (char) => char === '#';
class Galaxy {
  constructor(x, y, label) {
    this.x = x;
    this.y = y;
    this.label = label;
  }
}

const universe = detectGalaxies(image);
const grandUniverse = detectGalaxies(image);
function detectGalaxies(skyImage) {
  const lines = skyImage.split('\n');
  const galaxyCoordinates = [];
  let galaxyLabels = 1;

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    for (let j = line.length - 1; j >= 0; j--) {
      const char = line[j];
      const offset = line.length - j;
      if (isGalaxy(char)) {
        galaxyCoordinates.push(new Galaxy(offset, i, galaxyLabels++));
      }
    }
  }

  return galaxyCoordinates;
}
function expandUniverse(universe, expansionFactor = 2) {
  for (let type of ['x', 'y']) {
    let expansion = 0;
    universe = universe.sort((a, b) => a[type] - b[type]);
    for (let i = 0; i < universe.length - 1; i++) {
      const currentGalaxy = universe[i];
      const nextGalaxy = universe[i + 1];

      const diff = (nextGalaxy[type] + expansion) - currentGalaxy[type] - 1;
      if (diff >= 1) {
        expansion += diff * (expansionFactor - 1);
      }
      universe[i + 1][type] += expansion;
    }

  }

  return universe;
}
function distanceBetweenGalaxies(galaxyA, galaxyB) {
  return Math.abs(galaxyA.y - galaxyB.y) + Math.abs(galaxyA.x - galaxyB.x)
}
function totalUniversePaths(universe) {
  return universe.reduce((totalPairsPath, galaxy, i) =>
      totalPairsPath +
      universe.slice(i + 1).reduce((pairPath, nextGalaxy) => pairPath + distanceBetweenGalaxies(galaxy, nextGalaxy), 0),
  0);
}

const basicExpandedUniverse = expandUniverse(universe);
const expandedUniverse = expandUniverse(grandUniverse, 1000000);

console.log(totalUniversePaths(basicExpandedUniverse))
console.log(totalUniversePaths(expandedUniverse))

//p1 9543156
//p2 625243292686
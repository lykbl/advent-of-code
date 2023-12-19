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
function expandUniverse(sortedGalaxies, expansionType, expansionFactor = 2) {
  let expansion = 0;
  for (let i = 0; i < sortedGalaxies.length - 1; i++) {
    const currentGalaxy = sortedGalaxies[i];
    const nextGalaxy = sortedGalaxies[i + 1];

    const diff = (nextGalaxy[expansionType] + expansion) - currentGalaxy[expansionType] - 1;
    if (diff >= 1) {
      expansion += diff * (expansionFactor - 1);
    }
    sortedGalaxies[i + 1][expansionType] += expansion;
  }

  return sortedGalaxies;
}
function totalUniversePaths(universe) {
  let totalPath = 0;
  for (let i = 0; i < universe.length; i++) {
    const galaxy = universe[i];
    for (let y = i + 1; y < universe.length; y++) {
      const nextGalaxy = universe[y];
      const shortestPath = Math.abs(nextGalaxy.y - galaxy.y) + Math.abs(nextGalaxy.x - galaxy.x);
      totalPath += Math.abs(shortestPath);
    }
  }

  return totalPath;
}

const expansionFactor = 1000000;
const basicExpandedUniverse = expandUniverse(
  expandUniverse(JSON.parse(JSON.stringify(universe)).sort((a, b) => a.x - b.x), 'x').sort((a, b) => a.y - b.y),
  'y',
);
const expandedUniverse = expandUniverse(
  expandUniverse(universe.sort((a, b) => a.x - b.x), 'x', expansionFactor).sort((a, b) => a.y - b.y),
  'y',
  expansionFactor
);

console.log(totalUniversePaths(basicExpandedUniverse))
console.log(totalUniversePaths(expandedUniverse))

//p1 9543156
//p2 625243292686
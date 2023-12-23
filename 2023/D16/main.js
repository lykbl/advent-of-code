const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

const grid = input.split('\n').map(row => row.split(''));
const addMatrix = (a, b) => a.map((_, i) => a[i] + b[i])
const rotateMatrix = ([x, y], clockWise = true) => clockWise ? [y, x] : [-y, -x]
function energizeGrid(originalGrid, beams = [[-1, 0, [1, 0]]]) {
  const energized = new Map();
  const gridLoops = new Map();
  let beamIndex = 0;
  while (beamIndex < beams.length) {
    while (true) {
      const beam = beams[beamIndex];
      if (beam === null) break;
      const [x, y, direction] = beam;
      const nextCoordinates = addMatrix([x, y], direction);
      const [nextX, nextY] = nextCoordinates;
      const cellHash = `${ nextX }_${ nextY }_${ direction[0] }_${ direction[1] }`;
      if (grid[nextY] === undefined || grid[nextY][nextX] === undefined || gridLoops.has(cellHash)) {
        beams[beamIndex] = null;
        break;
      }

      gridLoops.set(cellHash, true)
      energized.set(`${nextX}_${nextY}`, true);
      const nextTile = grid[nextY][nextX];
      if (nextTile === '|' && direction[0] !== 0) {
        beams.push([nextX, nextY, [0, -1]], [nextX, nextY, [0, 1]])
        beams[beamIndex] = null;
      } else if (nextTile === '-' && direction[1] !== 0) {
        beams.push([nextX, nextY, [-1, 0]], [nextX, nextY, [1, 0]])
        beams[beamIndex] = null;
      } else if (['/', '\\'].includes(nextTile)) {
        beams[beamIndex] = [nextX, nextY, rotateMatrix(direction, nextTile === '\\')];
      } else {
        beams[beamIndex] = [nextX, nextY, direction];
      }
    }
    beamIndex++;
  }

  return energized.size;
}

const { performance, PerformanceObserver } = require('perf_hooks');
const obs = new PerformanceObserver((items) => {
  const durations = items.getEntries().reduce((acc, entry) => entry.duration, 0);
  console.log('Average durations', durations / items.getEntries().length)
  performance.clearMarks();
});
obs.observe({ entryTypes: ['measure'] });

function maxEnergized(grid) {
  let maxEnergized = 0;
  for (let x of [-1, grid[0].length]) {
    const direction = x === -1 ? [1, 0] : [-1, 0];
    for (let y = 0; y < grid.length; y++) {
      maxEnergized = Math.max(maxEnergized, energizeGrid(grid, [[x, y, direction]]));
    }
  }
  for (let y of [-1, grid.length]) {
    const direction = y === -1 ? [0, 1] : [0, -1];
    for (let x = 0; x < grid[0].length; x++) {
      maxEnergized = Math.max(maxEnergized, energizeGrid(grid, [[x, y, direction]]));
    }
  }
  return maxEnergized;
}

const p1 = energizeGrid(grid);
console.log(p1) // 7060
for (let i = 0; i < 10; i++) {
  performance.mark(`start_${i}`);
  const p2 = maxEnergized(grid);
  performance.mark(`end_${i}`);
  performance.measure('start to end', `start_${ i }`, `end_${ i }`);
  console.log(p2) // 7493
}

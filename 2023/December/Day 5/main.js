const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
// const BinaryHeap = require("../../../utils/binaryHeap");
const { extractLineNumbers } = require("../../../utils/helpers");

const rl = readline.createInterface({
  input: fileStream,
  crlfDelay: Infinity,
});
// let seeds = new BinaryHeap();
// let sources = new BinaryHeap();
let seeds = [];
let seedRanges = [];
let sources = [];
let map = {};
let seedsLoaded = false

function recalibrateSeedRanges(seedRanges, sources, map) {
  let mapIndex = 0;
  let seedIndex = 0;
  let newSeeds = [];
  while (seedIndex < seedRanges.length) {
    const [seedRangeStart, seedRangeEnd] = [seedRanges[seedIndex], seedRanges[seedIndex + 1]];
    const [sourceRangeStart, sourceRangeEnd] = [sources[mapIndex], sources[mapIndex + 1]];
    let [keepIndexes, seedRangeStartToTransform, seedRangeEndToTransform] = [false, seedRangeStart, seedRangeEnd];

    if (sourceRangeStart === undefined) {
      newSeeds.push(seedRangeStart, seedRangeEnd)
      seedIndex += 2;
      continue;
    }

    if (seedRangeStart > sourceRangeEnd) {
      mapIndex += 2;
      continue;
    }

    const leftMostRange = Math.min(seedRangeStart, sourceRangeStart);
    const rightMostRange = Math.max(seedRangeEnd, sourceRangeEnd);

    if (leftMostRange === seedRangeStart && leftMostRange !== sourceRangeStart) {
      const oldSeedRangeStart = seedRangeStart;
      const oldSeedRangeEnd = sourceRangeStart - 1;
      newSeeds.push(oldSeedRangeStart, oldSeedRangeEnd)

      seedRangeStartToTransform = sourceRangeStart;
    }

    if (rightMostRange === seedRangeEnd && rightMostRange !== sourceRangeEnd) {
      seedRangeEndToTransform = sourceRangeEnd;

      keepIndexes = true;
      seedRanges[seedIndex] = sourceRangeEnd + 1;
      mapIndex += 2;
    }

    const startOffset = seedRangeStartToTransform - sourceRangeStart;
    const endOffset = seedRangeEndToTransform - sourceRangeStart;
    newSeeds.push(map[sourceRangeStart] + startOffset, map[sourceRangeStart] + endOffset)

    seedIndex += keepIndexes ? 0 : 2;
  }

  return newSeeds;
}

function recalibrateSeeds(seeds, sources, map) {
  let mapIndex = 0;
  let seedIndex = 0;
  while (seedIndex < seeds.length) {
    const seed = seeds[seedIndex];
    if (seed >= sources[mapIndex] && seed <= sources[mapIndex + 1]) {
      const offset = seed - sources[mapIndex];
      seeds[seedIndex] = map[sources[mapIndex]] + offset
      seedIndex++;
    } else if (seed > sources[mapIndex + 1]) {
      mapIndex += 2
    } else if (seed < sources[mapIndex] || sources[mapIndex] === undefined) {
      seeds[seedIndex] = seed;
      seedIndex++;
    }
  }

  return seeds;
}

rl.on('line', (line) => {
  if (line.startsWith('seeds')) {
    const rangesExtractor = extractLineNumbers(line, null, 2);
    while (true) {
      const {value, done} = rangesExtractor.next();
      if (done) {
        break;
      }

      const [seedRangeStart, seedsCount] = value || [];
      seedRanges.push(seedRangeStart, seedRangeStart + seedsCount - 1);
    }
    seeds.push(...extractLineNumbers(line)); // p1
    seedsLoaded = true;
    return;
  }

  if (line.length === 0 && seedsLoaded && sources.length) {
    seeds.sort((a, b) => a - b)
    && seedRanges.sort((a, b) => a - b)
    && sources.sort((a, b) => a - b);

    seeds = recalibrateSeeds(seeds, sources, map);
    seedRanges = recalibrateSeedRanges(seedRanges, sources, map);

    sources = [];
    map = {};
  }

  if (!line.includes('map') && line.length) {
    const extractor = extractLineNumbers(line, null, 3);
    const [destination, source, range] = extractor.next().value;
    map[source] = destination;
    sources.push(source, source + range - 1);
  }
})

rl.on('close', () => {
  seeds.sort((a, b) => a - b)
  && seedRanges.sort((a, b) => a - b)
  && sources.sort((a, b) => a - b);

  seeds = recalibrateSeeds(seeds, sources, map);
  seeds.sort((a, b) => a - b)

  seedRanges = recalibrateSeedRanges(seedRanges, sources, map);
  seedRanges.sort((a, b) => a - b)
  console.log(seeds[0])
  console.log(seedRanges[0])
})

//165788812 answer P1
//1928058 answer p2
//TODO improve with binary heaps?
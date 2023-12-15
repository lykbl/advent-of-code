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
let sources = [];
let map = {};
let seedsLoaded = false

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
    seeds.push(...extractLineNumbers(line, '|'));
    seedsLoaded = true;
  } else if (line.length === 0 && seedsLoaded) {
    seeds.sort((a, b) => a - b) && sources.sort((a, b) => a - b);
    seeds = recalibrateSeeds(seeds, sources, map);
    sources = [];
    map = {};
  }

  if (!line.includes('map') && line.length) {
    const extractor = extractLineNumbers(line);
    const [destination, source, range] = Array.from({ length: 3 }, () => extractor.next().value);
    map[source] = destination;
    sources.push(source, source + range - 1);
  }
})

rl.on('close', () => {
  seeds.sort((a, b) => a - b) && sources.sort((a, b) => a - b);
  seeds = recalibrateSeeds(seeds, sources, map);
  console.log(seeds[0])
})

//165788812 answer
//172838688 too big
//367925813 too big
const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

class DesertMap {
  constructor() {
    this.nodes = new Map();
  }

  addNode(key, left, right) {
    this.nodes.set(key, { left, right });
  }

  getNode(key) {
    return this.nodes.get(key);
  }
}

const lines = input.split('\n');
const instructions = lines[0];
const nodeLines = lines.slice(2);

function buildMap(nodeLines) {
  const desertMap = new DesertMap();
  for (const nodeLine of nodeLines) {
    const {key, left, right} = nodeLine.match(/(?<key>[A-Z]{3}) = \((?<left>[A-Z]{3}), (?<right>[A-Z]{3})\)/).groups;
    desertMap.addNode(key, left, right);
  }

  return desertMap;
}

const desertMap = buildMap(nodeLines)
const startNode = 'AAA';
const finalNode = 'ZZZ';

let currentNode = startNode;
let i = 0;
let stepsCount = 0;

while (currentNode !== finalNode) {
  const { left, right } = desertMap.getNode(currentNode);
  const nextStep = instructions[i];

  currentNode = nextStep === 'L' ? left : right;
  i++;
  stepsCount++;

  if (i === instructions.length) {
    i = 0;
  }
}

console.log(stepsCount)

//14429 p1
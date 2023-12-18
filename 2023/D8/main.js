const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

class DesertMap {
  constructor() {
    this.nodes = new Map();
    this.startNodes = [];
  }

  addNode(key, left, right) {
    this.nodes.set(key, { left, right });
    if (key[2] === 'A') {
      this.startNodes.push(key);
    }
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

const startingNodes = desertMap.startNodes;
const cycles = [];
for (const startNode of startingNodes) {
  let stepsCount = 0;
  let i = 0;
  let currentNode = startNode;
  while (currentNode[2] !== 'Z') {
    const { left, right } = desertMap.getNode(currentNode);
    const nextStep = instructions[i];

    currentNode = nextStep === 'L' ? left : right;
    i++;
    stepsCount++;

    if (i >= instructions.length) {
      i = 0;
    }
  }
  cycles.push(stepsCount);
}

function gcd(a, b) {
  return b === 0 ? a : gcd(b, a % b);
}

// Function to calculate the Least Common Multiple (LCM)
function lcm(a, b) {
  return (a * b) / gcd(a, b);
}

console.log(cycles.reduce((accumulator, currentValue) => lcm(accumulator, currentValue)))

// console.log(stepsCount)

//14429 p1
//10921547990923 p2
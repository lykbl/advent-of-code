const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
const { extractLineNumbers } = require("../../utils/helpers");

const lines = input.split('\n');
const times = [...extractLineNumbers(lines[0])];
const distances = [...extractLineNumbers(lines[1])];
const p2time = [...extractLineNumbers(lines[0], null, 1, [' '])][0];
const p2distance = [...extractLineNumbers(lines[1], null, 1, [' '])][0];

function solveQuadratic(b, c, a = 1) {
  const x1 =  (-b + Math.sqrt(b ** 2 - 4 * a * c)) / (2 * a);
  const x2 =  (-b - Math.sqrt(b ** 2 - 4 * a * c)) / (2 * a);

  const from = Math.ceil(x1);
  const to = Math.floor(x2);

  return Math.abs((to) - (from < x1 ? from : from - 1));
}

let resultp1 = 1;
for (let i = 0; i < times.length; i++) {
  const solutionsCount = solveQuadratic(-times[i], distances[i]);
  resultp1 *= solutionsCount;
}
const resultp2 = solveQuadratic(-p2time, p2distance);

console.log(resultp1);
console.log(resultp2);

//p1 275724
//p2 37286485
const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
// const BinaryHeap = require("../../../utils/binaryHeap");
const { extractLineNumbers } = require("../../utils/helpers");

const rl = readline.createInterface({
  input: fileStream,
  crlfDelay: Infinity,
});

rl.on('line', (line) => {
})

rl.on('close', () => {
})

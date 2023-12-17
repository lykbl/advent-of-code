const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
// const { extractLineNumbers } = require("../../utils/helpers");
const BinaryHeap = require("../../utils/binaryHeap");

const rs = readline.createInterface({
  input: fileStream,
  crlfDelay: Infinity,
});

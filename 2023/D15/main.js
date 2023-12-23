const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

let p1 = 0;
const stringsToHash = input.split(',');
const boxes = new Map();
const stringToHash = (stringToHash) => stringToHash.reduce((acc, char) => ((acc + char.charCodeAt(0)) * 17) % 256, 0)

for (let string of stringsToHash) {
  let stringHash = stringToHash(string.split(''));
  p1 += stringHash;

  const splitChar = string.indexOf('=') !== -1 ? '=' : '-';
  const [newLabel, newFocalLength] = string.split(splitChar);
  const labelHash = stringToHash(newLabel.split(''));
  let currentBox = boxes.get(labelHash) || [];

  if (splitChar === '=') {
    const currentBoxIndex = currentBox.findIndex(({ label }) => label === newLabel);
    const replacement = { label: newLabel, focalLength: newFocalLength };
    currentBoxIndex === -1
      ? currentBox.push(replacement)
      : currentBox[currentBoxIndex] = replacement;
  } else {
    currentBox = currentBox.filter(({ label }) => label !== newLabel);
  }
  boxes.set(labelHash, currentBox);
}

let p2 = 0;
boxes.forEach((box, boxI) => {
  box.forEach(({ label, focalLength }, i) => {
    const multiplier = i + 1;
    p2 += (boxI + 1) * focalLength * multiplier;
  })
})

console.log(p1)
console.log(p2)

//507666 p1
//233537 p2
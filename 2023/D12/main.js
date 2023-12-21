const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const springArrangements = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

const springArrangedAt = (schema, springStart, springSize) => {
  let springEnd = springStart;
  while (springEnd < springStart + springSize) {
    if (schema[springEnd] === '.') {
      return -1;
    }
    springEnd++;
  }

  if ((['.', '?'].includes(schema[springEnd]) || springEnd === schema.length)) {
    return springEnd;
  }

  return -1;
}
function calculateArrangements(schema, springSizes, basePath = '', memory = {}) {
  if (springSizes.length === 0) {
    if (schema.indexOf('#') === -1) {
      // console.log(basePath.replaceAll('?', '.'))
      return 1;
    } else {
      return 0;
    }
  }

  let currentSpringSize = springSizes[0];
  let currentArrangements = 0;
  let springBank = currentSpringSize;

  for (let i = 0; i < schema.length; i++) {
    if (springBank < currentSpringSize) {
      continue;
    }
    let springEnd = springArrangedAt(schema, i, springBank)
    if (schema[i] === '#') {
      springBank -= 1;
    }
    if (springEnd === -1) {
      continue;
    }
    const remainingSprings = springSizes.slice(1);
    const remainingSchema = schema.slice(springEnd + 1); //slice is stupid

    let toAdd = schema.slice(0, springEnd).split('').fill('#', i, springEnd);
    toAdd[springEnd] = remainingSprings.length ? '.' : '';
    toAdd = toAdd.join('').replaceAll('?', '.');

    const memoryHash = (remainingSchema + JSON.stringify(remainingSprings));
    let subResult;
    if (memory[memoryHash] !== undefined) {
      subResult = memory[memoryHash];
    } else {
      subResult = calculateArrangements(
        remainingSchema, remainingSprings, basePath.replaceAll('?', '.') + toAdd, memory
      );
      memory[memoryHash] = subResult;
    }
    currentArrangements += subResult;
  }

  return currentArrangements;
}

let result = 0;
for (const arrangement of springArrangements.split('\n')) {
  let [schema, springSizes] = arrangement.split(' ');
  if (springSizes === undefined) {
    continue;
  }

  // springSizes = springSizes.split(',').map(Number);
  springSizes = Array(5).fill(springSizes).join(',').split(',').map(Number);
  schema = Array(5).fill(schema).join('?');
  console.log(schema, springSizes)
  let temp = calculateArrangements(schema, springSizes);
  console.log(temp)
  result += temp;
}

//7195 PERFECT
//33992866292225

console.log(result)

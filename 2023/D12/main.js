const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const springArrangements = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');

const findSpringEnd = (schema, springSize) => {
  const closestDot = schema.indexOf('.');
  if (closestDot >= 0 && closestDot < springSize) {
    return -1;
  }

  const isEndValid =
    ['.', '?'].includes(schema[springSize])
    || springSize === schema.length
  ;
  return isEndValid ? springSize : -1;
}

function calculateArrangements(schema, springSizes, memory = new Map()) {
  if (springSizes.length === 0) {
    return schema.indexOf('#') === -1 ? 1 : 0;
  }

  const currentSpringSize = springSizes[0];
  let currentArrangements = 0;
  let springBank = currentSpringSize;

  for (let i = 0; i < schema.length; i++) {
    if (springBank < currentSpringSize) {
      break;
    }
    const springEnd = findSpringEnd(schema.slice(i), springBank)
    if (schema[i] === '#') {
      springBank -= 1;
    }
    if (springEnd === -1) {
      continue;
    }

    const remainingSprings = springSizes.slice(1);
    const remainingSchema = schema.slice(springEnd + i + 1); //slice is stupid
    const memoryHash = (remainingSchema + JSON.stringify(remainingSprings));
    if (memory[memoryHash] === undefined) {
      memory[memoryHash] = calculateArrangements(
        remainingSchema, remainingSprings, memory
      );
    }

    currentArrangements += memory[memoryHash];
  }

  return currentArrangements;
}

let [result, expandedResult] = [0, 0];
for (const arrangement of springArrangements.split('\n')) {
  let [schema, springSizes] = arrangement.split(' ');
  result += calculateArrangements(schema, springSizes.split(',').map(Number));

  const expandedSpringSizes = Array(5).fill(springSizes).join(',').split(',').map(Number);
  const expandedSchema = Array(5).fill(schema).join('?');
  expandedResult += calculateArrangements(expandedSchema, expandedSpringSizes);
}

//7195 p1
//33992866292225 p2
console.log('P1: ', result)
console.log('P2: ', expandedResult)

const fs = require('fs');

function main() {
  // const commands = fs.readFileSync('./test.txt', 'utf8').split('\n');
  const commands = fs.readFileSync('./tests.txt', 'utf8').split('\n');
  // const commands = fs.readFileSync('./input.txt', 'utf8').split('\n');
  let lagoon = new Map();
  let curPoint = {
    x: 0,
    y: 0,
  };

  for (const command of commands) {
    if (command === '') continue;
    let [dir, steps, color] = command.split(' ');

    let i = 0;
    while (i < steps) {
      if (dir === 'U') {
        curPoint.y -= 1;
      } else if (dir === 'D') {
        curPoint.y += 1;
      } else if (dir === 'L') {
        curPoint.x -= 1;
      } else if (dir === 'R') {
        curPoint.x += 1;
      }

      if (lagoon.has(curPoint.y)) {
        const currentEdge = lagoon.get(curPoint.y);

        lagoon.set(curPoint.y, [Math.min(currentEdge[0], curPoint.x), Math.max(currentEdge[1], curPoint.x)]);
      } else {
        lagoon.set(curPoint.y, [curPoint.x, curPoint.x]);
      }

      i++;
    }
  }

  let area = 0;
  let pool = [];
  console.log(lagoon)
  // for (const line in lagoon) {
  //   // area += Number(line[1] + 1) - Number(line[0]);
  // }

  lagoon.forEach((line, key, map) => {
    let i = 0;
    let displayLine = '';
    console.log(line)
    while (i <= line[1]) {
      if (i < line[0]) {
        displayLine += '.';
      } else {
        displayLine += '#';
        area += 1;
      }
      i++;
    }
    pool.push(displayLine);

  })


  // console.log(pool.join('\n'))
  console.log(`Total area is: ${area}`)
  // too high 39340
  // too low 4931
}

main();

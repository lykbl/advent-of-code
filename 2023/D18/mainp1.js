const fs = require('fs');

function calculateLavaLagoonArea(input) {
  const instructions = input.trim().split('\n').map(line => {
    const [direction, distance, color] = line.split(' ');
    return { direction, distance: parseInt(distance), color: color.slice(1, -1) };
  });

  let x = 0, y = 0;
  let perimeter = 0;
  let area = 0;

  for (const { direction, distance } of instructions) {
    let dx = 0, dy = 0;

    switch (direction) {
      case 'U': dy = -distance; break;
      case 'D': dy = distance; break;
      case 'L': dx = -distance; break;
      case 'R': dx = distance; break;
    }

    // Calculate area using Shoelace formula
    area += x * dy - y * dx;

    // Update perimeter
    perimeter += distance;

    // Move to next point
    x += dx;
    y += dy;
  }

  // Finalize area calculation
  area = Math.abs(area) / 2;

  // Add half of the perimeter to account for the trench width
  const totalArea = area + perimeter / 2 + 1; // +1 to account for the starting point

  return Math.floor(totalArea);
}

// Read input from file
const input = fs.readFileSync('input.txt', 'utf8');
// const input = fs.readFileSync('test.txt', 'utf8');

// Calculate and print the result
console.log(calculateLavaLagoonArea(input));

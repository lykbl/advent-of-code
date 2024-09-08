const fs = require('fs');

function calculateLavaLagoonArea(input) {
  const instructions = input.trim().split('\n').map(line => {
    const [direction, distance, color] = line.split(' ');
    return { direction, distance: parseInt(distance), color: color.slice(1, -1) };
  });

  let x = 0, y = 0;
  let perimeter = 0;
  let area = 0;

  for (const { color } of instructions) {
    const [distanceEncoded, directionEncoded] = [color.slice(1, -1), color.slice(-1)]
    const distance = Number(`0x${distanceEncoded}`);
    console.log(`Doing: ${directionEncoded}:${distance}`);
    let dx = 0, dy = 0;

    // switch (direction) {
    switch (directionEncoded) {
      case '3': dy = -distance; break;
      case '1': dy = distance; break;
      case '2': dx = -distance; break;
      case '0': dx = distance; break;
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

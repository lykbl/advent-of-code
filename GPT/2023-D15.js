function runHashAlgorithm(input) {
  let current_value = 0;
  const steps = input.split(',');

  for (const step of steps) {
    let charCode = 0;

    for (let i = 0; i < step.length; i++) {
      charCode += step.charCodeAt(i);
    }

    current_value += charCode;
    current_value *= 17;
    current_value %= 256;
  }

  return current_value;
}

function runHashmap(initializationSequence) {
  const boxes = new Array(256).fill([]); // Initialize 256 boxes

  const steps = initializationSequence.split(',');

  for (const step of steps) {
    const [label, operation] = step.split(/[=-]/);
    const boxIndex = runHashAlgorithm(label);

    if (operation === '') {
      // Insert lens
      const focalLengthMatch = step.match(/\d+/);
      const focalLength = focalLengthMatch ? parseInt(focalLengthMatch[0]) : 1;
      const lens = `${ label } ${ focalLength }`;

      const box = boxes[boxIndex];
      const existingLensIndex = box.findIndex((l) => l.startsWith(label));

      if (existingLensIndex !== -1) {
        // Replace existing lens
        box[existingLensIndex] = lens;
      } else {
        // Add new lens
        box.unshift(lens); // Add lens to the front
      }
    } else if (operation === '-') {
      // Remove lens
      const lensIndex = boxes[boxIndex]?.findIndex((l) => l.startsWith(label));

      if (lensIndex !== -1) {
        // Lens found, remove it
        boxes[boxIndex].splice(lensIndex, 1);
      }
    }
  }

  // Calculate focusing power
  let totalFocusingPower = 0;

  boxes.forEach((box, boxIndex) => {
    box.forEach((lens, slotIndex) => {
      const [, focalLength] = lens.split(' ');
      totalFocusingPower += (boxIndex + 1) * (slotIndex + 1) * parseInt(focalLength);
    });
  });

  return totalFocusingPower;
}

const initializationSequence = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";
const result = runHashmap(initializationSequence);
console.log(result);

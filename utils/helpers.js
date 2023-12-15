const isDigit = (character) => !isNaN(parseInt(character));

function* extractLineNumbers(line, terminationCharacter = null) {
  let numberAcc = '';
  for (let i = 0; i < line.length; i++) {
    let character = line[i];
    if (character === terminationCharacter) {
      return null;
    }

    if (isDigit(character)) {
      numberAcc += character;
    } else {
      if (numberAcc.length) {
        yield parseInt(numberAcc);
        numberAcc = '';
      }
    }
  }

  if (numberAcc.length) {
    yield parseInt(numberAcc);
  }

  return null;
}

module.exports = {
  isDigit,
  extractLineNumbers,
};
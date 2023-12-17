const isDigit = (character) => !isNaN(parseInt(character));

function* extractLineNumbers(line, terminationCharacter = null, itemsCount = 1, ignoreCharacters = [], acceptedValues = /\d/) {
  let numberAcc = '';
  const numbers = [];
  for (let i = 0; i < line.length; i++) {
    let character = line[i];
    if (character === terminationCharacter) {
      return null;
    }

    const isAccepted = acceptedValues.test(character);
    if (isAccepted) {
      numberAcc += character;
    }

    if ((!isAccepted && !ignoreCharacters.includes(character) || i === line.length - 1) && numberAcc.length) {
      numbers.push(parseInt(numberAcc));
      numberAcc = '';

      while (numbers.length >= itemsCount) {
        yield itemsCount === 1
          ? numbers.shift()
          : numbers.slice(0, itemsCount) && numbers.splice(0, itemsCount)
      }
    }
  }

  return null;
}

module.exports = {
  isDigit,
  extractLineNumbers,
};
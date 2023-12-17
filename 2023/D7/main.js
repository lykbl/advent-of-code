const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
// const { extractLineNumbers } = require("../../utils/helpers");
const BinaryHeap = require("../../utils/binaryHeap");

const HAND_TYPES = {
  'High Card': 0,
  'One Pair': 2,
  'Two Pairs': 3,
  'Three of a Kind': 4,
  'Full House': 5,
  'Four of a Kind': 6,
  'Five of a Kind': 7,
}

const CARD_WEIGHTS = {
  'A': 14,
  'K': 13,
  'Q': 12,
  // 'J': 10,
  'J': 1,
  'T': 10,
  '9': 9,
  '8': 8,
  '7': 7,
  '6': 6,
  '5': 5,
  '4': 4,
  '3': 3,
  '2': 2,
}

const rs = readline.createInterface({
  input: fileStream,
  crlfDelay: Infinity,
});

class Hand {
  constructor(weight, type, bid, hand) {
    this._weight = weight;
    this._type = type;
    this._bid = parseInt(bid);
    this._hand = hand;
  }

  weight = () => this._weight;
  type = () => this._type;
  bid = () => this._bid;

  toString = () => `${this._hand} ${this._weight} ${this._type} ${this._bid}`;
}

const compareHands = (nodeA, nodeB) => {
  const handA = nodeA.value;
  const handB = nodeB.value;
  if (handA.type() > handB.type()) {
    return 1;
  }
  if (handA.type() < handB.type()) {
    return -1;
  }

  return handA.weight() - handB.weight();
}

const hands = new BinaryHeap(compareHands);
rs.on('line', (line) => {
  const [hand, bid] = line.split(' ');
  const cardsCountMap = new Map();
  let type = HAND_TYPES['High Card'];
  let [handWeight, wildCardsCount] = [0, 0];

  for (let i = 0; i < hand.length; i++) {
    const card = hand[i];
    const cardWeight = CARD_WEIGHTS[card];
    const cardCount = cardsCountMap.get(card) + 1 || 1;

    cardsCountMap.set(card, cardCount);
    handWeight += cardWeight * Math.pow(Object.keys(CARD_WEIGHTS).length + 1, hand.length - i);

    if (card === 'J') {
      wildCardsCount++;
      continue;
    }

    if (type === HAND_TYPES['High Card'] && cardCount === 2) {
      type = HAND_TYPES['One Pair'];
    } else if (type === HAND_TYPES['One Pair'] && cardCount === 2) {
      type = HAND_TYPES['Two Pairs'];
    } else if (type === HAND_TYPES['One Pair'] && cardCount === 3) {
      type = HAND_TYPES['Three of a Kind'];
    } else if ((type === HAND_TYPES['Two Pairs'] && cardCount === 3) || (type === HAND_TYPES['Three of a Kind'] && cardCount === 2)) {
      type = HAND_TYPES['Full House'];
    } else if (type === HAND_TYPES['Three of a Kind'] && cardCount === 4) {
      type = HAND_TYPES['Four of a Kind'];
    } else if (type === HAND_TYPES['Four of a Kind'] && cardCount === 5) {
      type = HAND_TYPES['Five of a Kind'];
      wildCardsCount = 0;
    }
  }

  while (wildCardsCount) {
    type = Math.min(type + 2, HAND_TYPES['Five of a Kind'])
    wildCardsCount--;
  }

  hands.insert(new Hand(handWeight, type, bid, hand));
});

rs.on('close', () => {
  let totalWinnings = 0;
  let i = 1;
  while (hands.size()) {
    totalWinnings += i * hands.extractMinimum().value.bid();
    i++;
  }

  console.log(totalWinnings)
})

//248836197 p1
//251195607 p2

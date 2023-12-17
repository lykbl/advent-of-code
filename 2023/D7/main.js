const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const readline = require('readline');
const fileStream = fs.createReadStream(isTest ? './test.txt' : './input.txt');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
// const { extractLineNumbers } = require("../../utils/helpers");
const BinaryHeap = require("../../utils/binaryHeap");

const HAND_TYPES = {
  'High Card': 1,
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
  'J': 11,
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
  constructor(weight, type, bid) {
    this._weight = weight;
    this._type = HAND_TYPES[type];
    this._bid = parseInt(bid);
  }

  weight = () => this._weight;
  type = () => this._type;
  bid = () => this._bid;
}

const hands = [];
rs.on('line', (line) => {
  const [hand, bid] = line.split(' ');
  const cardsCountMap = new Map();
  let type = 'High Card';
  let handWeight = 0;

  for (let i = 0; i < hand.length; i++) {
    const card = hand[i];
    const cardWeight = CARD_WEIGHTS[card];
    const cardCount = cardsCountMap.get(card) || 1;

    cardsCountMap.set(card, cardCount + 1);

    if (type === 'High Card' && cardCount === 2) {
      type = 'One Pair';
    } else if (type === 'One Pair' && cardCount === 2) {
      type = 'Two Pairs';
    } else if (type === 'One Pair' && cardCount === 3) {
      type = 'Three of a Kind';
    } else if ((type === 'Two Pairs' && cardCount === 3) || (type === 'Three of a Kind' && cardCount === 2)) {
      type = 'Full House';
    } else if (type === 'Three of a Kind' && cardCount === 4) {
      type = 'Four of a Kind';
    } else if (type === 'Four of a Kind' && cardCount === 5) {
      type = 'Five of a Kind';
    }

    handWeight += cardWeight * Math.pow(Object.keys(CARD_WEIGHTS).length, hand.length - i);
  }

  hands.push(new Hand(handWeight, type, bid, hand));
});


rs.on('close', () => {
  hands.sort((handA, handB) => {
    if (handA.type() > handB.type()) {
      return 1;
    }
    if (handA.type() < handB.type()) {
      return -1;
    }

    return handA.weight() - handB.weight();
  });

  let totalWinnings = 0;
  for (let i = 0; i < hands.length; i++) {
    totalWinnings += (i + 1) * hands[i].bid();
  }

  // console.log(hands)
  console.log(totalWinnings)
})

//248836197
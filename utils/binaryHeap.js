'use strict';

/**
 * Creates a binary heap.
 *
 * @constructor
 * @param {function} customCompare An optional custom node comparison
 * function.
 */
var BinaryHeap = function (customCompare = null) {
  this.list = [];

  if (customCompare) {
    this.compare = customCompare;
  }
};

BinaryHeap.prototype.buildHeap = function (...values) {
  const nodeArray = [];
  for (let i = 0; i < values.length; i++) {
    nodeArray.push(new Node(values[i]));
  }

  buildHeapFromNodeArray(this, nodeArray);
};

BinaryHeap.prototype.clear = function () {
  this.list.length = 0;
};

BinaryHeap.prototype.extractMinimum = function () {
  if (!this.list.length) {
    return undefined;
  }
  if (this.list.length === 1) {
    return this.list.shift();
  }
  const min = this.list[0];
  this.list[0] = this.list.pop();
  heapify(this, 0);
  return min;
};

BinaryHeap.prototype.findMinimum = function () {
  return this.isEmpty() ? undefined : this.list[0];
};

BinaryHeap.prototype.insert = function (value) {
  const node = new Node(value);
  let i = this.list.length;
  this.list.push(node);
  let parent = getParent(i);
  while (typeof parent !== 'undefined' && this.compare(this.list[i], this.list[parent]) < 0) {
    swap(this.list, i, parent);
    i = parent;
    parent = getParent(i);
  }
  return node;
};

BinaryHeap.prototype.isEmpty = function () {
  return !this.list.length;
};

BinaryHeap.prototype.size = function () {
  return this.list.length;
};

// BinaryHeap.prototype.union = function (otherHeap) {
//   const array = this.list.concat(otherHeap.list);
//   buildHeapFromNodeArray(this, array);
// };

BinaryHeap.prototype.compare = function (a, b) {
  if (a.value > b.value) {
    return 1;
  }
  if (a.value < b.value) {
    return -1;
  }
  return 0;
};

function heapify(heap, i) {
  const l = getLeft(i);
  const r = getRight(i);
  let smallest = i;
  if (l < heap.list.length && heap.compare(heap.list[l], heap.list[i]) < 0) {
    smallest = l;
  }
  if (r < heap.list.length && heap.compare(heap.list[r], heap.list[smallest]) < 0) {
    smallest = r;
  }
  if (smallest !== i) {
    swap(heap.list, i, smallest);
    heapify(heap, smallest);
  }
}

function buildHeapFromNodeArray(heap, nodeArray) {
  heap.list = nodeArray;
  for (let i = Math.floor(heap.list.length / 2); i >= 0; i--) {
    heapify(heap, i);
  }
}

function swap(array, a, b) {
  [array[a], array[b]] = [array[b], array[a]];
}

function getParent(i) {
  if (i === 0) {
    return undefined;
  }
  return Math.floor((i - 1) / 2);
}

function getLeft(i) {
  return 2 * i + 1;
}

function getRight(i) {
  return 2 * i + 2;
}

function Node(value) {
  this.value = value;
}

module.exports = BinaryHeap;
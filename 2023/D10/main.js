const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');
const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
const maze = input;

const TILE_TYPES = {
  '|': 'VERTICAL',
  '-': 'HORIZONTAL',
  'L': 'NORTH-EAST',
  'J': 'NORTH-WEST',
  '7': 'SOUTH-WEST',
  'F': 'SOUTH-EAST',
  '.': 'GROUND',
  'S': 'START', // animal
};
const tileConnections = (tileIndex, maze, rowSize) => {
  const type = TILE_TYPES[maze[tileIndex]];
  const canConnectOn = [];

  if (type === 'VERTICAL') {
    canConnectOn.push(tileIndex - rowSize, tileIndex + rowSize);
  } else if (type === 'HORIZONTAL') {
    canConnectOn.push(tileIndex - 1, tileIndex + 1);
  } else if (type === 'NORTH-EAST') {
    canConnectOn.push(tileIndex - rowSize, tileIndex + 1);
  } else if (type === 'NORTH-WEST') {
    canConnectOn.push(tileIndex - rowSize, tileIndex - 1);
  } else if (type === 'SOUTH-EAST') {
    canConnectOn.push(tileIndex + rowSize, tileIndex + 1);
  } else if (type === 'SOUTH-WEST') {
    canConnectOn.push(tileIndex + rowSize, tileIndex - 1);
  } else if (type === 'START') {
    canConnectOn.push(tileIndex - rowSize, tileIndex + rowSize, tileIndex - 1, tileIndex + 1)
  }

  return canConnectOn;
}
const getConnectedTiles = (maze, currentTileIndex, rowSize) => {
  const currentTileConnections = tileConnections(currentTileIndex, maze, rowSize);
  const allowedConnections = [];

  for (let currentTileConnection of currentTileConnections) {
    const tileToConnectConnections = tileConnections(currentTileConnection, maze, rowSize);
    if (tileToConnectConnections.includes(currentTileIndex)) {
      allowedConnections.push(currentTileConnection);
    }
  }

  return allowedConnections;
}

class Graph {
  constructor() {
    this.adjacencyList = {};
    this.vertices = [];
    this.rootNode = null;
  }

  addNode(node) {
    if (!this.adjacencyList[node.key]) {
      this.adjacencyList[node.key] = [];
      this.vertices.push(node);

      if (node.value === 'S') {
        this.rootNode = node.key;
      }
    }
  }

  addEdge(key1, key2) {
    this.adjacencyList[key1].push(key2);
    this.adjacencyList[key1].push(key2);
  }

  bfsMaxDistance(startNode) {
    const queue = [];
    const visited = new Set();
    let maxDistance = 0;
    queue.push({ node: startNode, distance: 0 });
    visited.add(startNode);

    while (queue.length > 0) {
      const { node, distance } = queue.shift();
      maxDistance = Math.max(maxDistance, distance);

      // Explore neighbors
      for (const neighbor of this.adjacencyList[node]) {
        if (!visited.has(neighbor)) {
          queue.push({ node: neighbor, distance: distance + 1 });
          visited.add(neighbor);
        }
      }
    }

    return maxDistance;
  }
}

const rowSize = maze.indexOf('\n') + 1;
const graph = new Graph();

for (let i = 0; i < maze.length; i++) {
  const connectedTiles = getConnectedTiles(maze, i, rowSize);
  if (connectedTiles.length === 0) continue;

  graph.addNode({ key: i, value: maze[i] });
  for (let connectedTile of connectedTiles) {
    graph.addNode({ key: connectedTile, value: maze[connectedTile] });
    graph.addEdge(i, connectedTile);
  }
}
//6907 p1 correct

console.log(graph.bfsMaxDistance(graph.rootNode));
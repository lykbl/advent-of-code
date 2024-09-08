const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');

// const input = fs.readFileSync(isTest ? './test.txt' : './input.txt', 'utf8');
const input = fs.readFileSync(isTest ? './test-small.txt' : './input.txt', 'utf8');

class PriorityQueue {
  constructor() {
    this.queue = [];
  }

  enqueue(node, priority, meta) {
    this.queue.push({ node, priority, meta });
    this.queue.sort((a, b) => a.priority - b.priority);
  }

  dequeue() {
    return this.queue.shift();
  }

  isEmpty() {
    return this.queue.length === 0;
  }
}

class Graph {
  constructor() {
    this.vertices = {};
  }

  addVertex({ label, heatLoss }) {
    if (!this.vertices[label]) {
      this.vertices[label] = {
        edges: [],
        heatLoss: heatLoss,
      };
    }
  }

  addEdge(source, destination) {
    this.vertices[source].edges.push({ node: destination });
  }

  getShortestPath(start, end, visited = new Map(), totalHeatLoss = 0, samePathRepeats = 3, cameFrom = null) {
    if (start === end) {
      visited.set(start, true);

      return this.vertices[start].heatLoss
    }

    const possibleMoves = this.vertices[start].edges.filter(({ node }) => !visited.has(node));
    if (possibleMoves.length === 0 || samePathRepeats === 0) {
      return -1;
    }

    const nextStepHeatLosses = possibleMoves.map(({ node }) => {
      const nextVisited = new Map(visited);
      nextVisited.set(start, true);
      const [currentX, currentY] = start.split('_').map(Number);
      const [nextX, nextY] = node.split('_').map(Number);
      const dimension = currentX === nextX ? 'vertical' : 'horizontal';
      const direction = currentX < nextX || currentY < nextY ? 'forward' : 'backward';
      const nextDirection = dimension + direction;
      const sameDirection = nextDirection === cameFrom;
      return this.getShortestPath(node, end, nextVisited, totalHeatLoss + this.vertices[node].heatLoss, sameDirection ? samePathRepeats - 1 : 3);
    })

    const shortestPath = Math.min(...nextStepHeatLosses.filter(path => path !== -1));

    console.log(shortestPath)

    return shortestPath + this.vertices[start].heatLoss
  }
}

function gridToGraph(grid) {
  const graph = new Graph();
  for (let y = 0; y < grid.length; y++) {
    const row = grid[y];
    for (let x = 0; x < row.length; x++) {
      const heatLoss = grid[y][x];
      graph.addVertex({label: `${ x }_${ y }`, heatLoss: Number(heatLoss) });
      const directions = [
        { coords: [x + 1, y] },
        { coords: [x, y + 1] },
        { coords: [x, y - 1] },
      ];
      for (let { coords: [nextX, nextY] } of directions) {
        const nextHeatLoss = grid[nextY] && grid[nextY][nextX];
        if (nextHeatLoss === undefined) {
          continue;
        }
        graph.addVertex({label: `${ nextX }_${ nextY }`, heatLoss: Number(nextHeatLoss) });
        graph.addEdge(`${ x }_${ y }`, `${ nextX }_${ nextY }`);
      }
    }
  }

  return graph;
}

const grid = input.split('\n').map(row => row.split(''));
const directedGraph = gridToGraph(grid);
console.log(directedGraph)

const shortestPath = directedGraph.getShortestPath('0_0', `${ grid[0].length - 1 }_${ grid.length - 1 }`);

console.log(shortestPath);
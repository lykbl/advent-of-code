class Graph {
  constructor() {
    this.nodes = [];
    this.edges = [];
  }

  addNode(node) {
    this.nodes.push(node);
  }

  addEdge(node1, node2) {
    this.edges.push([node1, node2]);
  }
}
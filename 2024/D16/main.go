package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Step struct {
  curNode *Node
  steps int
  turns int
}

type Node struct {
  adj map[string]int
  pos [2]int
}

type Graph struct {
  nodes map[string]*Node
}

func BuildGraph(grid [][]rune) Graph {
  graph := Graph { nodes: make(map[string]*Node, 0) }
  dirs := [][2]int{
    { 0, -1 },
    { 1, 0 },
    { 0, 1 },
    { -1, 0 },
  }

  for y := range grid {
    for x := range grid[y] {
      if grid[y][x] == '#' {
        continue
      }

      n := 0
      nDirs := []int{}
      for dirI, dir := range dirs {
        if grid[y + dir[1]][x + dir[0]] == '.' {
          n++
          nDirs = append(nDirs, dirI)
        }
      }

      if n >= 2 {
        if n == 2 && (nDirs[0] % 2 == nDirs[1] % 2) {
          continue
        }
        graph.nodes[fmt.Sprintf("%d_%d", x, y)] = &Node{ adj: make(map[string]int, 0), pos: [2]int{x, y} }
      }
    }
  }

  for _, node := range graph.nodes {
    for _, dir := range dirs {
      newPos := [2]int{node.pos[0] + dir[0], node.pos[1] + dir[1]}
      stepsTaken := 1
      for grid[newPos[1]][newPos[0]] == '.' {
        if adjNode, exists := graph.nodes[fmt.Sprintf("%d_%d", newPos[0], newPos[1])]; exists {
          node.adj[fmt.Sprintf("%d_%d", adjNode.pos[0], adjNode.pos[1])] = stepsTaken
        }

        newPos[0] += dir[0]
        newPos[1] += dir[1]
        stepsTaken++
      }
    }
  }

  return graph
}

func main() {
  file, err := os.Open("input.txt")
  // file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal("FAIL")
  }

  scanner := bufio.NewScanner(file)
  grid := make([][]rune, 0)
  var start, end [2]int
  y := 0
  for scanner.Scan() {
    line := scanner.Text()

    x := 0
    grid = append(grid, make([]rune, 0))
    for _, char := range line {
      if char == 'S' {
        start = [2]int{x, y}
      }
      if char == 'E' {
        end = [2]int{x, y}
      }
      grid[y] = append(grid[y], char)
      x ++
    }
    y++
  }
  grid[end[1]][end[0]] = '.'
  grid[start[1]][start[0]] = '.'

  graph := BuildGraph(grid)
  steps := make([]*Step, 0)
  startNode, _ := graph.nodes[fmt.Sprintf("%d_%d", start[0], start[1])]
  steps = append(steps, &Step{
    curNode: startNode,
    steps: 0,
    turns: 1,
  })
  // N E S W
  // 0 1 2 3

  distances := make(map[string]float64, 0)
  for key := range graph.nodes {
    distances[key] = math.MaxFloat64
  }

  visited := make(map[string]struct{}, 0)
  for len(steps) > 0 {
    currentStep := steps[0]
    steps = steps[1:]
    currentNodeKey := fmt.Sprintf("%d_%d", currentStep.curNode.pos[0], currentStep.curNode.pos[1])

    if currentStep.curNode.pos[0] == end[0] && currentStep.curNode.pos[1] == end[1] {
      currentStep.turns--
    }
    nextNodeDistance := float64(currentStep.steps + 1000 * currentStep.turns)
    currentDistance, _ := distances[currentNodeKey]
    distances[currentNodeKey] = math.Min(nextNodeDistance, currentDistance)
    fmt.Printf("Doing now: %s\n", currentNodeKey)
    if _, alreadyVisited := visited[currentNodeKey]; alreadyVisited {
      continue
    }

    visited[currentNodeKey] = struct{}{}
    for nextNodeKey, nextNodeSteps := range currentStep.curNode.adj {
      steps = append(steps, &Step {
        curNode: graph.nodes[nextNodeKey],
        steps: currentStep.steps + nextNodeSteps,
        turns: currentStep.turns + 1,
      })
    }
  }

  p1 := int(distances[fmt.Sprintf("%d_%d", end[0], end[1])])
  fmt.Printf("P1: %d\n", p1)
}


package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
  x int
  y int
}

type BuildStep struct {
  FromNodeKey string
  CurrentTile Point
  StepsTaken int
}

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatalf("File not found")
  }
  defer file.Close()

  grid := make([][]rune, 0)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    grid = append(grid, []rune(scanner.Text()))
  }

  graph := make(map[string]map[string]int, 0)
  graph["1_0"] = make(map[string]int, 0)
  graph[fmt.Sprintf("%d_%d", len(grid[0]) - 2, len(grid[0]) - 1)] = make(map[string]int, 0)
  queue := []BuildStep{
    {FromNodeKey: "1_0", CurrentTile: Point { x: 1, y: 0 }, StepsTaken: 0},
  }
  targetCell := Point { x: len(grid[0]) - 2, y: len(grid) - 1  }

  visitedCells := make(map[string]struct{})
  for len(queue) > 0 {
    current := queue[len(queue) - 1]
    queue = queue[0:len(queue) - 1]
    x, y := current.CurrentTile.x, current.CurrentTile.y
    visitedCells[fmt.Sprintf("%d_%d", x, y)] = struct{}{}
    currentTileRune := grid[y][x]
    if x == targetCell.x && y == targetCell.y {
      log.Print("ADDED LAST")
      graph[current.FromNodeKey][fmt.Sprintf("%d_%d", x, y)] = current.StepsTaken
    }

    if currentTileRune == '.' {
      for _, nextTile := range []Point {
        {x: x + 1, y: y},
        {x: x - 1, y: y},
        {x: x, y: y + 1},
        {x: x, y: y - 1},
      } {
        if nextTile.x < 0 || nextTile.y < 0 || nextTile.x >= len(grid[0]) || nextTile.y >= len(grid) {
          continue
        }
        if grid[nextTile.y][nextTile.x] == '#' {
          continue
        }
        nextTileKey := fmt.Sprintf("%d_%d", nextTile.x, nextTile.y)
        if _, visited := visitedCells[nextTileKey]; visited {
          continue
        }

        queue = append(queue, BuildStep{FromNodeKey: current.FromNodeKey, CurrentTile: Point { x: nextTile.x, y: nextTile.y }, StepsTaken: current.StepsTaken + 1})
      }
    } else {
        var newNodeStart Point
        if currentTileRune == 'v' {
          newNodeStart = Point { x: x, y: y + 1 }
        }
        if currentTileRune == '>' {
          newNodeStart = Point { x: x + 1, y: y }
        }
        newNodeKey := fmt.Sprintf("%d_%d", newNodeStart.x, newNodeStart.y)
        if err != nil {
          log.Fatal("Keyt failed")
        }

        graph[current.FromNodeKey][newNodeKey] = current.StepsTaken + 1
        if _, exists := graph[newNodeKey]; exists {
          graph[newNodeKey][current.FromNodeKey] = current.StepsTaken + 1
          continue
        }

        graph[newNodeKey] = make(map[string]int, 0)
        graph[newNodeKey][current.FromNodeKey] = current.StepsTaken + 1
        for _, nextTile := range []Point {
          {x: newNodeStart.x + 1, y: newNodeStart.y},
          {x: newNodeStart.x, y: newNodeStart.y + 1},
        } {
          if nextTile.x < 0 || nextTile.y < 0 || nextTile.x >= len(grid[0]) || nextTile.y >= len(grid) {
            continue
          }
          if grid[nextTile.y][nextTile.x] == '#' {
            continue
          }

          if grid[nextTile.y][nextTile.x] == '>' {
            visitedCells[fmt.Sprintf("%d_%d", nextTile.x, nextTile.y)] = struct{}{}
            queue = append(queue, BuildStep{FromNodeKey: newNodeKey, CurrentTile: Point { x: nextTile.x + 1, y: nextTile.y }, StepsTaken: 1})
          }
          if grid[nextTile.y][nextTile.x] == 'v' {
            visitedCells[fmt.Sprintf("%d_%d", nextTile.x, nextTile.y)] = struct{}{}
            queue = append(queue, BuildStep{FromNodeKey: newNodeKey, CurrentTile: Point { x: nextTile.x, y: nextTile.y + 1}, StepsTaken: 1})
          }
        }
      }
  }

  for nodeKey, neighbours := range graph {
    result := strings.Split(nodeKey, "_")
    x, _ := strconv.Atoi(result[0])
    y, _ := strconv.Atoi(result[1])

    fmt.Printf("%d %d: [", y + 1, x + 1)
    for neighbourKey := range neighbours {
      result := strings.Split(neighbourKey, "_")
      nX, _ := strconv.Atoi(result[0])
      nY, _ := strconv.Atoi(result[1])
      fmt.Printf("%d %d, ", nY + 1, nX + 1)
    }
    fmt.Print("]\n")
  }

  nodesQueue := []struct{CurrentNodeKey string; StepsTaken int; VisitedNodes map[string]struct{}}{
    {CurrentNodeKey: "1_0", StepsTaken: 0, VisitedNodes: make(map[string]struct{})},
  }
  
  maxSteps := 0
  for len(nodesQueue) > 0 {
    currentStep := nodesQueue[len(nodesQueue) - 1]
    nodesQueue = nodesQueue[0:len(nodesQueue) - 1]
    // log.Printf("New vis: %v", currentStep.VisitedNodes)

    if currentStep.CurrentNodeKey == fmt.Sprintf("%d_%d", targetCell.x, targetCell.y) {
      maxSteps = max(maxSteps, currentStep.StepsTaken)
    }

    newVisited := make(map[string]struct{})
    for t := range currentStep.VisitedNodes {
      newVisited[t] = struct{}{}
    }
    newVisited[currentStep.CurrentNodeKey] = struct{}{}
    for nextNodeKey, stepsToTake := range graph[currentStep.CurrentNodeKey] {
      if _, alreadyVisited := newVisited[nextNodeKey]; alreadyVisited {
        continue
      }

      nodesQueue = append(nodesQueue, struct{CurrentNodeKey string; StepsTaken int; VisitedNodes map[string]struct{}}{
        CurrentNodeKey: nextNodeKey,
        StepsTaken: currentStep.StepsTaken + stepsToTake + 1,
        VisitedNodes: newVisited,
      })
    }
  }

  log.Printf("Result: %d", maxSteps - 1)
}

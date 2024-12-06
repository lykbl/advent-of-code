package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vec2 struct {
  x int
  y int
}

type Guard struct {
  pos Vec2
  direction rune
}

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal("File failed")
  }

  grid := make([][]rune, 0)
  scanner := bufio.NewScanner(file)
  
  var guard *Guard
  y := 0
  for scanner.Scan() {
    line := scanner.Text()
    grid = append(grid, []rune(line))

    if guard == nil {
      for x, char := range line {
        if char != '.' && char != '#' {
          guard = &Guard { pos: Vec2 { x, y }, direction: char }
        }
      }
    }
    y++
  }
  if scanner.Err() != nil {
    log.Fatalf("SCANNER FAILED: %s", scanner.Err().Error())
  }

  grid[guard.pos.y][guard.pos.x] = '.'
  currentDir := guard.direction
  currentPos := guard.pos
  visitedCells := make(map[string]struct{})
  visitedCells[fmt.Sprintf("%d_%d", guard.pos.y, guard.pos.x)] = struct{}{}
  // loopRules := make(map[string]struct{})
  // wallCount := 3
  for {
    var nextPos Vec2
    switch currentDir {
    case '^':
      nextPos = Vec2 { y: currentPos.y - 1, x: currentPos.x }
    case 'v':
      nextPos = Vec2 { y: currentPos.y + 1, x : currentPos.x }
    case '>':
      nextPos = Vec2 { y: currentPos.y, x: currentPos.x + 1 }
    case '<':
      nextPos = Vec2 { y: currentPos.y, x: currentPos.x - 1 }
    default:
      log.Fatalf("Unexpected char: %c", currentDir)
    }

    if !IsValidPos(grid, nextPos) {
      break
    }

    nextTile := grid[nextPos.y][nextPos.x]
    nextTileKey := fmt.Sprintf("%d_%d", nextPos.y, nextPos.x)
    log.Printf("Next: %c (%d_%d)", nextTile, nextPos.y, nextPos.x)
    if nextTile == '.' {
      if _, alreadyVisited := visitedCells[nextTileKey]; !alreadyVisited {
        visitedCells[nextTileKey] = struct{}{}
      }
      currentPos = Vec2 { x: nextPos.x, y: nextPos.y }
    }
    if nextTile == '#' {
      // if currentDir == '^' {
      //   ruleKey = fmt.Sprintf("x%d_%c", nextPos.x, '<')
      // }
      // if currentDir == 'v' {
      //   ruleKey = fmt.Sprintf("x%d_%c", nextPos.x, '>')
      // }
      // if currentDir == '<' {
      //   ruleKey = fmt.Sprintf("y%d_%c", nextPos.y, 'v')
      // }
      // if currentDir == '>' {
      //   ruleKey = fmt.Sprintf("y%d_%c", nextPos.y, '^')
      // }

      var newDirection rune
      switch currentDir {
      case '^':
        newDirection = '>'
      case '>':
        newDirection = 'v'
      case 'v':
        newDirection = '<'
      case '<':
        newDirection = '^'
      }
      currentDir = newDirection
    } 
  }

  p1 := len(visitedCells)
  log.Printf("P1: %d", p1)

  // p2 := len(obstacleCells)
  // log.Printf("P2: %d", p2)
}

func IsValidPos(grid [][]rune, pos Vec2) bool {
  return pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[pos.y])
}

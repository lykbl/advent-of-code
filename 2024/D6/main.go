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

  loopCells := make(map[string]struct{})
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
    // log.Printf("Next: %c (%d_%d)", nextTile, nextPos.y, nextPos.x)
    if nextTile == '.' {
      if _, alreadyVisited := visitedCells[nextTileKey]; !alreadyVisited {
        visitedCells[nextTileKey] = struct{}{}
      }
      currentPos = Vec2 { x: nextPos.x, y: nextPos.y }
    }
    if nextTile == '#' {
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

  for y := range grid {
    for x := range grid {
      cellKey := fmt.Sprintf("%d_%d", y, x)
      if _, ok := visitedCells[cellKey]; ok && HasLoop(grid, guard.pos, guard.direction, Vec2 { x, y }) {
        loopCells[cellKey] = struct{}{}
      }
    }
  }

  p1 := len(visitedCells)
  log.Printf("P1: %d", p1)

  p2 := len(loopCells)
  log.Printf("P2: %d", p2)
}

func HasLoop(grid [][]rune, pos Vec2, dir rune, wallPos Vec2) bool {
  visitedCells := make(map[string]struct{})
  currentPos := pos
  currentDir := dir

  grid[wallPos.y][wallPos.x] = '#'
  hasLoop := false
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
    if nextTile == '.' {
      currentCellKey := fmt.Sprintf("%d_%d_%c", currentPos.y, currentPos.x, currentDir)
      if _, isLoop := visitedCells[currentCellKey]; isLoop {
        hasLoop = true
        break
      }
      visitedCells[currentCellKey] = struct{}{}
      currentPos = Vec2 { x: nextPos.x, y: nextPos.y }
    }
    if nextTile == '#' {
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
  grid[wallPos.y][wallPos.x] = '.'

  return hasLoop
}

func IsValidPos(grid [][]rune, pos Vec2) bool {
  return pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[pos.y])
}

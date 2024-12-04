package main

import (
	"bufio"
	"log"
	"os"
)

type Vec2 struct {
  x int
  y int
}

func IsXMas(grid [][]rune, pos Vec2) bool {
  if pos.y - 1 < 0 || pos.x - 1 < 0 || pos.x + 1 >= len(grid[0]) || pos.y + 1 >= len(grid) {
    return false
  }
  if grid[pos.y][pos.x] != 'A' {
    return false
  }
  //"M*M",
  //"*A*",
  //"S*S",
  //
  //"S*S",
  //"*A*",
  //"M*M",
  //
  //"M*S",
  //"*A*",
  //"M*S",
  //
  //"S*M",
  //"*A*",
  //"S*M",

  if grid[pos.y - 1][pos.x - 1] == grid[pos.y - 1][pos.x + 1] &&
     grid[pos.y + 1][pos.x - 1] == grid[pos.y + 1][pos.x + 1] &&
    (grid[pos.y - 1][pos.x - 1] == 'M' && grid[pos.y + 1][pos.x - 1] == 'S' ||
     grid[pos.y + 1][pos.x - 1] == 'M' && grid[pos.y - 1][pos.x - 1] == 'S') {
    return true
  }

  if grid[pos.y - 1][pos.x - 1] == grid[pos.y + 1][pos.x - 1] &&
     grid[pos.y - 1][pos.x + 1] == grid[pos.y + 1][pos.x + 1] &&
     (grid[pos.y - 1][pos.x - 1] == 'M' && grid[pos.y - 1][pos.x + 1] == 'S' ||
     grid[pos.y - 1][pos.x - 1] == 'S' && grid[pos.y - 1][pos.x + 1] == 'M') {
    return true
  }

  return false
}

func CountXmas(grid [][]rune, pos Vec2) int {
  target := "XMAS"
  directions := []Vec2{
    { x: 1, y: 0 },
    { x: -1, y: 0 },
    { x: 0, y: 1 },
    { x: 0, y: -1 },
    { x: 1, y: 1 },
    { x: -1, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 1 },
  }

  found := 0
  failed := make(map[int]struct{}, len(directions))
  for i, targetChar := range target {
    for dirI, direction := range directions {
      if targetChar == 'X' {
        continue
      }
      if _, alreadyFailed := failed[dirI]; alreadyFailed {
        continue
      }

      curPos := Vec2 { x: pos.x + direction.x * i, y: pos.y + direction.y * i }
      if curPos.x < 0 || curPos.y < 0 || curPos.x >= len(grid[0]) || curPos.y >= len(grid) {
        failed[dirI] = struct{}{}
        continue
      }

      curChar := grid[curPos.y][curPos.x]
      if targetChar != curChar {
        failed[dirI] = struct{}{}
        continue
      }

      if targetChar == 'S' {
        found++
      }
    }
  }

  return found
}

func main() {
  // file, _ := os.Open("test.txt")
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)

  var grid [][]rune
  for scanner.Scan() {
    grid = append(grid, []rune(scanner.Text()))
  }

  p1 := 0
  p2 := 0
  for y := range grid {
    for x := range grid[y] {
      if grid[y][x] == 'X' {
        p1 += CountXmas(grid, Vec2 { x, y })
      }
      if IsXMas(grid, Vec2 { x, y }) {
        p2 += 1
      }
    }
  }
  log.Printf("P1: %d,", p1)
  log.Printf("P2: %d", p2)
}

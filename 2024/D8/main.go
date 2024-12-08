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

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("File missing")
  }

  scanner := bufio.NewScanner(file)
  antennaLocations := map[rune][]Vec2{}
  grid := [][]rune{}

  y := 0
  for scanner.Scan() {
    line := scanner.Text()
    grid = append(grid, []rune(line))
    for x, char := range line {
      if char != '.' {
        antennaLocations[char] = append(antennaLocations[char], Vec2 { x, y })
      }
    }
    y++
  }

  antinodeLocations := map[string]struct{}{}
  for _, locations := range antennaLocations {
    for a, aL := range locations {
      for _, bL := range locations[a+1:] {
        for _, antinode := range GetAntinodes(grid, aL, bL) {
          antinodeLocations[fmt.Sprintf("%d_%d", antinode.y, antinode.x)] = struct{}{}
        }
      }
    }
  }

  p2 := len(antinodeLocations)
  log.Printf("P1: %d", p2)
}

func GetAntinodes(grid [][]rune, a Vec2, b Vec2) []Vec2 {
  result := make([]Vec2, 0)
  dx, dy := b.x - a.x, b.y - a.y
  for i := 0; IsValidPos(grid, Vec2 { a.x - dx * i, a.y - dy * i }); i++ {
    result = append(result, Vec2 { a.x - dx * i, a.y - dy * i })
  }
  for i := 0; IsValidPos(grid, Vec2 { b.x + dx * i, b.y + dy * i }); i++ {
    result = append(result, Vec2 { b.x + dx * i, b.y + dy * i })
  }

  return result
}

func IsValidPos(grid [][]rune, pos Vec2) bool {
  return pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[pos.y])
}

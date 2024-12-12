package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var cache = make(map[string]int)
func main() {
  // f, err := os.Open("test.txt")
  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("F")
  }

  scanner := bufio.NewScanner(f)

  grid := make([][]rune, 0)
  for scanner.Scan() {
    line := scanner.Text()
    grid = append(grid, []rune(line))
  }

  farmPlots := make(map[rune][][][2]int, 0)
  visited := make(map[string]struct{}, 0)
  for y := range grid {
    for x := range grid {
      farmPlot := make([][2]int, 0)
      farmType := grid[y][x]
      if _, exists := visited[fmt.Sprintf("%d_%d", y, x)]; exists {
        continue
      }

      toVisit := make([][2]int, 0)
      toVisit = append(toVisit, [2]int{y,x})
      for len(toVisit) > 0 {
        currentPos := toVisit[len(toVisit) - 1]
        toVisit = toVisit[:len(toVisit) - 1]
        if _, exists := visited[fmt.Sprintf("%d_%d", currentPos[0], currentPos[1])]; exists {
          continue
        }

        visited[fmt.Sprintf("%d_%d", currentPos[0], currentPos[1])] = struct{}{}
        farmPlot = append(farmPlot, currentPos)
        dirs := [][2]int{
          { currentPos[0] + 1, currentPos[1] },
          { currentPos[0] - 1, currentPos[1] },
          { currentPos[0], currentPos[1] - 1},
          { currentPos[0], currentPos[1] + 1 },
        }

        for _, newPos := range dirs {
          if IsValid(grid, newPos) && grid[newPos[0]][newPos[1]] == farmType {
            toVisit = append(toVisit, newPos)
          }
        }
      }
      farmPlots[farmType] = append(farmPlots[farmType], farmPlot)
    }
  }

  p1 := 0
  for farmType, farmPlots := range farmPlots {
    for _, farmPlot := range farmPlots {
      totalArea := PlotArea(farmPlot)
      totalPerim := PlotPerim(grid, farmPlot)
      log.Printf("A region of %c plants with price %d * %d = %d.\n", farmType, totalArea, totalPerim, totalArea * totalPerim)
      p1 += totalArea * totalPerim
    }
  }

  log.Printf("P1: %d", p1)
}

func PlotArea(plot [][2]int) int {
  return len(plot)
}

func PlotPerim(grid [][]rune, plot [][2]int) int {
  farmType := grid[plot[0][0]][plot[0][1]]
  perim := 0
  for _, currentPos := range plot {
    dirs := [][2]int{
      { currentPos[0] + 1, currentPos[1] },
      { currentPos[0] - 1, currentPos[1] },
      { currentPos[0], currentPos[1] - 1},
      { currentPos[0], currentPos[1] + 1 },
    }

    neighbours := 0
    for _, dir := range dirs {
      if IsValid(grid, dir) && grid[dir[0]][dir[1]] == farmType {
        neighbours++
      }
    }

    perim += 4 - neighbours
  }

  return perim
}

func IsValid(grid [][]rune, pos [2]int) bool {
  return pos[0] >= 0 && pos[1] >= 0 && pos[0] < len(grid) && pos[1] < len(grid[pos[0]])
}
 

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  // f, _ := os.Open("test.txt")
  f, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(f)

  grid := [][]rune{}
  y := 0

  var start [2]int
  var end [2]int
  for scanner.Scan() {
    line := scanner.Text()

    grid = append(grid, []rune{})
    for x, char := range line {
      grid[y] = append(grid[y], char)
      if char == 'S' {
        start = [2]int{x, y}
        grid[y][x] = '.'
      }
      if char == 'E' {
        end = [2]int{x, y}
        grid[y][x] = '.'
      }
    }
    y++
  }

  trackTimes := map[string]int{}
  queue := make([][2]int, 0)
  queue = append(queue, start)
  dirs := [][2]int{
    { 1, 0 },
    { -1, 0 },
    { 0, 1 },
    { 0, -1 },
  }

  trackTime := 0
  currentCell := start
  for {
    trackTimes[fmt.Sprintf("%d_%d", currentCell[0], currentCell[1])] = trackTime
    trackTime++
    if currentCell == end {
      break
    }

    for _, dir := range dirs {
      nextPos := [2]int{currentCell[0] + dir[0], currentCell[1] + dir[1]}
      if grid[nextPos[1]][nextPos[0]] != '.' {
        continue
      }
      if _, visited := trackTimes[fmt.Sprintf("%d_%d", nextPos[0], nextPos[1])]; visited {
        continue
      }

      currentCell = nextPos
      break
    }
  }

  // fmt.Printf("Track times: %v\n", trackTimes)
  // cheatTime := 2
  cheatTimes := make(map[int]int, 0)
  for cellKey, cellTime := range trackTimes {
    parts := strings.Split(cellKey, "_")
    x, _ := strconv.Atoi(parts[0])
    y, _ := strconv.Atoi(parts[1])

    for _, dir := range dirs {
      if y + dir[1] * 2 >= len(grid) || y + dir[1] * 2 < 0 || x + dir[0] * 2 < 0 || x + dir[0] * 2 >= len(grid[y]) {
        continue
      }

      if grid[y + dir[1]][x + dir[0]] == '#' && grid[y + dir[1] * 2][x + dir[0] * 2] == '.' {
        skipCellKey := fmt.Sprintf("%d_%d", x + dir[0] * 2, y + dir[1] * 2)
        skipCellTime := trackTimes[skipCellKey]
        timeSaved := skipCellTime - cellTime - 2
        if timeSaved > 0 {
          cheatTimes[timeSaved] += 1
        }
      }
    }
  }

  p1 := 0
  for timeSaved, cheatsCount := range cheatTimes {
    if timeSaved >= 100 {
      p1 += cheatsCount
    }
  }
  fmt.Printf("P1: %d\n", p1)

  maxCheatSteps := 20
  cheatSavedTimes := make(map[string]int, 0)
  for cellKey, cellTime := range trackTimes {
    parts := strings.Split(cellKey, "_")
    x, _ := strconv.Atoi(parts[0])
    y, _ := strconv.Atoi(parts[1])

    visited := make(map[string]struct{}, 0)
    queue := make([]CheatStep, 0)
    queue = append(queue, CheatStep{
      curPos: [2]int{x, y},
      stepsTaken: 0,
    })

    maxCheatSteps = 20
    for len(queue) > 0 {
      curStep := queue[0]
      queue = queue[1:]
      if curStep.stepsTaken > maxCheatSteps {
        continue
      }
      if _, alreadyVisited := visited[fmt.Sprintf("%d_%d", curStep.curPos[0], curStep.curPos[1])]; alreadyVisited {
        continue
      }

      visited[fmt.Sprintf("%d_%d", curStep.curPos[0], curStep.curPos[1])] = struct{}{}
      if curStep.stepsTaken <= maxCheatSteps && grid[curStep.curPos[1]][curStep.curPos[0]] == '.' {
        skipCellTime := trackTimes[fmt.Sprintf("%d_%d", curStep.curPos[0], curStep.curPos[1])]
        timeSaved := skipCellTime - cellTime - curStep.stepsTaken

        maxTimeSaved := cheatSavedTimes[fmt.Sprintf("%d_%d-%d_%d", x, y, curStep.curPos[0], curStep.curPos[1])]
        if maxTimeSaved < timeSaved {
          cheatSavedTimes[fmt.Sprintf("%d_%d-%d_%d", x, y, curStep.curPos[0], curStep.curPos[1])] = timeSaved
        }
      }

      for _, dir := range dirs {
        nextPos := [2]int{curStep.curPos[0] + dir[0], curStep.curPos[1] + dir[1]}
        if nextPos[1] >= len(grid) || nextPos[1] < 0 || nextPos[0] < 0 || nextPos[0] >= len(grid[y]) {
          continue
        }

        queue = append(queue, CheatStep {
          curPos: nextPos,
          stepsTaken: curStep.stepsTaken + 1,
        })
      }
    }
  }

  counts := make(map[int]int, 0)
  for _, timeSaved := range cheatSavedTimes {
    counts[timeSaved]++
  }
  
  p2 := 0
  for timeSaved, count := range counts {
    if timeSaved >= 100 {
      p2 += count
    }
  }

  fmt.Printf("P2: %d\n", p2)
}

type CheatStep struct {
  curPos [2]int
  stepsTaken int
  path [][2]int
}

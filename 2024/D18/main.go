package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
  height, width, simLimit := 70, 70, 1024
  f, _ := os.Open("input.txt")
  // height, width, simLimit := 6, 6, 12
  // f, _ := os.Open("test.txt")

  scanner := bufio.NewScanner(f)
  grid := make([][]rune, 0)
  for y := 0; y <= height; y++ {
    grid = append(grid, make([]rune, 0))
    for x := 0; x <= width; x++ {
      grid[y] = append(grid[y], '.')
    }
  }

  stepI := 0
  corruptedBytes := [][2]int{}
  for scanner.Scan() {
    line := scanner.Text()
    r := regexp.MustCompile("([0-9]*),([0-9]*)")
    matches := r.FindStringSubmatch(line)
    xS, yS := matches[1], matches[2]
    x, _ := strconv.Atoi(xS)
    y, _ := strconv.Atoi(yS)
    corruptedBytes = append(corruptedBytes, [2]int{x, y})
    if stepI <= simLimit {
      grid[y][x] = '#'
    }

    stepI++
  }

  start := [2]int{0, 0}
  end := [2]int{width, height}
  queue := []Step{
    {curPos: start, steps: 0},
  }
  p1 := 0
  visited := make(map[string]struct{}, 0)
  visited[fmt.Sprintf("%d_%d", start[0], start[1])] = struct{}{}
  for len(queue) > 0 {
    curStep := queue[0]
    queue = queue[1:]
    if curStep.curPos == end {
      p1 = curStep.steps
      break
    }

    dirs := [][2]int{
      {0, 1},
      {1, 0},
      {-1, 0},
      {0, -1},
    }
    for _, dir := range dirs {
      newPos := [2]int{curStep.curPos[0] + dir[0], curStep.curPos[1] + dir[1]}
      if newPos[0] < 0 || newPos[1] < 0 || newPos[0] > width || newPos[1] > height {
        continue
      }

      if grid[newPos[1]][newPos[0]] == '#' {
        continue
      }
      if _, alreadyVisited := visited[fmt.Sprintf("%d_%d", newPos[0], newPos[1])]; alreadyVisited {
        continue
      }
      visited[fmt.Sprintf("%d_%d", newPos[0], newPos[1])] = struct{}{}

      queue = append(queue, Step {
        curPos: newPos,
        steps: curStep.steps + 1,
      })
    }
  }
  fmt.Printf("P1: %d", p1)

  for curSimLimit := simLimit; curSimLimit < len(corruptedBytes); curSimLimit++ {
    answerExists := false
    queue = []Step{
      {curPos: start, steps: 0},
    }
    grid[corruptedBytes[curSimLimit][1]][corruptedBytes[curSimLimit][0]] = '#'
    visited = make(map[string]struct{}, 0)
    visited[fmt.Sprintf("%d_%d", start[0], start[1])] = struct{}{}
    for len(queue) > 0 {
      curStep := queue[0]
      queue = queue[1:]
      if curStep.curPos == end {
        answerExists = true
        break
      }

      dirs := [][2]int{
        {0, 1},
        {1, 0},
        {-1, 0},
        {0, -1},
      }
      for _, dir := range dirs {
        newPos := [2]int{curStep.curPos[0] + dir[0], curStep.curPos[1] + dir[1]}
        if newPos[0] < 0 || newPos[1] < 0 || newPos[0] > width || newPos[1] > height || grid[newPos[1]][newPos[0]] == '#' {
          continue
        }
        if _, alreadyVisited := visited[fmt.Sprintf("%d_%d", newPos[0], newPos[1])]; alreadyVisited {
          continue
        }
        visited[fmt.Sprintf("%d_%d", newPos[0], newPos[1])] = struct{}{}

        queue = append(queue, Step {
          curPos: newPos,
          steps: curStep.steps + 1,
        })
      }
    }

    if answerExists == false {
      fmt.Printf("No longer found at: %d, %v\n", curSimLimit, corruptedBytes[curSimLimit])
      break
    }
  }
}

type Step struct {
  curPos [2]int
  steps int
}

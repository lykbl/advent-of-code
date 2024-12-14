package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
  // file, err := os.Open("test.txt")
  // roomSize := [2]int{7, 11}
  // steps := 25

  file, err := os.Open("input.txt")
  roomSize := [2]int{103, 101}
  steps := 100

  if err != nil {
    log.Fatal("FAIL")
  }

  scanner := bufio.NewScanner(file)
  robots := make([]struct{ x int; y int; vx int; vy int }, 0)
  for scanner.Scan() {
    line := scanner.Text()

    pattern := `p=([\-\d]*),([\-\d]*) v=([\-\d]*),([\-\d]*)`
    re := regexp.MustCompile(pattern)
    matches := re.FindStringSubmatch(line)

    x, err := strconv.Atoi(matches[1])
    y, err := strconv.Atoi(matches[2])
    vx, err := strconv.Atoi(matches[3])
    vy, err := strconv.Atoi(matches[4])
    if err != nil {
      log.Fatal("ATOI")
    }
    robots = append(robots, struct{x int; y int; vx int; vy int}{
      x, y, vx, vy,
    })
  }

  finalPositions := make([][2]int, 0)
  for _, robot := range robots {
    finalPos := [2]int{ robot.y, robot.x }
    for i := 0; i < steps; i++ {
      newY := robot.vy + finalPos[0]
      if newY >= roomSize[0] || newY < 0 {
        newY = mod(newY, roomSize[0])
      }

      newX := robot.vx + finalPos[1]
      if newX >= roomSize[1] || newX < 0 {
        newX = mod(newX, roomSize[1])
      }
      finalPos = [2]int{newY, newX}
    }
    
    finalPositions = append(finalPositions, finalPos)
  }

  grid := make([][]rune, 0)
  for y := 0; y < roomSize[0]; y++ {
    grid = append(grid, []rune{})
    for x := 0; x < roomSize[1]; x++ {
      grid[y] = append(grid[y], '.')
    }
  }

  steps = 10043
  stepsMemory := make(map[string]int)
  outputF, _ := os.Create("output.txt")
  writer := bufio.NewWriter(outputF)
  for step := 1; step < steps; step++ {
    positionsMap := make(map[string]struct{}, 0)
    for i, robot := range robots {
      newY := robot.vy + robot.y
      if newY >= roomSize[0] || newY < 0 {
        newY = mod(newY, roomSize[0])
      }

      newX := robot.vx + robot.x
      if newX >= roomSize[1] || newX < 0 {
        newX = mod(newX, roomSize[1])
      }

      robots[i].y = newY
      robots[i].x = newX
      positionsMap[fmt.Sprintf("%d_%d", robots[i].y, robots[i].x)] = struct{}{}
    }

    toPrint := fmt.Sprintf("\n\n\nStep: %d\n", step)
    memory := ""
    for y := 0; y < roomSize[0]; y++ {
      line := ""
      for x := 0; x < roomSize[1]; x++ {
        if _, exists := positionsMap[fmt.Sprintf("%d_%d", y, x)]; exists {
          memory += "*"
          line += "*"
        } else {
          memory += "."
          line += "."
        }
      }
      memory += "\n"
    }

    writer.WriteString(toPrint + memory)
    if _, found := stepsMemory[memory]; found {
      break
    }
  }

  quadrants := make(map[int]int, 4)
  for _, fp := range finalPositions  {
    if fp[0] < roomSize[0] / 2 && fp[1] < roomSize[1] / 2 {
      quadrants[0]++
    }
    if fp[0] < roomSize[0] / 2 && fp[1] > roomSize[1] / 2 {
      quadrants[1]++
    }
    if fp[0] > roomSize[0] / 2 && fp[1] < roomSize[1] / 2 {
      quadrants[2]++
    }
    if fp[0] > roomSize[0] / 2 && fp[1] > roomSize[1] / 2 {
      quadrants[3]++
    }
  }
  p1 := 1
  for _, quadrant := range quadrants {
    p1 *= quadrant
  }

  log.Printf("P1: %d", p1)
}

func mod(a, b int) int {
    return (a % b + b) % b
}


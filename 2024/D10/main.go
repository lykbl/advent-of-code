package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetMoves(g [][]rune, pos [2]int) [][2]int {
  moves := [][2]int{
   { pos[0] + 1, pos[1] },
   { pos[0] - 1, pos[1] },
   { pos[0], pos[1] + 1 },
   { pos[0], pos[1] - 1},
  }

  validMoves := make([][2]int, 0)
  for _, move := range moves {
    log.Printf("Trying move: %v", move)
    if IsValidPos(g, move) {
      validMoves = append(validMoves, move)
      log.Printf("added")
    }
  }

  return validMoves
}

func IsValidPos(g [][]rune, pos [2]int) bool {
  return pos[1] >= 0 && pos[0] >= 0 && pos[0] < len(g) && pos[1] < len(g[pos[0]])
} 

type Step struct {
  visitedCells map[string]struct{}
  currentPos [2]int
}

func main() {
  // f, err := os.Open("test.txt")
  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("F")
  }

  scanner := bufio.NewScanner(f)

  grid := make([][]rune, 0)
  trailheads := make([][2]int, 0)
  y := 0 
  for scanner.Scan() {
    line := scanner.Text()
    grid = append(grid, []rune(line))
    for x, char := range line {
      if char == '0' {
        trailheads = append(trailheads, [2]int{ y, x })
      }
    }

    y++
  }

  scores := make(map[string]int, 0)
  for _, startingPos := range trailheads {
    visitedCells := make(map[string]struct{}, 0)
    visitQueue := make([]Step, 0)
    visitQueue = append(visitQueue, Step {
      currentPos: startingPos,
      visitedCells: visitedCells,
    })

    score := 0
    scoreKey := fmt.Sprintf("%d_%d", startingPos[0], startingPos[1])
    log.Printf("Doing: %v", startingPos)
    for len(visitQueue) > 0 {
      currentStep := visitQueue[len(visitQueue) - 1]
      visitQueue = visitQueue[0:len(visitQueue) - 1]
      log.Printf("New queue: %v, %v", currentStep, visitQueue)
      if _, alreadyVisited := currentStep.visitedCells[fmt.Sprintf("%d_%d", currentStep.currentPos[0], currentStep.currentPos[1])]; alreadyVisited {
        // continue
      }
      currentStep.visitedCells[fmt.Sprintf("%d_%d", currentStep.currentPos[0], currentStep.currentPos[1])] = struct{}{}
      currentHeight := grid[currentStep.currentPos[0]][currentStep.currentPos[1]]
      if currentHeight == '9' {
        score++
        continue
      }

      moves := GetMoves(grid, currentStep.currentPos)
      log.Printf("Your movews: %v", moves)
      for _, move := range moves {
        moveHeight := grid[move[0]][move[1]]
        if moveHeight - currentHeight == 1 {
          visitQueue = append(visitQueue, Step {
            visitedCells: visitedCells,
            currentPos: move,
          })
        }
      }
    }
    scores[scoreKey] = score
  }

  p1 := 0
  for _, score := range scores {
    p1 += score
  }

  log.Printf("P1: %d", p1)
}


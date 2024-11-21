package main

import (
	"bufio"
	"fmt"
	// "fmt"
	"log"
	"os"
)


const (
  StartingRune rune = 'S'
  GardenRune rune = '.'
  RockRune rune = '#'
  VisitedRune rune = 'O'
)

type Garden struct {
  grid [][]rune
  elfPosition [2]int
}

func fileToGarden(filePath string) (*Garden, error) {
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  garden := &Garden{ grid: make([][]rune, 0)  }

  y := 0
  for scanner.Scan() {
    line := []rune(scanner.Text())

    for x, char := range line {
      if char == StartingRune {
        garden.elfPosition = [2]int{x, y}
      }
    }

    garden.grid = append(garden.grid, line)
    y += 1
  }

  if err := scanner.Err(); err != nil {
		return nil, err
	}

  return garden, nil
}

type Step struct {
  position [2]int
  stepsLeft int
}

// func (g *Garden) Ge

func (g *Garden) CountPossiblePlots(stepsToTake int) int {
  possiblePlots := 1
  moveDirections := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
  // visitedPlots := make(map[string]bool)
  cellsToCheck := []Step{
    {position: g.elfPosition, stepsLeft: stepsToTake},
  }

  log.Print(fmt.Printf("%d", 5))
  // visitedPositions := make(map[string]bool)
  g.grid[g.elfPosition[1]][g.elfPosition[0]] = VisitedRune
  // cellKey, err := fmt.Printf("%d_%d", g.elfPosition[0], g.elfPosition[1])
  // visitedPositions[cellKey] = true
  for len(cellsToCheck) > 0 {
    currentStep := cellsToCheck[0]
    cellsToCheck = cellsToCheck[1:]
    if currentStep.stepsLeft < 0 {
      continue
    }

    for _, moveOffset := range moveDirections {
      potentialPlotPosition := [2]int{moveOffset[0] + currentStep.position[0], moveOffset[1] + currentStep.position[1]}

      isValidPlot := potentialPlotPosition[0] >= 0 && potentialPlotPosition[1] >= 0 && potentialPlotPosition[0] < len(g.grid[0]) && potentialPlotPosition[1] < len(g.grid)
      if !isValidPlot {
        continue
      }

      potentialRune := g.grid[potentialPlotPosition[1]][potentialPlotPosition[0]]
      if potentialRune == RockRune {
        continue
      }

      currentRune := g.grid[currentStep.position[1]][currentStep.position[0]]
      if currentRune == VisitedRune && potentialRune == GardenRune {
        cellsToCheck = append(cellsToCheck, Step { position: potentialPlotPosition, stepsLeft: currentStep.stepsLeft - 1 })
      }
      if currentRune == GardenRune && potentialRune == GardenRune {
        g.grid[potentialPlotPosition[1]][potentialPlotPosition[0]] = VisitedRune
        possiblePlots += 1
        cellsToCheck = append(cellsToCheck, Step { position: potentialPlotPosition, stepsLeft: currentStep.stepsLeft - 1 })
      }
    }
  }

  return possiblePlots
}

func main() {
  // fileName := "test.txt"
  fileName := "input.txt"
  
  var stepsToTake int
  if fileName == "test.txt" {
    stepsToTake = 100
  } else {
    stepsToTake = 64
  }
  
  garden, err  := fileToGarden(fileName)

  if err != nil {
    log.Fatalf("File not opened: %s", err)
  }

  movesCount := garden.CountPossiblePlots(stepsToTake)

  for _, line := range garden.grid {
    // for _, char := range line {
    //   log.Printf("%s", char)
    // }
    line := string(line)
    log.Printf("%s\n", line)
  }

  log.Printf("Result: %d", movesCount)
}

// 128 low

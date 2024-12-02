package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal("File rr")
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  result := 0
  minChange := 1
  maxChange := 3
  for scanner.Scan() {
    line := scanner.Text()
    levelsStr := strings.Split(line, " ")
    
    levels := make([]int, len(levelsStr))
    for i, levelStr := range levelsStr {
      level, err := strconv.Atoi(levelStr)
      if err != nil {
        log.Fatal("conv err")
      }

      levels[i] = level
    }

    for y := 0; y < len(levels); y++ {
      lastDiff := 0
      isValid := true

      newLevels := make([]int, 0, len(levels)-1)
      newLevels = append(newLevels, levels[:y]...)
      newLevels = append(newLevels, levels[y+1:]...)

      for i := 0; i < len(newLevels) - 1 && isValid; i++ {
        currDif := newLevels[i] - newLevels[i + 1]

        if math.Abs(float64(currDif)) < float64(minChange) || math.Abs(float64(currDif)) > float64(maxChange) || (currDif > 0 && lastDiff < 0 || currDif < 0 && lastDiff > 0) {
            isValid = false
        }

        lastDiff = currDif
      }

      if isValid {
        result++
        break
      }
    }
  }

  log.Printf("Result: %d", result)
}

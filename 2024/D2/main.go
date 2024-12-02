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

  p1 := 0
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

    lastDiff := levels[0] - levels[1]
    isValid := levels[0] != levels[1]
    for i := 0; i < len(levels) - 1 && isValid; i++ {
      currDif := levels[i] - levels[i + 1]

      if math.Abs(float64(currDif)) < float64(minChange) || math.Abs(float64(currDif)) > float64(maxChange) || currDif ^ lastDiff < 0 {
        isValid = false
      }

      lastDiff = currDif
    }

    if isValid {
      p1++
    }
  }

  log.Printf("Result: %d", p1)
}

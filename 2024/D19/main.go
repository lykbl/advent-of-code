package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  f, _ := os.Open("input.txt")
  // f, _ := os.Open("test.txt")
  scanner := bufio.NewScanner(f)
  i := 0
  patterns := make([]string, 0)
  designs := make([]string, 0)
  for scanner.Scan() {
    line := scanner.Text()
    if i == 0 {
      patterns = strings.Split(line, ",")
      for i := range patterns {
        patterns[i] = strings.TrimSpace(patterns[i])
      }
    } else if i == 1 {
    } else {
      designs = append(designs, strings.TrimSpace(line))
    }
    i++
  }

  fmt.Printf("Patterns: %v\nDeisngs: %v\n", patterns, designs)
  valid := 0
  memory := make(map[string]int, 0)
  p2 := 0
  for i, design := range designs {
    fmt.Printf("Doing %d out of %d\n", i, len(designs))
    counts := IsDesignPossible(patterns, design, memory)
    if counts > 0 {
      valid++
      p2 += counts
    }
  }

  fmt.Printf("Valid: %d\n", valid)
  fmt.Printf("P2: %d\n", p2)
}

func IsDesignPossible(patterns []string, design string, memory map[string]int) int {
  if cached, alreadyFound := memory[design]; alreadyFound {
    return cached
  }

  if design == "" {
    return 1
  }

  count := 0
  for _, pattern := range patterns {
    newDesign, found := strings.CutPrefix(design, pattern)
    if found {
      count += IsDesignPossible(patterns, newDesign, memory)
    }
  }

  memory[design] = count
  return count
}

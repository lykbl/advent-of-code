package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("file")
  }

  scanner := bufio.NewScanner(file)

  i := 0
  maxCal := i
  elfCalories := make(map[int]int, 0)
  for scanner.Scan() {
    line := scanner.Text()

    if strings.TrimSpace(line) == "" {
      if elfCalories[i] > elfCalories[maxCal] {
        maxCal = i
      }
      i++
      continue
    } else {
      lineCalories, err := strconv.Atoi(line)
      if err != nil {
        log.Fatal("ATOI")
      }

      elfCalories[i] += lineCalories
    }
  }

  log.Printf("P1: %d", elfCalories[maxCal])
}

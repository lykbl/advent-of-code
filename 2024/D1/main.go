package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  
  if err != nil {
    log.Fatal("File failed")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  listLeft := make([]int, 0)
  listRight := make([]int, 0)

  for scanner.Scan() {
    line := scanner.Text()

    result := strings.Split(line, "   ")
    if len(result) > 2 {
      log.Fatal("Faield to split")
    }

    left, _ := strconv.Atoi(result[0])
    right, _ := strconv.Atoi(result[1])

    listLeft = append(listLeft, left)
    listRight = append(listRight, right)
  }

  if scanner.Err() != nil {
    log.Fatalf("Scan failed, %v", scanner.Err())
  }

  slices.Sort(listLeft)
  slices.Sort(listRight)

  p1 := 0
  for i := range listLeft {
    p1 += int(math.Abs(float64(listRight[i] - listLeft[i])))
  }
  log.Printf("P1: %d", p1)
}

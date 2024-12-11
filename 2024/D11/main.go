package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cache = make(map[string]int)
func main() {
  // f, err := os.Open("test.txt")
  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("F")
  }

  scanner := bufio.NewScanner(f)

  stones := make([]int, 0)
  for scanner.Scan() {
    line := scanner.Text()

    stonesString := strings.Split(line, " ")
    for _, stoneString := range stonesString {
      stone, err := strconv.Atoi(stoneString)
      if err != nil {
        log.Fatal("ATOI")
      }

      stones = append(stones, stone)
    }
  }

  log.Printf("P1: %d", Blink(slices.Clone(stones), 25))
  log.Printf("P2: %d", Blink(slices.Clone(stones), 75))
}

func FlipStones(stones []int) []int {
  for i := 0; i < len(stones); i++ {
    if stones[i] == 0 {
      stones[i] = 1
    } else if stoneStr := strconv.Itoa(stones[i]); len(strconv.Itoa(stones[i])) % 2 == 0 {
      leftHalf,  _ := strconv.Atoi(stoneStr[:len(stoneStr) / 2])
      rightHalf, _ := strconv.Atoi(stoneStr[len(stoneStr) / 2:])

      stones[i] = leftHalf
      i++
      stones = slices.Insert(stones, i, rightHalf)
    } else {
      stones[i] = stones[i] * 2024
    }
  }

  return stones
}

func Blink(stones []int, blinksRemaining int) int {
  if blinksRemaining == 0 {
    return len(stones)
  }

  stones = FlipStones(stones)
  result := 0
  for _, stone := range stones {
    cacheKey := fmt.Sprintf("%d_%d", blinksRemaining - 1, stone)
    if _, alreadyCached := cache[cacheKey]; !alreadyCached {
      cache[cacheKey] = Blink([]int{stone}, blinksRemaining - 1)
    }

    result += cache[cacheKey]
  }

  return result
}


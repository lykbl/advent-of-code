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

  blinkTimes := 25
  stonesP1 := slices.Clone(stones)
  for y := 0; y < blinkTimes; y++ {
    stonesP1 = FlipStones(stonesP1)
  }

  p1 := len(stonesP1)
  log.Printf("P1: %d", p1)

  blinkTimes = 75
  cache := make(map[string]int, 0)
  p2 := Blink(stones, blinkTimes, cache)
  log.Printf("P2: %d", p2)
}

func FlipStones(stones []int) []int {
  for i := 0; i < len(stones); i++ {
    stone := stones[i]
    stoneStr := strconv.Itoa(stone) 
    if stone == 0 {
      stones[i] = 1
    } else if len(stoneStr) % 2 == 0 {
      stoneA := stoneStr[:len(stoneStr)/ 2]
      stoneAN, _ := strconv.Atoi(stoneA)

      stoneB := stoneStr[len(stoneStr) / 2:]
      stoneBN, _ := strconv.Atoi(stoneB)

      stones[i] = stoneAN
      i++
      stones = slices.Insert(stones, i, stoneBN)
    } else {
      stones[i] = stones[i] * 2024
    }
  }

  return stones
}

func Blink(stones []int, blinksRemaining int, cache map[string]int) int {
  if blinksRemaining == 0 {
    return len(stones)
  }

  stones = FlipStones(stones)
  result := 0
  for i := 0; i < len(stones); i++ {
    if cachedResult, alreadyCached := cache[fmt.Sprintf("%d_%d", blinksRemaining - 1, stones[i])]; alreadyCached {
      result += cachedResult
    } else {
      calculated := Blink([]int{stones[i]}, blinksRemaining - 1, cache)
      cache[fmt.Sprintf("%d_%d", blinksRemaining - 1, stones[i])] = calculated
      result += calculated
    }
  }

  return result
}


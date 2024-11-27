package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"

	// "slices"
	"strconv"
	"strings"
)

type Brick struct {
  minX int
  maxX int
  minY int
  maxY int
  minZ int
  maxZ int
  label *rune
}

func (a *Brick) IsOverlapping(b *Brick) bool {
  return max(a.minX, b.minX) <= min(a.maxX, b.maxX) && max(a.minY, b.minY) <= min(a.maxY, b.maxY)
}

func BrickFromCorners(minP [3]int, maxP[3] int, label *rune) *Brick {
  return &Brick {
    minX: minP[0],
    minY: minP[1],
    minZ: minP[2],
    maxX: maxP[0],
    maxY: maxP[1],
    maxZ: maxP[2],
    label: label,
  }
}

func (b *Brick) String() string {
  return fmt.Sprintf("%d,%d,%d~%d,%d,%d (%c)\n", b.minX, b.minY, b.minZ, b.maxX, b.maxY, b.maxZ, *b.label)
}


func main() {
  // filename := "debug.txt"
  filename := "input.txt"
  f, err := os.Open(filename)
  if err != nil {
    log.Fatalf("File: %s", err)
  }

  scanner := bufio.NewScanner(f)

  defer f.Close()

  bricks := make([]*Brick, 0)
  startingLabel := 'A'
  for scanner.Scan() {
    line := scanner.Text()

    var a, b [3]int
    parts := strings.Split(line, "~")
    for i, part := range parts {
      numStrs := strings.Split(part, ",")
      if len(numStrs) != 3 {
        panic("Each part must have exactly 3 numbers")
      }

      for j, numStr := range numStrs {
        num, err := strconv.Atoi(numStr)
        if err != nil {
          panic("Invalid number in input")
        }
        if i == 0 {
          a[j] = num
        } else {
          b[j] = num
        }
      }
    }

    label := startingLabel + rune(len(bricks))
    brickToAdd := BrickFromCorners(a, b, &label)
    bricks = append(bricks, brickToAdd)
  }

  log.Printf("Starting: %v", bricks)

  zComp := func(a, b *Brick) int {
    return cmp.Compare(a.maxZ, b.maxZ)
  }
  slices.SortFunc(bricks, zComp)

  for i, loweringBrick := range bricks {
    finalZ := 1
    for _, brickToCheck := range bricks[:i] {
      if loweringBrick.IsOverlapping(brickToCheck) {
        finalZ = max(finalZ, brickToCheck.maxZ + 1)
      }
    }

    loweringBrick.maxZ -= loweringBrick.minZ - finalZ
    loweringBrick.minZ = finalZ
  }

  slices.SortFunc(bricks, zComp)
  log.Printf("Bricks (S): %v", bricks)

  k_supports_v := make(map[int]map[int]bool, len(bricks))
  v_supports_k := make(map[int]map[int]bool, len(bricks))

  for i := range bricks {
    k_supports_v[i] = make(map[int]bool, len(bricks))
    v_supports_k[i] = make(map[int]bool, len(bricks))
  }

  for j, upperBrick := range bricks {
    for i, lowerBrick := range bricks[:j] {
      if upperBrick.IsOverlapping(lowerBrick) && upperBrick.minZ == lowerBrick.maxZ + 1 {
        k_supports_v[i][j] = true
        v_supports_k[j][i] = true
      }
    }
  }
  log.Printf("Supports: %v", k_supports_v)
  log.Printf("Supported: %v", v_supports_k)

  result := 0
  for i := range bricks {
    all := true
    for j := range k_supports_v[i] {
      if len(v_supports_k[j]) < 2 {
        all = false
      }
    }

    if all {
      result++
    }
  }

  log.Printf("Result: %d", result)
}

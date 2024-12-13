package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Hook struct {
  a struct {
    x int
    y int
  }
  b struct {
    x int
    y int
  }
  targetX int
  targetY int
}

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal("FAIL")
  }

  scanner := bufio.NewScanner(file)
  hooks := make([]Hook, 0)
  currentHook := Hook {}
  for scanner.Scan() {
    line := scanner.Text()

    if strings.Contains(line, "Button") {
      pattern := `Button ([A-Z]): X\+(\d+), Y\+(\d+)`
      re := regexp.MustCompile(pattern)
      matches := re.FindStringSubmatch(line)
      label := matches[1]
      x, _ := strconv.Atoi(matches[2])
      y, _ := strconv.Atoi(matches[3])

      if label == "A" {
        currentHook.a = struct{x int; y int}{x, y}
      }
      if label == "B" {
        currentHook.b = struct{x int; y int}{x, y}
      }
    }
    if strings.Contains(line, "Prize") {
      pattern := `Prize: X=(\d+), Y=(\d+)`
      re := regexp.MustCompile(pattern)
      matches := re.FindStringSubmatch(line)
      targetX, err := strconv.Atoi(matches[1])
      if err != nil {
        log.Printf("ATOI: %s", err.Error())
      }
      targetY, err := strconv.Atoi(matches[2])
      if err != nil {
        log.Printf("ATOI: %s", err.Error())
      }

      currentHook.targetX = targetX
      currentHook.targetY = targetY
    }
    if strings.TrimSpace(line) == "" {
      hooks = append(hooks, currentHook)
      currentHook = Hook{}
    }
  }

  hooks = append(hooks, currentHook)
  currentHook = Hook{}

  p1 := 0
  for _, hook := range hooks {
    a, b, _ := SolveForHook(float64(hook.targetX), float64(hook.a.x), float64(hook.b.x), float64(hook.targetY), float64(hook.a.y), float64(hook.b.y))
    p1 += int(a * 3 + b * 1)
  }
  log.Printf("P1: %d", p1)

  p2 := 0
  offset := 10000000000000
  for _, hook := range hooks {
    a, b, _ := SolveForHook(float64(hook.targetX + offset), float64(hook.a.x), float64(hook.b.x), float64(hook.targetY + offset), float64(hook.a.y), float64(hook.b.y))
    p2 += int(a * 3 + b * 1)
  }

  log.Printf("P2: %d", p2)
}

func SolveForHook(targetX, aX, bX, targetY, aY, bY float64) (float64, float64, error) {
	det := aX*bY - bX*aY

	if math.Abs(det) < 1e-10 {
		return 0, 0, fmt.Errorf("no unique solution exists")
	}

	detA := targetX * bY - targetY * bX
	A := detA / det
	detB := aX*targetY - aY*targetX
	B := detB / det

  if A != float64(int64(A)) || B != float64(int64(B)) {
		return 0, 0, fmt.Errorf("Non-int")
  }

	return A, B, nil
}

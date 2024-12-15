package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
  file, err := os.Open("input.txt")
  // file, err := os.Open("test.txt")
  // file, err := os.Open("small.txt")
  if err != nil {
    log.Fatal("FAIL")
  }

  scanner := bufio.NewScanner(file)
  y := 0
  grid := make([][]rune, 0)
  moves := make([]rune, 0)
  robot := struct{ x int; y int; }{}
  mode := 0
  for scanner.Scan() {
    line := scanner.Text()

    x := 0
    if mode == 0 {
      grid = append(grid, make([]rune, 0))
      // grid = append(grid, []rune(line))
      for _, char := range line {
        if char == '@' {
          robot.x = x
          robot.y = y
          grid[y] = append(grid[y], '.')
          grid[y] = append(grid[y], '.')
        }
        if char == '#' || char == '.' {
          grid[y] = append(grid[y], char)
          grid[y] = append(grid[y], char)
        }
        if char == 'O' {
          grid[y] = append(grid[y], '[')
          grid[y] = append(grid[y], ']')
        }
        x += 2
      }
    }
    if strings.TrimSpace(line) == "" {
      mode = 1
    }
    if mode == 1 {
      for _, move := range line {
        moves = append(moves, move)
      }
    }
    y++
  }

  grid[robot.y][robot.x] = '.'
  log.Print("Start")
  PrintGrid(grid, robot)
  i := 0
  for len(moves) > 0 {
    move := moves[0]
    moves = moves[1:]
    fmt.Printf("Doing move (%d): %c\n", i, move)
    if move == '^' {
      nextPos := [2]int{robot.x, robot.y - 1}
      if !IsValidPos(grid, nextPos) {
        continue
      }

      if grid[nextPos[1]][nextPos[0]] == ']' || grid[nextPos[1]][nextPos[0]] == '[' {
        boxPositions := DetectBoxes(grid, nextPos, true)
        fmt.Printf("\nGonna move: %v\n", boxPositions)
        if CanMoveBoxes(grid, boxPositions, true) {
          fmt.Printf("Can move\n")
          MoveBoxes(grid, boxPositions, true, robot)
          robot.y -= 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == 'O' {
        movePos := [2]int{robot.x, robot.y - 2}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[1] -= 1
        }
        if IsValidPos(grid, movePos) {
          grid[nextPos[1]][nextPos[0]] = '.'
          grid[movePos[1]][movePos[0]] = 'O'
          robot.y -= 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == '.' {
        grid[robot.y][robot.x] = '.'
        robot.y -= 1
      }
    }

    if move == 'v' {
      nextPos := [2]int{robot.x, robot.y + 1}

      if !IsValidPos(grid, nextPos) {
        continue
      }
      if grid[nextPos[1]][nextPos[0]] == ']' || grid[nextPos[1]][nextPos[0]] == '[' {
        boxPositions := DetectBoxes(grid, nextPos, false)
        fmt.Printf("Detected pos: %v\n", boxPositions)
        if CanMoveBoxes(grid, boxPositions, false) {
          fmt.Printf("Can move\n")
          MoveBoxes(grid, boxPositions, false, robot)
          robot.y += 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == 'O' {
        movePos := [2]int{robot.x, robot.y + 2}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[1] += 1
        }
        if IsValidPos(grid, movePos) {
          grid[nextPos[1]][nextPos[0]] = '.'
          grid[movePos[1]][movePos[0]] = 'O'
          robot.y += 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == '.' {
        grid[robot.y][robot.x] = '.'
        robot.y += 1
      }
    }

    if move == '>' {
      nextPos := [2]int{robot.x + 1, robot.y}
      if !IsValidPos(grid, nextPos) {
        continue
      }
      if grid[nextPos[1]][nextPos[0]] == '[' {
        movePos := [2]int{robot.x + 3, robot.y}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[0] += 1
        }
        if IsValidPos(grid, movePos) {
          for i := 0; i < movePos[0] - nextPos[0]; i++ {
            newX := nextPos[0] + 1 + i
            if i % 2 == 0 {
              grid[nextPos[1]][newX] = '['
            } else {
              grid[nextPos[1]][newX] = ']'
            }
          }
          grid[nextPos[1]][nextPos[0]] = '.'
          robot.x += 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == 'O' {
        movePos := [2]int{robot.x + 2, robot.y}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[0] += 1
        }
        if IsValidPos(grid, movePos) {
          grid[nextPos[1]][nextPos[0]] = '.'
          grid[movePos[1]][movePos[0]] = 'O'
          robot.x += 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == '.' {
        grid[robot.y][robot.x] = '.'
        robot.x += 1
      }
    }

    if move == '<' {
      nextPos := [2]int{robot.x - 1, robot.y}
      if !IsValidPos(grid, nextPos) {
        continue
      }

      if grid[nextPos[1]][nextPos[0]] == ']' {
        movePos := [2]int{robot.x - 3, robot.y}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[0] -= 1
        }
        fmt.Printf("Moving pos: n %d m %d\n", nextPos[0], movePos[0])
        if IsValidPos(grid, movePos) {
          fmt.Printf("Looks valid\n")
          for i := 0; i < nextPos[0] - movePos[0]; i++ {
            newX := nextPos[0] - 1 - i
            fmt.Printf("replacing at: %d\n", newX)
            if i % 2 == 0 {
              grid[nextPos[1]][newX] = ']'
            } else {
              grid[nextPos[1]][newX] = '['
            }
          }
          grid[nextPos[1]][nextPos[0]] = '.'
          robot.x -= 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == 'O' {
        movePos := [2]int{robot.x - 2, robot.y}
        for IsValidPos(grid, movePos) && grid[movePos[1]][movePos[0]] != '.' {
          movePos[0] -= 1
        }
        if IsValidPos(grid, movePos) {
          grid[nextPos[1]][nextPos[0]] = '.'
          grid[movePos[1]][movePos[0]] = 'O'
          robot.x -= 1
        }
      } else if grid[nextPos[1]][nextPos[0]] == '.' {
        grid[robot.y][robot.x] = '.'
        robot.x -= 1
      }
    }
    PrintGrid(grid, robot)
    i++
  }

  p1 := 0
  for y, line := range grid {
    for x, char := range line {
      if char == '[' {
        p1 += x + y * 100
      }
    }
  }

  log.Printf("P1: %d", p1)
}

func IsValidPos(grid [][]rune, pos[2]int) bool {
  return pos[0] >= 0 && pos[1] >= 0 && pos[1] < len(grid) && pos[0] < len(grid[pos[1]]) && grid[pos[1]][pos[0]] != '#'
}

func DetectBoxes(grid [][]rune, pos[2]int, IsUp bool) [][2]int {
  toCheck := [][2]int{pos}
  checked := make(map[string]struct{})

  if grid[pos[1]][pos[0]] == ']' {
    toCheck = append(toCheck, [2]int{pos[0] - 1, pos[1]})
  } else {
    toCheck = append(toCheck, [2]int{pos[0] + 1, pos[1]})
  }

  fmt.Printf("Initial: %v", toCheck)
  boxPositions := [][2]int{}
  for len(toCheck) > 0 {
    curPos := toCheck[0]
    toCheck = toCheck[1:]
    if _, exists := checked[fmt.Sprintf("%d_%d", curPos[0], curPos[1])]; exists {
      continue
    }

    checked[fmt.Sprintf("%d_%d", curPos[0], curPos[1])] = struct{}{}
    boxPositions = append(boxPositions, curPos)
    posToCheck := [2]int{curPos[0], curPos[1]}
    if IsUp {
      posToCheck[1] -= 1
    } else {
      posToCheck[1] += 1
    }

    if IsValidPos(grid, posToCheck) && (grid[posToCheck[1]][posToCheck[0]] == ']') {
      toCheck = append(toCheck, posToCheck)
      toCheck = append(toCheck, [2]int{posToCheck[0] - 1, posToCheck[1]})
    }

    if IsValidPos(grid, posToCheck) && (grid[posToCheck[1]][posToCheck[0]] == '[') {
      toCheck = append(toCheck, posToCheck)
      toCheck = append(toCheck, [2]int{posToCheck[0] + 1, posToCheck[1]})
    }
  }

  return boxPositions
}

func CanMoveBoxes(grid [][]rune, boxPositions [][2]int, IsUp bool) bool {
  for _, boxPos := range boxPositions {
    posToMove := [2]int{boxPos[0], boxPos[1]}
    if IsUp {
      posToMove[1] -= 1
    } else {
      posToMove[1] += 1
    }

    if !IsValidPos(grid, posToMove) {
      return false
    }
  }

  return true
}

func MoveBoxes(grid [][]rune, boxPositions [][2]int, IsUp bool, robot struct{x int; y int;}) {
  fmt.Printf("Moving boxes\n")
  sorted := make([][2]int, len(boxPositions))
  copy(sorted, boxPositions[:])
  sort.Slice(sorted, func(i, j int) bool {
    if IsUp {
      if sorted[i][1] == sorted[j][1] {
        return sorted[i][0] < sorted[j][0]
      }
      return sorted[i][1] > sorted[j][1]
    } else {
      if sorted[i][1] == sorted[j][1] {
        return sorted[i][0] < sorted[j][0]
      }
      return sorted[i][1] < sorted[j][1]
    }
  })

  for _, boxPos := range sorted {
    grid[boxPos[1]][boxPos[0]] = '.'
  }
  fmt.Printf("Reset \n")
  PrintGrid(grid, robot)

  y := 0
  for i, boxPos := range sorted {
    if i == 0 || boxPos[1] != sorted[i][1] {
      y = 0
    }
    if y % 2 == 0 {
      if IsUp {
        grid[boxPos[1] - 1][boxPos[0]] = '['
      } else {
        grid[boxPos[1] + 1][boxPos[0]] = '['
      }
    } else {
      if IsUp {
        grid[boxPos[1] - 1][boxPos[0]] = ']'
      } else {
        grid[boxPos[1] + 1][boxPos[0]] = ']'
      }
    }
    y++
  }
  PrintGrid(grid, robot)
} 

func PrintGrid(grid [][]rune, robot struct{x int; y int}) {
  return
  for y, line := range grid {
    for x, char := range line {
      if x == robot.x && robot.y == y {
        fmt.Print("@")
      } else {
        fmt.Printf("%c", char)
      }
    }
    fmt.Print("\n")
  }
}


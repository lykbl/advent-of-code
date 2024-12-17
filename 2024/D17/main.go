package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  f, _ := os.Open("input.txt")
  // f, _ := os.Open("test.txt")
  // f, _ := os.Open("debug.txt")
  scanner := bufio.NewScanner(f)
  registers := make(map[rune]int, 0)
  commands := make([]int, 0)
  for scanner.Scan() {
    line := scanner.Text()

    r, _ := regexp.Compile("Register ([A-Z]): ([0-9]*)")
    matches := r.FindStringSubmatch(line)
    if len(matches) > 1 {
      register := matches[1]
      registerValue, _ := strconv.Atoi(matches[2])
      registers[rune(register[0])] = registerValue
    }
    if strings.Contains(line, "Program") {
      commandsS := strings.Split(strings.TrimLeft(line, "Program: "), ",")
      for _, command := range commandsS {
        commandI, _ := strconv.Atoi(command)
        commands = append(commands, commandI)
      }
    }
    // for _, match := range matches {
    //   fmt.Printf("Match: %v\n", match)
    // }
  }

  outputs := make([]int, 0)
  i := 0
  for i < len(commands) {
    instruction := commands[i]
    literal := commands[i + 1]
    combo := literal
    var output int
    if literal == 4 {
      combo = registers['A']
    }
    if literal == 5 {
      combo = registers['B']
    }
    if literal == 6 {
      combo = registers['C']
    }

    if instruction == 0 {
      output = registers['A'] / int(math.Pow(2, float64(combo)))
      registers['A'] = output
      fmt.Printf("0: new A: %d\n", registers['A'])
    }
    if instruction == 1 {
      output = registers['B'] ^ literal
      registers['B'] = output
    }
    if instruction == 2 {
      output = combo % 8
      registers['B'] = output
    }
    if instruction == 3 && registers['A'] != 0 {
      i = literal
      fmt.Printf("3: Jumping to: %d\n", i)
      continue
    }
    if instruction == 4 {
      output = registers['B'] ^ registers['C']
      registers['B'] = output
      //Ignore operand?
    }
    if instruction == 5 {
      output = combo % 8
      fmt.Printf("5: Outputting: %d\n", output)
      outputs = append(outputs, output)
    }
    if instruction == 6 {
      output = registers['A'] / int(math.Pow(2, float64(combo)))
      registers['B'] = output
    }
    if instruction == 7 {
      output = registers['A'] / int(math.Pow(2, float64(combo)))
      registers['C'] = output
    }
    i += 2
  }
    fmt.Printf("Registers: %v\n", registers)
    fmt.Printf("Outputs: %v\n", outputs)
}

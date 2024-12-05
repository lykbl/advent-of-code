package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type READ_MODE int
const (
  RULES_MODE = 0
  UPDATES_MODE = 1
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatalf("File error: %s", err)
  }

  scanner := bufio.NewScanner(file)
  mode := 0
  pageRules := make(map[int]*mapset.Set[int], 0)
  log.Println("Loading rules...")
  for scanner.Scan() {
    line := scanner.Text()
    if strings.TrimSpace(line) == "" {
      mode = UPDATES_MODE
      break
    }

    if mode == RULES_MODE {
      rulesParts := strings.Split(line, "|")
      if len(rulesParts) > 2 {
        log.Fatal("Unsupported format")
      }

      before, err := strconv.Atoi(rulesParts[0])
      if err != nil {
        log.Fatal("Atoi failed")
      }
      after, err := strconv.Atoi(rulesParts[1])
      if err != nil {
        log.Fatal("Atoi failed")
      }

      if currentRules, exists := pageRules[before]; exists {
        (*currentRules).Add(after)
      } else {
        rules := mapset.NewSet(after)
        pageRules[before] = &rules
      }
    }
  }
  log.Println("Loading updates...")
  p1 := 0
  for scanner.Scan() {
    line := scanner.Text()
    pagesString := strings.Split(line, ",")

    updateRules := mapset.NewSet[int]()
    ruleValid := true
    for _, pageString := range pagesString {
      page, err := strconv.Atoi(strings.TrimSpace(pageString))
      if err != nil {
        log.Fatal("Atoi failed")
      }
      pagePrintRules, hasRules := pageRules[page]
      if hasRules && !updateRules.Intersect(*pagePrintRules).IsEmpty() {
        log.Printf("Updates invalid: %s", line)
        ruleValid = false
        break
      }

      updateRules.Add(page)
    }

    if ruleValid {
      middlePageString := pagesString[int(math.Ceil(float64(len(pagesString) - 1) / 2))]
      middlePage, err := strconv.Atoi(middlePageString)
      if err != nil {
        log.Fatal("Atoi failed (middlePage)")
      }
      p1 += middlePage
    }
  }

  log.Printf("P1: %d", p1)
}

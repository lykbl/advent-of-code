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

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatalf("File error: %s", err)
  }

  scanner := bufio.NewScanner(file)
  pageRules := make(map[int]*mapset.Set[int], 0)
  log.Println("Loading rules...")
  for scanner.Scan() {
    line := scanner.Text()
    if strings.TrimSpace(line) == "" {
      break
    }

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
  log.Println("Loading updates...")
  p1 := 0
  p2 := 0
  for scanner.Scan() {
    line := scanner.Text()
    pagesString := strings.Split(line, ",")

    currentPages := []int{}
    ruleChanged := false
    for _, pageString := range pagesString {
      page, err := strconv.Atoi(strings.TrimSpace(pageString))
      if err != nil {
        log.Fatal("Atoi failed")
      }
      pagePrintRules, hasRules := pageRules[page]
      currentPages = append(currentPages, page)
      if hasRules == false {
        continue
      }

      pagesToCheck := currentPages
      moved := 0
      for !(*pagePrintRules).Intersect(mapset.NewSet(pagesToCheck...)).IsEmpty() {
        ruleChanged = true
        currentPages[len(currentPages) - 1 - moved], currentPages[len(currentPages) - 2 - moved] = currentPages[len(currentPages) - 2 - moved], currentPages[len(currentPages) - 1 - moved]
        pagesToCheck = pagesToCheck[:len(pagesToCheck) - 1]
        moved++
      }
    }

    if ruleChanged {
      p2 += currentPages[int(math.Ceil(float64(len(currentPages) - 1) / 2))]
    } else {
      p1 += currentPages[int(math.Ceil(float64(len(currentPages) - 1) / 2))]
    }
  }

  log.Printf("P1: %d", p1)
  log.Printf("P2: %d", p2)
}

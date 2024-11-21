package main

import (
	"bufio"
	"container/heap"
	// "crypto/sha256"
	// "encoding/hex"
	"fmt"
	"log"
	"os"
	"pulse/stack"
	"strconv"
	"strings"
)


type Pulse bool
const (
  PulseHigh Pulse = true
  PulseLow Pulse = false
)

func (p Pulse) String() string {
  switch p {
    case true:
      return "high"
    case false:
      return "low"
  }

  panic("Unknown type")
}

func (p Pulse) IntString() string {
  switch p {
    case PulseHigh:
      return "1"
    case PulseLow:
      return "0"
  }

  panic("Unknown Type")
}

type Signal struct {
  from string
  to string
  pulse Pulse
  weight int
}

type PriorityQueue []*Signal
func (h PriorityQueue) Len() int { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].weight > h[j].weight }
func (h PriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
  if signal, ok := x.(*Signal); ok {
    *h = append(*h, signal)
  }
}

func (h *PriorityQueue) Pop() any {
  old := *h
  x := old[len(old) - 1]
  *h = old[0: len(old) - 1]

  return x
}


type Circuit struct {
  modules map[string]Module
  pulsesCounter map[Pulse]int
  modulesStateCache map[string][]int
  buttonPresses int
  cyclesMemory map[string]int
}

type BaseModule struct {
  outputs []string
}

func (m BaseModule) Outputs() []string {
  return m.outputs
}

func (m *BaseModule) AttachOutput(name string) {
  m.outputs = append(m.outputs, name)
}

type Module interface {
  HandleSignal(signal Pulse, from string) *Pulse

  Outputs() []string

  AttachOutput(name string)
}

type FlipFlop struct {
  BaseModule
  state Pulse
}

type Conjunctor struct {
  BaseModule
  lastReceieved map[string]Pulse
  inputsCount int
  signalsReceieved int
}

type Broadcaster struct {
  BaseModule
}

func (m *Broadcaster) HandleSignal(signal Pulse, from string) *Pulse {
  return &signal
}

func (m *FlipFlop) HandleSignal(signal Pulse, from string) *Pulse {
  if signal {
    return nil
  }
  m.state = !m.state

  return &m.state
}

func (m *Conjunctor) HandleSignal(signal Pulse, from string) *Pulse {
  m.signalsReceieved += 1
  m.lastReceieved[from] = signal

  // if m.signalsReceieved != m.inputsCount {
  //   return nil
  // }

  returnSignal := PulseLow
  // if m.signalsReceieved == m.inputsCount {
    // m.signalsReceieved = 0
    for _, storedSignal := range(m.lastReceieved) {
      if storedSignal == PulseLow {
        returnSignal = PulseHigh
        break
      }
    }
  // }

  return &returnSignal
}

func (m *Conjunctor) AttachInput(from string) {
  if m.lastReceieved == nil {
    m.lastReceieved = make(map[string]Pulse)
  }

  m.inputsCount += 1
  m.lastReceieved[from] = false
}

func (c *Circuit) AddModule(name string, m Module){
  c.modules[name] = m
}

func (c *Circuit) PressButton() int {
  c.buttonPresses += 1
  hOld := &PriorityQueue{
    {from: "button", to: "broadcaster", pulse: false, weight: 0},
  }
  path := ""
  heap.Init(hOld)
  h := &stack.Stack[*Signal]{}
  h.Push(&Signal{from: "button", to: "broadcaster", pulse: false, weight: 0})

  for h.Len() > 0 {
    currentSignal, _ := h.Shift()
    path = path + currentSignal.pulse.String()
    c.pulsesCounter[currentSignal.pulse] += 1
    // log.Printf("Next %v", next)
    // if currentSignal, ok := next.(*Signal); ok {
      currentModule := c.modules[currentSignal.to]
      if currentSignal.to == "output" {
        continue
      }
      if currentModule == nil {
        error := fmt.Sprintf("Module not found %s", currentSignal.to)
        panic(error)
      }

      // log.Printf("Handling signal: %s -%s> %v (%d)", currentSignal.from, currentSignal.pulse, currentSignal.to, currentSignal.weight)
      
      if currentSignal.to == "mg" && currentSignal.pulse == PulseHigh {
        savedCycle, ok := c.cyclesMemory[currentSignal.from]
        if !ok || savedCycle == 0 {
          c.cyclesMemory[currentSignal.from] = c.buttonPresses
          log.Printf("Cycle %s - %d", currentSignal.from, c.buttonPresses)
        }
      }

      if currentSignal.to == "rx" && currentSignal.pulse == PulseLow {
        return 0
      }
      newPulse  := currentModule.HandleSignal(currentSignal.pulse, currentSignal.from)
      // log.Printf("Result: %s", newPulse)
      if newPulse == nil {
        continue
      }

      // outputsCount := len(currentModule.Outputs())
      for _, nextModule := range(currentModule.Outputs()) {
        // newWeight := currentSignal.weight + len(currentModule.Outputs()) - i
        newWeight := currentSignal.weight
        // log.Printf("New weight %s%s", currentSignal.to, newWeight)
        // heap.Push(h, &Signal{ from: currentSignal.to, to: nextModule, weight: newWeight, pulse: *newPulse })
        h.Push(&Signal{ from: currentSignal.to, to: nextModule, weight: newWeight, pulse: *newPulse })
      }
    // }
  }

  buttonPresses := 1
  internalKey := c.InternalStateKey()
  visitedAt, ok := c.modulesStateCache[internalKey]
  if ok {
    if visitedAt[1] == 0 {
      c.modulesStateCache[internalKey] = []int{visitedAt[0], c.buttonPresses}
    } else {
      log.Printf("LOOP DETECTED: %d - %d: %d!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!11", visitedAt[0], visitedAt[1], visitedAt[1] - visitedAt[0])
      buttonPresses = visitedAt[1] - visitedAt[0]
    }
  } else {
    c.modulesStateCache[internalKey] = []int{c.buttonPresses, 0}
  }

  return buttonPresses
  // log.Printf("Resulting pulses: %v", c.pulsesCounter)
}

func (c *Circuit) InternalStateKey() string {
  stateKey := ""
  for _, module := range c.modules {
    if flipFlop, ok := module.(*FlipFlop); ok {
      stateKey = stateKey + flipFlop.state.IntString()
    }
    if conjunctor, ok := module.(*Conjunctor); ok {
      for _, state := range conjunctor.lastReceieved {
        stateKey = stateKey + state.IntString()
      }
    }
  }

	//  hasher := sha256.New()
	// hasher.Write([]byte(stateKey)) 
	//  hash := hasher.Sum(nil)     // Get the hash as a byte slice
	//
	// return hex.EncodeToString(hash)
  //

  return stateKey
}

func buildCircuit() *Circuit {
  f, err := os.Open("./input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  circuit := &Circuit { modules: make(map[string]Module), pulsesCounter: map[Pulse]int{ true: 0, false: 0}, modulesStateCache: map[string][]int{}, cyclesMemory: map[string]int{} }
  var attachQueue [][]string
  // log.Printf("Cricuit: %v",circuit)
  //
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()

    parts := strings.Split(line, " -> ")
    inputKey, outputNamesString := parts[0], parts[1]

    var m Module
    var inputName string
    if strings.Contains(inputKey, "%") {
      m = &FlipFlop { state: false }
      inputName = inputKey[1:]
    } else if strings.Contains(inputKey, "&") {
      m = &Conjunctor { inputsCount: 0, signalsReceieved: 0  }
      inputName = inputKey[1:]
    } else if strings.Contains(inputKey, "broadcaster") {
      m = &Broadcaster {}
      inputName = "broadcaster"
    }

    circuit.AddModule(inputName, m)
    outputNames := strings.Split(outputNamesString, ",")
    for _, outputName := range outputNames {
      attachQueue = append(attachQueue, []string{inputName, strings.TrimSpace(outputName)})
    }
  }
  rxM := &Broadcaster{}
  circuit.AddModule("rx", rxM)

  for len(attachQueue) > 0 {
    toAttach := attachQueue[0]
    attachQueue = attachQueue[1:]
    inputName, outputName := toAttach[0], toAttach[1]
    circuit.modules[inputName].AttachOutput(outputName)
    if outputName == "mg" {
      circuit.cyclesMemory[inputName] = 0
    }
    if outputM, ok := circuit.modules[outputName].(*Conjunctor); ok {
      outputM.AttachInput(inputName)
    }

    // log.Printf("%s(%T) sends to %s(%T)\n", inputName, circuit.modules[inputName], outputName, circuit.modules[outputName])
  }
  // log.Printf("CircuitMap: %v", circuit.modules)

  return circuit
}

func main() {
  // test()
  // testStack()
  //
  //
  if len(os.Args) < 2 {
    log.Fatalln("Arg missing")
  }

  pressTimesArg := os.Args[1]

  pressTimes, err := strconv.Atoi(pressTimesArg)
  log.Printf("%d", pressTimes)
  if err != nil {
    log.Fatalf("Error: %s is not integer\n", pressTimesArg)
  }

  circuit := buildCircuit()

  // for i := 0; i < pressTimes; i++ {
  //   log.Printf("Doing %d\n", i)
  //   circuit.PressButton()
  // }

  i := 0
  for {
    toPress := circuit.PressButton()
    if toPress == 0 {
      break
    }
    i += toPress
    // if i % 10e6 == 0 {
      // log.Printf("Now at: %d", i)
    // }
  }

  log.Printf("Answer is: %d", circuit.pulsesCounter[PulseHigh] * circuit.pulsesCounter[PulseLow])
  log.Printf("Min pressTimes: %d", i)
}

package day20

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"strings"
)

type ModuleType rune

const (
	MBroadcaster ModuleType = 'b'
	MFlipFlop    ModuleType = '%'
	MConjuction  ModuleType = '&'
	MNoType      ModuleType = ' '
)

type Module struct {
	name         string
	t            ModuleType
	on           bool
	received     map[string]bool
	destinations []*Module
}

type Task struct {
	sender   string
	receiver string
	signal   bool
}

type ModuleMap = map[string]*Module

func Solve(input string) {
	rows := utils.SplitRows(input)

	modules := make(ModuleMap)
	destinationNames := make(map[string][]string)

	for _, row := range rows {
		rowSplit := strings.Split(row, " -> ")

		moduleType := rowSplit[0][0]
		moduleName := rowSplit[0][1:]
		if moduleType == 'b' {
			moduleName = "broadcaster"
		}

		destinationNames[moduleName] = strings.Split(rowSplit[1], ", ")

		modules[moduleName] = &Module{
			name:         moduleName,
			on:           false,
			t:            ModuleType(moduleType),
			received:     make(map[string]bool),
			destinations: make([]*Module, len(destinationNames[moduleName])),
		}
	}

	for k, v := range destinationNames {
		for i, mName := range v {
			_, exists := modules[mName]

			if !exists {
				modules[mName] = &Module{
					name:         mName,
					on:           false,
					t:            MNoType,
					received:     make(map[string]bool),
					destinations: make([]*Module, 0),
				}
			}

			modules[k].destinations[i] = modules[mName]
			modules[mName].received[k] = false
		}
	}

	lowPulses := 0
	highPulses := 0
	counter := 1

	for counter >= 1 {
		queue := []Task{{"button", "broadcaster", false}}
		lowPulses++

		for len(queue) > 0 {
			task := queue[0]
			queue = queue[1:]
			newTasks := ProcessTask(&modules, task.sender, task.receiver, task.signal, &lowPulses, &highPulses, counter)
			queue = append(queue, newTasks...)
		}

		lowRxCount := 0
		for _, j := range modules["rx"].received {
			if !j {
				lowRxCount++
			}
		}
		//fmt.Println("Part 2: ", counter+1, lowRxCount)
		if lowRxCount == 1 {
			fmt.Println("Part 2: ", counter)
			break
		}

		if counter == 1000 {
			//fmt.Println("After ", counter, " => Low: ", lowPulses, "; High: ", highPulses)
			fmt.Println("Part 1: ", lowPulses, highPulses, lowPulses*highPulses)
		}
		counter++
	}

}

func ProcessTask(modules *ModuleMap, sender string, receiver string, pulse bool, lowPulses, highPulses *int, counter int) []Task {
	if receiver == "rx" && !pulse {
		fmt.Println("Part 2!", counter, pulse)
	}
	newTasks := make([]Task, 0)
	rec := (*modules)[receiver]
	rec.received[sender] = pulse
	result := false

	if rec.t == MConjuction {
		allHigh := true
		for _, j := range rec.received {
			if j == false {
				allHigh = false
				break
			}
		}
		result = !allHigh
	} else if rec.t == MFlipFlop {
		if pulse {
			return newTasks
		} else {
			rec.on = !rec.on
			result = rec.on
		}
	}

	for _, j := range rec.destinations {
		newTasks = append(newTasks, Task{
			sender:   rec.name,
			receiver: j.name,
			signal:   result,
		})
	}
	if result {
		*highPulses += len(rec.destinations)
	} else {
		*lowPulses += len(rec.destinations)
	}

	return newTasks
}

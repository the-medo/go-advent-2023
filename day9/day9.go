package day9

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"strconv"
	"strings"
)

type HistoryRow = []int

type HistorySteps struct {
	steps     []HistoryRow
	newNumber []int
}

func Solve(input string) {
	instructionRows := utils.SplitRows(input)

	historySteps := make([]HistorySteps, len(instructionRows))
	for i, row := range instructionRows {
		numSplits := strings.Split(row, " ")
		historyRow := make(HistoryRow, len(numSplits))
		for j, num := range numSplits {
			historyRow[j], _ = strconv.Atoi(num)
		}

		historySteps[i].steps = []HistoryRow{historyRow}
		allZeros := false
		step := 0
		for !allZeros {
			diffs, x := getDiffs(historySteps[i].steps[step])
			historySteps[i].steps = append(historySteps[i].steps, diffs)
			step, allZeros = step+1, x
		}
	}

	for _, historyStep := range historySteps {
		fmt.Println("hs:", historyStep)
	}

	part1 := 0
	for _, historyStep := range historySteps {
		newNum := findNum(historyStep, false)
		part1 += newNum
		fmt.Println("hs:", newNum)
	}
	fmt.Println("Part1: ", part1)

	part2 := 0
	for _, historyStep := range historySteps {
		newNum := findNum(historyStep, true)
		part2 += newNum
		fmt.Println("hs:", newNum)
	}
	fmt.Println("Part2: ", part2)

}

func findNum(input HistorySteps, isPrev bool) int {
	length := len(input.steps)
	lastNumber := 0
	if isPrev == false {
		for i := length - 1; i >= 0; i-- {
			stepI := input.steps[i]
			lastNumber = stepI[len(stepI)-1] + lastNumber
			input.steps[i] = append(stepI, lastNumber)
		}
	} else {
		for i := length - 1; i >= 0; i-- {
			stepI := input.steps[i]
			lastNumber = stepI[0] - lastNumber
			input.steps[i] = append([]int{lastNumber}, input.steps[i]...)
		}
	}

	return lastNumber
}

func getDiffs(input []int) ([]int, bool) {
	result := make([]int, len(input)-1)
	areZeroes := true

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		result[i-1] = diff
		if diff != 0 {
			areZeroes = false
		}
	}

	return result, areZeroes
}

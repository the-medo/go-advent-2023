package day15

import (
	"fmt"
	"strings"
)

func Solve(input string) {
	inputs := strings.Split(input, ",")
	total := 0
	for _, s := range inputs {
		val := hastString(s)
		total += val
		fmt.Println(s, " = ", val, total)
	}

	fmt.Println("Part 1: ", total)
}

func hastString(s string) int {
	currVal := 0
	for _, c := range s {
		currVal += int(c)
		currVal *= 17
		currVal = currVal % 256
	}

	return currVal
}

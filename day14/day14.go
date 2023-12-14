package day14

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

const ROUND = 'O'
const SQUARE = '#'
const EMPTY = '.'

type RockMap = [][]rune

func Solve(input string) {
	rows := utils.SplitRows(input)

	rockMap := make(RockMap, len(rows))
	for y, row := range rows {
		rockMap[y] = make([]rune, len(row))
		for x, c := range row {
			rockMap[y][x] = c
		}
	}

	moveMap(&rockMap)
	printMap(&rockMap)

	fmt.Println("Part 1: ", countMap(&rockMap))

}

func countMap(rockMap *RockMap) int {
	totalHeight := len(*rockMap)
	result := 0
	for y, row := range *rockMap {
		for _, c := range row {
			if c == ROUND {
				result += totalHeight - y
			}
		}
	}
	return result
}

func printMap(rockMap *RockMap) {
	for _, row := range *rockMap {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func moveMap(rockMap *RockMap) {
	for y, row := range *rockMap {
		for x, c := range row {
			if y > 0 && c == ROUND {
				moveUp(rockMap, x, y)
			}
		}
	}
}

func moveUp(rockMap *RockMap, x, y int) {
	for cY := y - 1; cY >= 0; cY-- {
		if (*rockMap)[cY][x] == EMPTY && (*rockMap)[cY+1][x] == ROUND {
			(*rockMap)[cY][x], (*rockMap)[cY+1][x] = (*rockMap)[cY+1][x], (*rockMap)[cY][x]
		}
	}
}

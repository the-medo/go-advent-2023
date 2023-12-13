package day13

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

func Solve(input string) {
	maps := utils.SplitByEmptyRow(input)

	total := 0

	for i, mp := range maps {
		mapRows := utils.SplitRows(mp)
		mapCols := make([]string, len(mapRows[0]))
		for _, mr := range mapRows {
			for x, c := range mr {
				mapCols[x] = mapCols[x] + string(c)
			}
		}

		fmt.Println("Map ", i)
		fmt.Println("Rows ", mapRows)
		fmt.Println("Cols ", mapCols)

		rowPoints := compare(mapRows, 100)
		colPoints := compare(mapCols, 1)
		fmt.Println("Result rows", rowPoints)
		fmt.Println("Result cols", colPoints)
		fmt.Println("=============")
		total += colPoints + rowPoints
	}
	fmt.Println("Part 1 ", total)
}

func compare(arr []string, multiplier int) int {

	mapTotal := 0

	for j := 0; j < len(arr)-1; j++ {
		k := len(arr) - 1

		found := true
		if arr[j] == arr[k] && j != k {
			for diff := 1; diff <= k-j; diff++ {
				if arr[j+diff] != arr[k-diff] || j+diff == k-diff {
					found = false
					break
				}
			}
		} else {
			found = false
		}
		if found {
			mapTotal = j + ((k - j) / 2) + 1
			fmt.Println("Found (part1)! j + ((k - j) / 2;;;", j, k, mapTotal)
		}
	}

	if mapTotal == 0 {

		for j := len(arr) - 1; j > 0; j-- {
			k := 0

			found := true
			if arr[j] == arr[k] && j != k {

				for diff := 1; diff <= j-k; diff++ {
					if arr[j-diff] != arr[k+diff] || j-diff == k+diff {
						found = false
						break
					}
				}
			} else {
				found = false
			}
			if found {
				mapTotal = k + ((j - k) / 2) + 1
				fmt.Println("Found (part2)!  k + ((j - k) / 2);;;", j, k, mapTotal)
			}
		}
	}

	return mapTotal * multiplier
}

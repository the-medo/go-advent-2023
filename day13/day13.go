package day13

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

func Solve(input string) {
	maps := utils.SplitByEmptyRow(input)

	total1 := 0
	total2 := 0

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

		rowPoints1 := compare(mapRows, 100, 0)
		colPoints1 := compare(mapCols, 1, 0)

		rowPoints2 := compare(mapRows, 100, 1)
		colPoints2 := compare(mapCols, 1, 1)
		fmt.Println("Result rows", rowPoints1)
		fmt.Println("Result cols", colPoints1)
		fmt.Println("=============")
		total1 += colPoints1 + rowPoints1
		total2 += colPoints2 + rowPoints2
	}
	fmt.Println("Part 1 ", total1)
	fmt.Println("Part 2 ", total2)
	fmt.Println("Test ", compareDiffs("#.#.###", "#######"))
}

func compare(arr []string, multiplier int, smudges int) int {

	mapTotal := findFromOneSide(arr, 1, smudges)

	if mapTotal == 0 {
		mapTotal = findFromOneSide(arr, -1, smudges)
	}

	return mapTotal * multiplier
}

func findFromOneSide(arr []string, increment int, smudges int) int {
	maxSmudges := smudges * 2
	result := 0
	start := 0
	end := len(arr) - 1
	if increment == -1 {
		start = len(arr) - 1
		end = 0
	}

	fmt.Println("STARTING! ", increment, smudges, "j, end", start, end)
	for j := start; cond(j, end, increment); j += increment {
		totalSmudges := 0
		k := end

		found := true
		totalSmudges += compareDiffs(arr[j], arr[k])
		fmt.Println("IF ", totalSmudges)
		if totalSmudges <= maxSmudges && j != k {
			for diff := 1; diff <= utils.AbsInt(k-j); diff++ {
				pos1 := j + (diff * increment)
				pos2 := k + (-diff * increment)
				totalSmudges += compareDiffs(arr[pos1], arr[pos2])
				if pos1 == pos2 || totalSmudges > maxSmudges {
					found = false
					break
				}
			}
		} else {
			found = false
		}
		if found && totalSmudges == maxSmudges {
			result = res(j, k, increment)
			fmt.Println("Found (part ", increment, ") ;;;", j, k, result)
		}
	}

	return result
}

func res(j, k, inc int) int {
	if inc == 1 {
		return j + ((k - j) / 2) + 1
	}
	return k + ((j - k) / 2) + 1
}

func cond(num, end, inc int) bool {
	if inc == 1 {
		return num < end
	}
	return num > end
}

func compareDiffs(s1, s2 string) int {
	diffChars := 0
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			diffChars++
		}
	}
	fmt.Println("Comparing: ", s1, s2, diffChars)
	return diffChars
}

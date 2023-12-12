package day12

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"strings"
)

type Template struct {
	template    string
	arrangement []int
	length      int
}

type Combo struct {
	combinations []string
	count        int
}
type Combos map[string]Combo

func Solve(input string) {
	rows := utils.SplitRows(input)
	templates := make([]*Template, len(rows))
	comboSum := 0
	for i, row := range rows {
		rowSplit := strings.Split(row, " ")
		templates[i] = &Template{
			template:    rowSplit[0],
			arrangement: utils.StringsToInts(strings.Split(rowSplit[1], ",")),
			length:      len(rowSplit[0]),
		}
		//fmt.Println("Template: ", templates[i])
		combinations := createCombinations(templates[i].length, templates[i].arrangement, templates[i].template)
		comboSum += len(combinations)
		//fmt.Println("Combinations", combinations, len(combinations))
	}

	fmt.Println("Part 1:", comboSum)
}

func createCombinations(length int, arrangement []int, template string) map[string]string {
	arrLen := len(arrangement)
	minLength := arrLen - 1
	movedBy := make([]int, arrLen)
	semiResult := make(map[string]string)
	result := make(map[string]string)

	for i, n := range arrangement {
		minLength += n
		movedBy[i] = 0
	}

	if minLength > length {
		//fmt.Println("minLength > length", arrangement)
		panic("!")
	}
	lastStart := length - minLength
	totalFree := lastStart + 1

	res := make([][]int, 5)
	allRes := make([][]int, 5)
	//final := make([]int, 0)
	//fmt.Println("res ", res)
	for i, _ := range arrangement {
		for _, arr := range allRes {
			res = createArr(totalFree, arr)
			resLen := len(res[0])

			for _, arr := range res {
				if len(arr) == resLen {
					allRes = append(allRes, arr)
				}
			}
			//fmt.Println("=> res ", res, totalFree)
		}

		if i == arrLen-1 {
			for _, arr := range allRes {
				if isValid(arr, totalFree, arrLen) {
					key := utils.JoinIntsToString(arr, "-")
					semiResult[key] = createString(length, arrangement, arr)
					if compareStringWithTemplate(semiResult[key], template) {
						result[key] = semiResult[key]
					}
				}
			}
		}
	}

	//fmt.Println("Positions: 0 - ", lastStart)
	return result
}

func createArr(max int, arr []int) [][]int {
	res := make([][]int, max)
	for i := 0; i < max; i++ {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		res[i] = append(newArr, i)
	}
	return res
}

func compareStringWithTemplate(s string, t string) bool {
	for i, _ := range s {
		if (t[i] == '.' || t[i] == '#') && t[i] != s[i] {
			return false
		}
	}
	return true
}

func isValid(arr []int, totalFree int, arrLen int) bool {
	if len(arr) < arrLen {
		return false
	}
	//fmt.Print("IS valid", arr)
	sum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		if sum >= totalFree {
			//fmt.Println("FALSE")
			return false
		}
	}
	//fmt.Println("TRUE")
	return true
}

func createString(length int, arrangement []int, movedBy []int) string {
	//fmt.Print("CREATE STRING", length, arrangement, movedBy)
	result := make([]rune, length)
	i := 0
	for n, _ := range arrangement {
		for j := 0; j < movedBy[n]; j++ {
			result[i] = '.'
			i++
		}
		for j := 0; j < arrangement[n]; j++ {
			result[i] = '#'
			i++
		}
		if n != len(arrangement)-1 {
			result[i] = '.'
			i++
		}
	}
	for i < length {
		result[i] = '.'
		i++
	}
	//fmt.Println(string(result))
	return string(result)
}

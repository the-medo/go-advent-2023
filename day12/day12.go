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

type Key struct {
	posInTemplate, rangeCount, currentRangeLength int
}

type CountMap = map[Key]int

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
		mp := make(CountMap)
		value := compute(&Key{0, 0, 0}, templates[i].arrangement, templates[i].template, &mp, '.')
		fmt.Println("Template: ", templates[i], value, mp)

		comboSum += value
	}

	fmt.Println("Part 1:", comboSum)
}

func compute(key *Key, amounts []int, template string, cMap *CountMap, lastChar rune) int {
	val, exists := (*cMap)[*key]
	if exists {
		return val
	}

	if key.posInTemplate == len(template) {
		if key.rangeCount == len(amounts) {
			if key.currentRangeLength > 0 {
				if amounts[key.rangeCount-1] == key.currentRangeLength {
					return 1
				}
				return 0
			}
			return 1
		}
		return 0
	}

	sum := 0
	char := template[key.posInTemplate]
	newPosInTemplate := key.posInTemplate + 1

	if (char == '.' || char == '?') && key.currentRangeLength == 0 {
		sum += compute(&Key{
			posInTemplate:      newPosInTemplate,
			rangeCount:         key.rangeCount,
			currentRangeLength: 0,
		}, amounts, template, cMap, '.')

	} else if (char == '.' || char == '?') && key.currentRangeLength > 0 && key.rangeCount <= len(amounts) && amounts[key.rangeCount-1] == key.currentRangeLength {
		sum += compute(&Key{
			posInTemplate:      newPosInTemplate,
			rangeCount:         key.rangeCount,
			currentRangeLength: 0,
		}, amounts, template, cMap, '.')
	}

	if char == '#' || char == '?' {
		newCurrentRangeLength := key.currentRangeLength + 1
		newRangeCount := key.rangeCount
		if lastChar == '.' {
			newRangeCount++
		}

		sum += compute(&Key{
			posInTemplate:      newPosInTemplate,
			rangeCount:         newRangeCount,
			currentRangeLength: newCurrentRangeLength,
		}, amounts, template, cMap, '#')
	}

	(*cMap)[*key] = sum

	return sum
}

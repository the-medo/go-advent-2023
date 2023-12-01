package day1

import (
	"github.com/the-medo/go-advent-2023/utils"
	"strings"
)

func Solve(input string) {
	rows := utils.SplitRows(input)

	println("Part1: ", computeSum(rows))

	mapping := map[string]string{
		"one":   "1e",
		"two":   "2o",
		"three": "3e",
		"four":  "4r",
		"five":  "5e",
		"six":   "6x",
		"seven": "7n",
		"eight": "8t",
		"nine":  "9e",
	}

	for i, _ := range rows {

		lowestIndex := 0

		for lowestIndex != -1 {
			lowestIndexReplacement := "one"
			lowestIndex = -1

			for k, _ := range mapping {
				index := strings.Index(rows[i], k)
				if index > -1 && (index < lowestIndex || lowestIndex == -1) {
					lowestIndexReplacement = k
					lowestIndex = index
				}
			}
			if lowestIndex > -1 {
				rows[i] = strings.Replace(rows[i], lowestIndexReplacement, mapping[lowestIndexReplacement], -1)
			}

		}
	}

	println("Part2: ", computeSum(rows))

}

func computeSum(rows []string) int32 {
	sum := int32(0)
	for _, row := range rows {
		firstDigit := int32(-1)
		lastDigit := int32(-1)
		for _, c := range row {
			if c >= 48 && c <= 57 {
				digit := c - 48
				if firstDigit == -1 {
					firstDigit = digit
				}
				lastDigit = digit
			}
		}
		if firstDigit != -1 {
			sum += (firstDigit * 10) + lastDigit
		}
	}

	return sum
}

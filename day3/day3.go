package day3

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

const RUNE_DOT = rune('.')
const RUNE_DIGIT_START = rune('0')
const RUNE_DIGIT_END = rune('9')

type SchemaRow = []rune
type Schema = []SchemaRow
type PartNumber struct {
	number          int
	adjacentSymbols []rune
}

func Solve(input string) {

	rows := utils.SplitRows(input)

	schema := make(Schema, len(rows))

	for i, row := range rows {
		schema[i] = make(SchemaRow, len(row))
		for j, r := range row {
			schema[i][j] = r
		}
	}

	// ========== part 1
	numbers := []PartNumber{}
	for y, _ := range schema {
		isNumber := false
		num := PartNumber{
			adjacentSymbols: make([]rune, 0),
		}
		for x, c := range schema[y] {
			if c >= RUNE_DIGIT_START && c <= RUNE_DIGIT_END {
				if isNumber { //it is not first digit of a number
					num.number = num.number*10 + runeNumber(c)
				} else { //it is first digit of a number
					isNumber = true
					num.number = runeNumber(c)

					if x > 0 {
						if y > 0 && isSymbol(schema[y-1][x-1]) {
							num.adjacentSymbols = append(num.adjacentSymbols, schema[y-1][x-1])
						}
						if y < len(schema)-1 && isSymbol(schema[y+1][x-1]) {
							num.adjacentSymbols = append(num.adjacentSymbols, schema[y+1][x-1])
						}
						if isSymbol(schema[y][x-1]) {
							num.adjacentSymbols = append(num.adjacentSymbols, schema[y][x-1])
						}
					}
					if y > 0 && isSymbol(schema[y-1][x]) {
						num.adjacentSymbols = append(num.adjacentSymbols, schema[y-1][x])
					}
					if y < len(schema)-1 && isSymbol(schema[y+1][x]) {
						num.adjacentSymbols = append(num.adjacentSymbols, schema[y+1][x])
					}
				}

				if x < len(schema[y])-1 {
					if y > 0 && isSymbol(schema[y-1][x+1]) {
						num.adjacentSymbols = append(num.adjacentSymbols, schema[y-1][x+1])
					}
					if y < len(schema)-1 && isSymbol(schema[y+1][x+1]) {
						num.adjacentSymbols = append(num.adjacentSymbols, schema[y+1][x+1])
					}
					if isSymbol(schema[y][x+1]) {
						num.adjacentSymbols = append(num.adjacentSymbols, schema[y][x+1])
					}
				}
			} else if isNumber {
				isNumber = false
				numbers = append(numbers, num)
				num = PartNumber{
					adjacentSymbols: make([]rune, 0),
				}
			}
		}
		if isNumber {
			numbers = append(numbers, num)
			num = PartNumber{
				adjacentSymbols: make([]rune, 0),
			}
		}
	}

	part1 := 0

	for _, num := range numbers {
		if len(num.adjacentSymbols) > 0 {
			part1 += num.number
		}
	}

	fmt.Println(numbers)
	fmt.Println("Part 1:", part1)
}

func runeNumber(r rune) int {
	return int(r - RUNE_DIGIT_START)
}

func isSymbol(r rune) bool {
	return r != RUNE_DOT && (r < RUNE_DIGIT_START || r > RUNE_DIGIT_END)
}

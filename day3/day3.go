package day3

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

const RUNE_DOT = rune('.')
const RUNE_STAR = rune('*')
const RUNE_DIGIT_START = rune('0')
const RUNE_DIGIT_END = rune('9')

type SchemaRow = []rune
type Schema = []SchemaRow
type PartNumber struct {
	id              int
	number          int
	adjacentSymbols []rune
}
type Gear struct {
	adjacentNumberIds []int
}
type GearMap map[string]*Gear

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
	gearMap := GearMap{}
	numId := 0

	for y, _ := range schema {
		isNumber := false
		numId++
		num := PartNumber{
			id:              numId,
			adjacentSymbols: make([]rune, 0),
		}
		for x, c := range schema[y] {
			if c >= RUNE_DIGIT_START && c <= RUNE_DIGIT_END {
				if isNumber { //it is not first digit of a number
					num.number = num.number*10 + runeNumber(c)
				} else { //it is first digit of a number
					isNumber = true
					num.number = runeNumber(c)

					checkSymbol(schema, y-1, x-1, &num, numId, &gearMap)
					checkSymbol(schema, y, x-1, &num, numId, &gearMap)
					checkSymbol(schema, y+1, x-1, &num, numId, &gearMap)

					checkSymbol(schema, y-1, x, &num, numId, &gearMap)
					checkSymbol(schema, y+1, x, &num, numId, &gearMap)

				}

				checkSymbol(schema, y-1, x+1, &num, numId, &gearMap)
				checkSymbol(schema, y, x+1, &num, numId, &gearMap)
				checkSymbol(schema, y+1, x+1, &num, numId, &gearMap)
			} else if isNumber {
				isNumber = false
				numbers = append(numbers, num)
				numId++
				num = PartNumber{
					id:              numId,
					adjacentSymbols: make([]rune, 0),
				}
			}
		}
		if isNumber {
			numbers = append(numbers, num)
			numId++
			num = PartNumber{
				id:              numId,
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

	part2 := 0
	for k, g := range gearMap {
		if len(g.adjacentNumberIds) == 2 {
			fmt.Print("Gear:", k, " => ")
			gearRatio := 1
			for _, gearNumId := range g.adjacentNumberIds {
				for _, num := range numbers {
					if num.id == gearNumId {
						gearRatio *= num.number
					}
				}
			}
			fmt.Print("Gear ratio = ", gearRatio)
			part2 += gearRatio
		}
		fmt.Println()
	}

	fmt.Println("Part 2:", part2)
}

func runeNumber(r rune) int {
	return int(r - RUNE_DIGIT_START)
}

func checkSymbol(schema Schema, y int, x int, partNum *PartNumber, numId int, gearMap *GearMap) {
	if x < 0 || y < 0 || y >= len(schema) || x >= len(schema[0]) {
		return
	}

	c := schema[y][x]

	if isSymbol(c) {
		partNum.adjacentSymbols = append(partNum.adjacentSymbols, schema[y][x])
	}

	if c == RUNE_STAR {
		key := fmt.Sprintf("%d-%d", x, y)
		gear, exists := (*gearMap)[key]

		if exists {
			gear.adjacentNumberIds = append(gear.adjacentNumberIds, numId)
		} else {
			gear = &Gear{
				adjacentNumberIds: []int{numId},
			}
			(*gearMap)[key] = gear
		}
	}
}

func isSymbol(r rune) bool {
	return r != RUNE_DOT && (r < RUNE_DIGIT_START || r > RUNE_DIGIT_END)
}

package day5

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"strconv"
	"strings"
)

type MapRange struct {
	oneStart  int
	twoStart  int
	rangeSize int
}

type SeedMap = []MapRange

type FixedSeed struct {
	start  int
	length int
}

func Solve(input string) {
	sections := utils.SplitByEmptyRow(input)

	seedNumbersSplit := strings.Split(sections[0], ": ")
	seedsSplit := strings.Split(seedNumbersSplit[1], " ")

	seeds := make([]int, len(seedsSplit))
	for i, s := range seedsSplit {
		seeds[i], _ = strconv.Atoi(s)
	}

	seedMaps := make([]SeedMap, len(sections)-1)
	for i, section := range sections[1:] {
		sectionRows := utils.SplitRows(section)
		seedMap := make(SeedMap, len(sectionRows)-1)
		for j, sectionRow := range sectionRows[1:] {
			rangeSplit := strings.Split(sectionRow, " ")
			twoStart, _ := strconv.Atoi(rangeSplit[0])
			oneStart, _ := strconv.Atoi(rangeSplit[1])
			rangeSize, _ := strconv.Atoi(rangeSplit[2])
			seedMap[j] = MapRange{
				oneStart:  oneStart,
				twoStart:  twoStart,
				rangeSize: rangeSize,
			}
		}
		seedMaps[i] = seedMap
	}

	fmt.Println(seedMaps, seeds)

	lowestLocation := math.MaxInt

	//part 1
	for _, seed := range seeds {
		fmt.Print("Seed: ", seed, " ")
		tmpValue := seed
		for _, seedMap := range seedMaps {
			for _, x := range seedMap {
				if x.oneStart <= tmpValue && x.oneStart+x.rangeSize > tmpValue {

					fmt.Print(" (", tmpValue, " - ", x.oneStart, " ", x.twoStart, " ", x.rangeSize, ") ")
					tmpValue = x.twoStart + (tmpValue - x.oneStart)
					break
				}
			}
			fmt.Print(" => ", tmpValue, " ")
		}
		if tmpValue < lowestLocation {
			lowestLocation = tmpValue
		}
		fmt.Println()
	}

	fmt.Println("Part 1:", lowestLocation)

	//part 2
	//fix seeds

}

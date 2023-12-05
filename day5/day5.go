package day5

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"strconv"
	"strings"
)

type MapRange struct {
	sourceStart int
	targetStart int
	rangeSize   int
}

type SeedMap = []MapRange

type FixedSeedInfo struct {
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
			oneStart, _ := strconv.Atoi(rangeSplit[1])
			twoStart, _ := strconv.Atoi(rangeSplit[0])
			rangeSize, _ := strconv.Atoi(rangeSplit[2])
			seedMap[j] = MapRange{
				sourceStart: oneStart,
				targetStart: twoStart,
				rangeSize:   rangeSize,
			}
		}
		seedMaps[i] = seedMap
	}

	fmt.Println("Part 1:", findLowestLocation(seeds, seedMaps))

	//part 2
	//fix seeds
	fixedSeedInfos := make([]FixedSeedInfo, len(seeds)/2)
	for i, seed := range seeds {
		if i%2 == 0 {
			fixedSeedInfos[i/2] = FixedSeedInfo{
				start: seed,
			}
		} else {
			fixedSeedInfos[i/2].length = seed
		}
	}

	fmt.Println("Fixed seeds:", fixedSeedInfos)
	fmt.Println("First seed map:", seedMaps[0])

	test := make([][]FixedSeedInfo, len(seedMaps)+1)
	test[0] = fixedSeedInfos
	lowestValue := 0
	for i, sm := range seedMaps {
		test[i+1], lowestValue = splitSeedInfos(test[i], sm)
		fmt.Println("Round", i, " [", lowestValue, "] : ", test[i+1])
	}

	//seedsFinal := make([]int, len(test[len(test)-1]))
	//for i, x := range test[len(test)-1] {
	//	seedsFinal[i] = x.start
	//}
	//fmt.Println("Final seeds:", seedsFinal)
	//fmt.Println("Lowest value:", lowestValue)
	//fmt.Println("Part 2:", findLowestLocation(seedsFinal, seedMaps))
}

func splitSeedInfos(fixedSeedInfos []FixedSeedInfo, seedMap SeedMap) ([]FixedSeedInfo, int) {
	lowestValue := math.MaxInt

	newSeedInfo := make([]FixedSeedInfo, 0)

	for _, fsi := range fixedSeedInfos {
		leftover := fsi.length
		tmpValue := fsi.start
		for leftover > 0 {
			//newSeedInfo = append(newSeedInfo, tmpValue)
			found := false
			nearestRangeStart := math.MaxInt
			for _, x := range seedMap {
				if x.sourceStart > tmpValue && x.sourceStart < nearestRangeStart {
					nearestRangeStart = x.sourceStart
				}
				if x.sourceStart <= tmpValue && x.sourceStart+x.rangeSize > tmpValue {
					found = true
					diff := tmpValue - x.sourceStart
					diffTillEnd := x.rangeSize - diff
					newLength := leftover
					leftover -= diffTillEnd
					if leftover < 0 {
						//diff += leftover
					} else {
						newLength = diffTillEnd
						tmpValue = x.sourceStart + x.rangeSize
					}
					newValue := x.targetStart + diff
					newSeedInfo = append(newSeedInfo, FixedSeedInfo{
						start:  newValue,
						length: newLength,
					})
					if newValue < lowestValue {
						lowestValue = newValue
					}
					//println("DIFF: ", diff, " LEFTOVER: ", leftover, " TMP: ", tmpValue)
					break
				}
			}
			if !found {
				//diff := nearestRangeStart - tmpValue
				diff := nearestRangeStart - tmpValue
				if nearestRangeStart == math.MaxInt {
					diff = leftover
				}
				newSeedInfo = append(newSeedInfo, FixedSeedInfo{
					start:  tmpValue,
					length: diff,
				})
				if tmpValue < lowestValue {
					lowestValue = tmpValue
				}
				leftover -= diff
				tmpValue += diff
			}
		}
	}
	//fmt.Println("Seeds to check:", newSeedInfo)
	return newSeedInfo, lowestValue
}

func findLowestLocation(seeds []int, seedMaps []SeedMap) int {
	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		fmt.Print("Seed: ", seed, " ")
		tmpValue := seed
		for _, seedMap := range seedMaps {
			for _, x := range seedMap {
				if x.sourceStart <= tmpValue && x.sourceStart+x.rangeSize > tmpValue {

					fmt.Print(" (", tmpValue, " - ", x.sourceStart, " ", x.targetStart, " ", x.rangeSize, ") ")
					tmpValue = x.targetStart + (tmpValue - x.sourceStart)
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

	return lowestLocation
}

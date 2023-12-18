package day18

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"strconv"
	"strings"
)

const (
	DirUp      = 'U'
	DirDown    = 'D'
	DirRight   = 'R'
	DirLeft    = 'L'
	Horizontal = 1
	Vertical   = 2
	Both       = 3
)

type Point struct {
	x, y int
	//color string
}

type ColRow struct {
	pos             int
	lowest, highest int
	points          map[*Point]bool
}

//
//type Point struct {
//	x, y int
//}

func Solve(input string) {
	inputRows := utils.SplitRows(input)

	trenchRows := make(map[int]*ColRow, 0)
	trenchCols := make(map[int]*ColRow, 0)

	x, y := 0, 0
	highestY, highestX, lowestY, lowestX := -math.MaxInt, -math.MaxInt, math.MaxInt, math.MaxInt
	//pointsToCheck := make([]*Point, 0)

	for _, row := range inputRows {
		splitRow := strings.Split(row, " ")
		dir := splitRow[0][0]
		value, _ := strconv.Atoi(splitRow[1])

		endX, incX := x, 0
		endY, incY := y, 0
		if dir == DirLeft {
			endX, incX = x-value, -1
		} else if dir == DirRight {
			endX, incX = x+value, 1
		} else if dir == DirUp {
			endY, incY = y-value, -1
		} else if dir == DirDown {
			endY, incY = y+value, 1
		}

		for x, y = x, y; x != endX || y != endY; x, y = x+incX, y+incY {

			point := &Point{
				x: x,
				y: y,
				//color: splitRow[2],
			}

			trenchRow, existRow := trenchRows[y]
			trenchCol, existCol := trenchCols[x]

			if !existRow {
				trenchRows[y] = &ColRow{
					pos:     y,
					highest: x,
					lowest:  x,
					points:  make(map[*Point]bool),
				}
				trenchRows[y].points[point] = true
			} else {
				if trenchRow.highest < x {
					trenchRow.highest = x
				}
				if trenchRow.lowest > x {
					trenchRow.lowest = x
				}
				trenchRow.points[point] = true
			}

			if !existCol {
				trenchCols[x] = &ColRow{
					pos:     x,
					highest: y,
					lowest:  y, points: make(map[*Point]bool),
				}
			} else {
				if trenchCol.highest < y {
					trenchCol.highest = y
				}
				if trenchCol.lowest > y {
					trenchCol.lowest = y
				}
				trenchCol.points[point] = true
			}

			if x > highestX {
				highestX = x
			}
			if x < lowestX {
				lowestX = x
			}
			if y > highestY {
				highestY = y
			}
			if y < lowestY {
				lowestY = y
			}
		}
	}

	totalOutside := 0

	fullMap := make([][]int, highestY-lowestY+1+2)

	for _, row := range trenchRows {
		y = row.pos - lowestY + 1
		fullMap[y] = make([]int, highestX-lowestX+1+2)
		for k, _ := range row.points {
			fullMap[y][k.x-lowestX+1] = 1
		}
	}

	fullMap[0] = make([]int, highestX-lowestX+1+2)
	fullMap[len(fullMap)-1] = make([]int, highestX-lowestX+1+2)

	queue := []Point{{0, 0}}

	fmt.Println(fullMap)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		//fmt.Println(cur)
		if fullMap[cur.y][cur.x] == 0 {
			totalOutside++
			fullMap[cur.y][cur.x] = 2

			if cur.x > 0 {
				queue = append(queue, Point{x: cur.x - 1, y: cur.y})
			}
			if cur.y > 0 {
				queue = append(queue, Point{x: cur.x, y: cur.y - 1})
			}
			if cur.y < len(fullMap)-1 {
				queue = append(queue, Point{x: cur.x, y: cur.y + 1})
			}
			if cur.x < len(fullMap[0])-1 {
				queue = append(queue, Point{x: cur.x + 1, y: cur.y})
			}

		}
	}

	fmt.Println("Outside: ", totalOutside)

	fmt.Println("Total: ", (len(fullMap)*len(fullMap[0]))-totalOutside)

	//for _, row := range trenchRows {
	//	started := true
	//	yArray := make([]int, len(row.points))
	//
	//	i := 0
	//	for k, _ := range row.points {
	//		yArray[i] = k.x
	//		fmt.Print(k, "; ")
	//		i++
	//	}
	//
	//	sort.Ints(yArray)
	//
	//	//lastVal := yArray[0]
	//	continuousLine := false
	//	rowTotal := 0
	//	for i := 0; i < len(yArray); i++ {
	//		val := yArray[i]
	//
	//		if !started {
	//			fmt.Print(val, "S;")
	//			started = true
	//			rowTotal += 1
	//		} else {
	//			if i == len(yArray)-1 {
	//				rowTotal += 1
	//			} else {
	//				diff := yArray[i+1] - val + 1
	//				if diff > 1 && !continuousLine {
	//					fmt.Print(val, "D-NL; ")
	//					i++
	//					started = false
	//					rowTotal += diff
	//				} else if diff > 1 && continuousLine {
	//					fmt.Print(val, "D-L; ")
	//					i++
	//					started = true
	//					rowTotal += 0
	//					continuousLine = false
	//				} else if diff == 1 {
	//					fmt.Print(val, "D=1-L; ")
	//					continuousLine = true
	//					rowTotal += 1
	//				} else {
	//					fmt.Print(val, "OOOPS; ")
	//				}
	//			}
	//		}
	//	}
	//	total += rowTotal
	//
	//	fmt.Println(" => ", rowTotal)
	//}

	//fmt.Println("Trench rows: ", trenchRows)
	//fmt.Println("Part 0.5: ", total)
}

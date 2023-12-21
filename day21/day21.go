package day21

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
)

type Point struct {
	x, y      int
	val       rune
	visitedAt map[string]int
}

type Garden = [][]*Point

type Task struct {
	p      *Point
	gX, gY int
}

func Solve(input string) {
	rows := utils.SplitRows(input)
	size := len(rows)

	g := make(Garden, len(rows))
	start := &Point{}

	freePoints := 0

	for y, row := range rows {
		g[y] = make([]*Point, len(row))
		for x, c := range row {
			visitedAt := make(map[string]int)
			g[y][x] = &Point{
				x:         x,
				y:         y,
				val:       c,
				visitedAt: visitedAt,
			}

			if c == 'S' || c == '.' {
				freePoints++
			}

			if c == 'S' {
				(*g[y][x]).val = '.'
				(*g[y][x]).visitedAt["0;0"] = 0
				*start = *g[y][x]
				fmt.Println("START!", *start, *g[y][x])
			}
		}
	}

	fmt.Println("Free points: ", freePoints)

	queue := []*Task{{start, 0, 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		newPoints := p.Process(&g, size)
		queue = append(queue, newPoints...)
	}

	stepCountPart1 := 64
	totalPoints1 := 0
	rowVals := make([]map[int]int, len(g))
	for y, row := range g {
		rowVals[y] = make(map[int]int)
		for _, p := range row {
			mainVal := math.MaxInt
			for i, x := range p.visitedAt {
				if i == k(0, 0) {
					mainVal = p.visitedAt[k(0, 0)]
					_, exists := rowVals[y][x%2]
					if !exists {
						rowVals[y][x%2] = 1
					} else {
						rowVals[y][x%2]++
					}
				}
				if x%2 == stepCountPart1%2 && x <= stepCountPart1 {
					totalPoints1++
				}
			}
			p.visitedAt = make(map[string]int)
			p.visitedAt[k(0, 0)] = mainVal
			//fmt.Println(p)
		}
	}

	fmt.Println(rowVals)

	stepCountPart2 := 26501365

	middleST := stepCountPart2 % 2
	yVal := 0
	startW := (stepCountPart2 * 2) + 1
	w := startW
	totalPoints2 := 0

	fmt.Println(*start)
	plusCache, minusCache := make(map[string]int), make(map[string]int)

	for w > 0 {
		rowPointsMinus := 0
		rowPointsPlus := 0
		//fmt.Println("================ W = ", w, " ====================== ")
		widthWithoutMain := w - size
		actualST := (middleST + 1) % 2
		atTheTip := false
		if widthWithoutMain > 0 {
			//fmt.Println("New: ", widthWithoutMain, "[+", rowVals[(*start).y+yVal][middleST], " ; - ", rowVals[(*start).y-yVal][middleST], "]")
			rowPointsPlus += rowVals[(*start).y+yVal][middleST]
			rowPointsMinus += rowVals[(*start).y-yVal][middleST]
		} else {
			widthWithoutMain = w
			//fmt.Println("Staying at old: ", widthWithoutMain)
			actualST = (actualST + 1) % 2
			atTheTip = true
		}

		repeats := widthWithoutMain / (size * 2)
		//fmt.Println("Repeats: ", repeats)
		if repeats > 0 {
			actualSTCount := repeats/2 + repeats%2
			nextSTCount := repeats - actualSTCount

			rowPointsPlus += rowVals[(*start).y+yVal][actualST] * 2 * actualSTCount
			rowPointsMinus += rowVals[(*start).y-yVal][actualST] * 2 * actualSTCount
			actualST = (actualST + 1) % 2
			rowPointsPlus += rowVals[(*start).y+yVal][actualST] * 2 * nextSTCount
			rowPointsMinus += rowVals[(*start).y-yVal][actualST] * 2 * nextSTCount

			if actualSTCount == nextSTCount {
				actualST = (actualST + 1) % 2
			}
		}
		repeatsModulo := widthWithoutMain % (size * 2)

		bothDirections := repeatsModulo / 2
		//fmt.Println("Repeats modulo: ", repeatsModulo, "Both directions: ", bothDirections)

		cacheKey := k2(bothDirections, yVal, actualST)
		rPlus, ePlus := plusCache[cacheKey]
		rMinus, eMinus := minusCache[cacheKey]
		if ePlus && eMinus && !atTheTip {
			rowPointsPlus += rPlus
			rowPointsMinus += rMinus
		} else {
			pointsP, pointsM := 0, 0

			for i := 0; i < bothDirections; i++ {
				offset1, offset2 := size-1-i, i
				if atTheTip {
					offset1, offset2 = (*start).x+i+1, (*start).x-i-1
				}

				val := g[(*start).y+yVal][offset1].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsP++
				}
				val = g[(*start).y+yVal][offset2].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsP++
				}
				val = g[(*start).y-yVal][offset1].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsM++
				}
				val = g[(*start).y-yVal][offset2].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsM++
				}
			}

			if atTheTip {
				val := g[(*start).y+yVal][(*start).x].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsP++
				}
				val = g[(*start).y-yVal][(*start).x].visitedAt[k(0, 0)]
				if val < math.MaxInt && val%2 == actualST%2 {
					pointsM++
				}
			} else {
				plusCache[cacheKey] = pointsP
				minusCache[cacheKey] = pointsM
			}

			rowPointsPlus += pointsP
			rowPointsMinus += pointsM
		}

		totalPoints2 += rowPointsPlus + rowPointsMinus
		if w == startW {
			totalPoints2 -= rowPointsPlus
		}

		//if (w-1)%100000 == 0 {
		//	fmt.Println("W", w, "yVal", yVal, ", total: ", totalPoints2, "rowPlus:", rowPointsPlus, "rowMinus:", rowPointsMinus)
		//}

		yVal++

		if yVal > ((size - 1) / 2) {
			yVal = -((size - 1) / 2)
			middleST = (middleST + 1) % 2
		}

		w -= 2
		//break
	}

	fmt.Println("Part 1: ", stepCountPart1, totalPoints1-1)
	fmt.Println("Part 2: ", stepCountPart2, totalPoints2)
}

func (t *Task) Process(garden *Garden, maxSteps int) []*Task {
	newTasks := make([]*Task, 0)
	step := t.p.visitedAt[k(t.gX, t.gY)]
	if step+1 > maxSteps {
		return newTasks
	}

	valid, gx, gy, point := ValidPoint(garden, t.p.x-1, t.p.y, step, t.gX, t.gY)
	if valid {
		newTasks = append(newTasks, &Task{point, gx, gy})
	}
	valid, gx, gy, point = ValidPoint(garden, t.p.x+1, t.p.y, step, t.gX, t.gY)
	if valid {
		newTasks = append(newTasks, &Task{point, gx, gy})
	}
	valid, gx, gy, point = ValidPoint(garden, t.p.x, t.p.y-1, step, t.gX, t.gY)
	if valid {
		newTasks = append(newTasks, &Task{point, gx, gy})
	}
	valid, gx, gy, point = ValidPoint(garden, t.p.x, t.p.y+1, step, t.gX, t.gY)
	if valid {
		newTasks = append(newTasks, &Task{point, gx, gy})
	}

	return newTasks
}

func ValidPoint(garden *Garden, x, y int, step int, gX, gY int) (bool, int, int, *Point) {
	if y < 0 {
		y = len(*garden) - 1
		gY--
	}
	if x < 0 {
		x = len((*garden)[0]) - 1
		gX--
	}
	if y >= len(*garden) {
		y = 0
		gY++
	}
	if x >= len((*garden)[0]) {
		x = 0
		gX++
	}

	point := (*garden)[y][x]
	if point.val == '#' {
		return false, 0, 0, nil
	}
	_, exists := point.visitedAt[k(gX, gY)]
	if exists {
		return false, 0, 0, nil
	}

	point.visitedAt[k(gX, gY)] = step + 1
	return true, gX, gY, point
}

func k(n1, n2 int) string {
	return fmt.Sprintf("%d-%d", n1, n2)
}

func k2(n1, n2, n3 int) string {
	return fmt.Sprintf("%d;%d;%d", n1, n2, n3)
}

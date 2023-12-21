package day21

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
)

type Point struct {
	x, y      int
	val       rune
	visitedAt int
}

type Garden = [][]*Point

func Solve(input string) {
	rows := utils.SplitRows(input)

	g := make(Garden, len(rows))
	start := &Point{}

	for y, row := range rows {
		g[y] = make([]*Point, len(row))
		for x, c := range row {
			g[y][x] = &Point{
				x:         x,
				y:         y,
				val:       c,
				visitedAt: math.MaxInt,
			}

			if c == 'S' {
				(*g[y][x]).val = '.'
				(*g[y][x]).visitedAt = 0
				*start = *g[y][x]
				fmt.Println("START!", *start, *g[y][x])
			}
		}
	}

	queue := []*Point{start}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		newPoints := p.Process(&g)
		queue = append(queue, newPoints...)
	}

	fmt.Println(start)

	stepCount := 64
	totalPoints := 0
	for _, row := range g {
		for _, p := range row {
			if p.visitedAt == math.MaxInt {
				fmt.Print("x")
			} else {
				fmt.Print(p.visitedAt)
			}
			//fmt.Print(string(p.val))
			if p.visitedAt < math.MaxInt && p.visitedAt%2 == stepCount%2 && p.visitedAt <= stepCount {
				totalPoints++
			}
		}
		fmt.Println()
	}

	fmt.Println("Part 1: ", totalPoints)
}

func (p *Point) Process(garden *Garden) []*Point {
	newPoints := make([]*Point, 0)

	if ValidPoint(garden, p.x-1, p.y, p.visitedAt) {
		newPoints = append(newPoints, (*garden)[p.y][p.x-1])
	}
	if ValidPoint(garden, p.x+1, p.y, p.visitedAt) {
		newPoints = append(newPoints, (*garden)[p.y][p.x+1])
	}
	if ValidPoint(garden, p.x, p.y-1, p.visitedAt) {
		newPoints = append(newPoints, (*garden)[p.y-1][p.x])
	}
	if ValidPoint(garden, p.x, p.y+1, p.visitedAt) {
		newPoints = append(newPoints, (*garden)[p.y+1][p.x])
	}

	return newPoints
}

func ValidPoint(garden *Garden, x, y int, step int) bool {
	if y < 0 || y >= len(*garden) || x < 0 || x >= len((*garden)[0]) {
		return false
	}
	point := (*garden)[y][x]
	if point.val == '#' {
		return false
	}
	if point.visitedAt < math.MaxInt {
		return false
	}
	point.visitedAt = step + 1
	return true
}

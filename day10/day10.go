package day10

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

type Dirs struct {
	x, y int
}

type Neighbors = map[rune][]Dirs
type TileMap = map[string]*Point

type Point struct {
	x, y  int
	value rune
}

var NeighborMap = Neighbors{
	'F': []Dirs{{0, 1}, {1, 0}},
	'L': []Dirs{{0, -1}, {1, 0}},
	'J': []Dirs{{-1, 0}, {0, -1}},
	'7': []Dirs{{-1, 0}, {0, 1}},
	'-': []Dirs{{-1, 0}, {1, 0}},
	'|': []Dirs{{0, -1}, {0, 1}},
	',': []Dirs{},
}

func Solve(input string) {
	rows := utils.SplitRows(input)

	var startingPoint *Point
	tiles := &TileMap{}

	for y, row := range rows {
		for x, c := range row {
			(*tiles)[key(x, y)] = &Point{
				x:     x,
				y:     y,
				value: c,
			}
			if c == 'S' {
				startingPoint = (*tiles)[key(x, y)]
			}
		}
	}

	fmt.Println("Start: ", startingPoint)

	startingPoint.value = getStartingShape(tiles, startingPoint.x, startingPoint.y)
	fmt.Println("Part1: ", loopNeighbors(startingPoint, tiles, startingPoint, startingPoint, 0)/2)
}

func loopNeighbors(start *Point, tiles *TileMap, cp *Point, pp *Point, length int) int {
	if cp.x == start.x && cp.y == start.y && length > 0 {
		return length
	}
	diffX, diffY := pp.x-cp.x, pp.y-cp.y
	neighbors := NeighborMap[cp.value]
	for _, n := range neighbors {
		if n.x != diffX || n.y != diffY {
			fmt.Print(string(cp.value))
			return loopNeighbors(start, tiles, (*tiles)[key(cp.x+n.x, cp.y+n.y)], cp, length+1)
		}
	}
	fmt.Println(" = Did not find correct neighbor...", *cp, string(cp.value))
	return length
}

func getStartingShape(tiles *TileMap, x, y int) rune {
	n, e, s, w := false, false, false, false
	if x > 0 {
		for _, neighbor := range NeighborMap[(*tiles)[key(x-1, y)].value] {
			if neighbor.x == 1 {
				w = true
			}
		}
	}
	if y > 0 {
		for _, neighbor := range NeighborMap[(*tiles)[key(x, y-1)].value] {
			if neighbor.y == 1 {
				n = true
			}
		}
	}

	for _, neighbor := range NeighborMap[(*tiles)[key(x+1, y)].value] {
		if neighbor.x == -1 {
			e = true
		}
	}
	for _, neighbor := range NeighborMap[(*tiles)[key(x, y+1)].value] {
		if neighbor.y == -1 {
			s = true
		}
	}

	if w && n {
		return 'J'
	}
	if w && s {
		return '7'
	}
	if e && n {
		return 'L'
	}
	if e && s {
		return 'F'
	}
	if e && w {
		return '-'
	}
	if n && s {
		return '|'
	}
	return '.'
}

func key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

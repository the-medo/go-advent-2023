package day24

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z float64
}

type Hailstone struct {
	ps, velocity *Point
}

const (
	testAreaStart = 200000000000000.0
	testAreaEnd   = 400000000000000.0
	//testAreaStart = 7
	//testAreaEnd   = 27
)

func Solve(input string) {
	rows := utils.SplitRows(input)
	hailstorms := make([]Hailstone, len(rows))
	for i, row := range rows {
		hailstorms[i] = ParseHailstone(row)
	}

	intersections := 0
	for i := 0; i < len(hailstorms); i++ {
		h1 := hailstorms[i]
		for j := i + 1; j < len(hailstorms); j++ {
			h2 := hailstorms[j]

			inc := IntersectHailstorms(&h1, &h2)
			intersections += inc
		}
	}

	fmt.Println("Part 1: ", intersections)
}

func IntersectHailstorms(h1 *Hailstone, h2 *Hailstone) int {
	t := h2.velocity.y*(h2.ps.x-h1.ps.x) - h2.velocity.x*(h2.ps.y-h1.ps.y)
	t = t / ((h1.velocity.x * h2.velocity.y) - (h2.velocity.x * h1.velocity.y))

	if t < 0 {
		fmt.Println("[T < 0] PAST")
		return 0
	}

	x := h1.ps.x + (h1.velocity.x * t)
	y := h1.ps.y + (h1.velocity.y * t)

	s := (x - h2.ps.x) / h2.velocity.x
	if s < 0 {
		fmt.Println("[S < 0]")
		return 0
	}

	withinRange := 0

	if x >= testAreaStart && x <= testAreaEnd && y >= testAreaStart && y <= testAreaEnd {
		withinRange = 1
		fmt.Print("[ OK ] ")
	} else {
		fmt.Print("[OOR]  ")
		withinRange = 0
	}

	fmt.Printf("[%d] The lines intersect at point (%f, %f)\n", withinRange, x, y)

	return withinRange
}

func (h *Hailstone) isPast(x, y float64) bool {
	if (h.velocity.x < 0 && x >= h.ps.x) ||
		(h.velocity.x > 0 && x <= h.ps.x) ||
		(h.velocity.y < 0 && y >= h.ps.y) ||
		(h.velocity.y > 0 && y <= h.ps.y) {
		return true
	}
	return false
}

func ParseHailstone(s string) Hailstone {
	spl := strings.Split(s, " @ ")
	h := Hailstone{
		ps:       ParsePoint(spl[0]),
		velocity: ParsePoint(spl[1]),
	}
	return h
}

func ParsePoint(s string) *Point {
	split := strings.Split(s, ", ")
	x, _ := strconv.Atoi(strings.TrimLeft(split[0], " "))
	y, _ := strconv.Atoi(strings.TrimLeft(split[1], " "))
	z, _ := strconv.Atoi(strings.TrimLeft(split[2], " "))

	return &Point{x: float64(x), y: float64(y), z: float64(z)}
}

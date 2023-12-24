package day24

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Hailstone struct {
	ps, pe, velocity *Point
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
		//fmt.Println("H ", i, " => ", hailstorms[i].ps, hailstorms[i].pe, hailstorms[i].velocity)
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

	//fmt.Println(math.MaxInt)
	a := (math.MaxInt + 1) - 100
	a += 1
	fmt.Println(a)
	fmt.Println("Part 1: ", intersections)
}

func (h *Hailstone) findEnd() {
	h.pe = &Point{
		x: h.ps.x - h.velocity.x,
		y: h.ps.y - h.velocity.y,
		z: h.ps.z - h.velocity.z,
	}
}

func IntersectHailstorms(h1 *Hailstone, h2 *Hailstone) int {
	//a1 := h1.pe.y - h1.ps.y
	//b1 := h1.ps.x - h1.pe.x
	//a2 := h2.pe.y - h2.ps.y
	//b2 := h2.ps.x - h2.pe.x

	//c1 := a1*h1.ps.x + b1*h1.ps.y
	//c2 := a2*h2.ps.x + b2*h2.ps.y

	//c1 := ((h1.pe.y - h1.ps.y)*h1.ps.x + (h1.ps.x - h1.pe.x)*h1.ps.y)
	//c2 := ((h2.pe.y - h2.ps.y)*h2.ps.x + (h2.ps.x - h2.pe.x)*h2.ps.y)
	//denominator := a1*b2 - a2*b1

	//denominator := ((h1.pe.y-h1.ps.y)*(h2.ps.x-h2.pe.x) - (h2.pe.y-h2.ps.y)*(h1.ps.x-h1.pe.x))
	//fmt.Println(" Denominator: ", denominator)

	if ((h1.pe.y-h1.ps.y)*(h2.ps.x-h2.pe.x) - (h2.pe.y-h2.ps.y)*(h1.ps.x-h1.pe.x)) == 0 {
		fmt.Println("The lines are parallel and will not intersect.")
		if ((h2.ps.x-h1.ps.x)/h1.velocity.x)*h1.velocity.y+h1.ps.y == h2.ps.y {
			fmt.Println("Meet!")
		}
		return 0
	}
	x := float64((h2.ps.x-h2.pe.x)*((h1.pe.y-h1.ps.y)*h1.ps.x+(h1.ps.x-h1.pe.x)*h1.ps.y)-(h1.ps.x-h1.pe.x)*((h2.pe.y-h2.ps.y)*h2.ps.x+(h2.ps.x-h2.pe.x)*h2.ps.y)) / float64(((h1.pe.y-h1.ps.y)*(h2.ps.x-h2.pe.x) - (h2.pe.y-h2.ps.y)*(h1.ps.x-h1.pe.x)))
	y := float64((h1.pe.y-h1.ps.y)*((h2.pe.y-h2.ps.y)*h2.ps.x+(h2.ps.x-h2.pe.x)*h2.ps.y)-(h2.pe.y-h2.ps.y)*((h1.pe.y-h1.ps.y)*h1.ps.x+(h1.ps.x-h1.pe.x)*h1.ps.y)) / float64(((h1.pe.y-h1.ps.y)*(h2.ps.x-h2.pe.x) - (h2.pe.y-h2.ps.y)*(h1.ps.x-h1.pe.x)))

	withinRange := 0

	if h1.isPast(x, y) || h2.isPast(x, y) {
		if h1.isPast(x, y) && h2.isPast(x, y) {
			fmt.Print("[Past] [Past] ")
		} else {
			fmt.Print("[Past] ")
		}
	} else {
		if x >= testAreaStart && x <= testAreaEnd && y >= testAreaStart && y <= testAreaEnd {
			withinRange = 1
			fmt.Print("[ OK ] ")
		} else {
			fmt.Print("[OOR]  ")
		}
	}

	fmt.Printf("[%d] The lines intersect at point (%f, %f)\n", withinRange, x, y)

	return withinRange
}

func (h *Hailstone) isPast(x, y float64) bool {
	if (h.velocity.x < 0 && x >= float64(h.ps.x)) ||
		(h.velocity.x > 0 && x <= float64(h.ps.x)) ||
		(h.velocity.y < 0 && y >= float64(h.ps.y)) ||
		(h.velocity.y > 0 && y <= float64(h.ps.y)) {
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
	h.findEnd()
	return h
}

func ParsePoint(s string) *Point {
	split := strings.Split(s, ", ")
	x, _ := strconv.Atoi(strings.TrimLeft(split[0], " "))
	y, _ := strconv.Atoi(strings.TrimLeft(split[1], " "))
	z, _ := strconv.Atoi(strings.TrimLeft(split[2], " "))

	return &Point{x, y, z}
}

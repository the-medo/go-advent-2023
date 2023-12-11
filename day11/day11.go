package day11

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

type Galaxy struct {
	x, y int
}

func (g *Galaxy) key() string {
	return fmt.Sprintf("%d-%d", g.x, g.y)
}

func (g *Galaxy) doubleKey(g2 *Galaxy) string {
	if g.x < g2.x {
		return fmt.Sprintf("%s-%s", g.key(), g2.key())
	} else if g.x == g2.x {
		if g.y < g2.y || g.y == g2.y {
			return fmt.Sprintf("%s-%s", g.key(), g2.key())
		}
	}
	return fmt.Sprintf("%s-%s", g2.key(), g.key())
}

func (g *Galaxy) isDiff(g2 *Galaxy) bool {
	if g.x != g2.x || g.y != g2.y {
		return true
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Galaxy) dist(g2 *Galaxy) int {
	return Abs(g2.x-g.x) + Abs(g2.y-g.y)
}

func Solve(input string) {
	rows := utils.SplitRows(input)
	colsHaveGalaxy := make([]bool, len(rows[0]))

	increaseY := 0
	startedGalaxies := make([]*Galaxy, 0)
	for y, row := range rows {
		rowHasGalaxy := false
		for x, c := range row {
			if c == '#' {
				rowHasGalaxy = true
				colsHaveGalaxy[x] = true
				startedGalaxies = append(startedGalaxies, &Galaxy{
					x: x,
					y: y + increaseY,
				})
			}
		}
		if !rowHasGalaxy {
			increaseY += 999999
		}
	}

	increaseX := 0
	galaxies := make(map[string]*Galaxy)

	for _, galaxy := range startedGalaxies {
		increaseX = 0
		for x, col := range colsHaveGalaxy {
			if !col && x < galaxy.x {
				increaseX += 999999
			}
		}
		galaxy.x += increaseX
		galaxies[galaxy.key()] = galaxy
		fmt.Println("Galaxy ", galaxy.key())
	}

	totalDist := 0
	pairs := 0
	paths := make(map[string]int)
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			_, exists := paths[g1.doubleKey(g2)]
			if !exists && g1.isDiff(g2) {
				dist := g1.dist(g2)
				paths[g1.doubleKey(g2)] = dist
				totalDist += dist
				pairs++
			}
		}
	}

	fmt.Println("====", pairs)
	fmt.Println("Part1: ", paths["1-6-9-10"], paths["4-0-9-10"], paths["0-2-12-7"], paths["0-11-5-11"])
	fmt.Println("Part1: ", paths)
	fmt.Println("Part1: ", totalDist)
}

func key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

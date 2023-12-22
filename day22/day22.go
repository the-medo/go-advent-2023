package day22

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Direction rune

const (
	dirX Direction = 'x'
	dirY Direction = 'y'
	dirZ Direction = 'z'
)

type Point struct {
	x, y, z int
}

type Brick struct {
	p1, p2      Point
	dir         Direction
	id          int
	supportedBy []*Brick
	supports    []*Brick
}

func Solve(input string) {
	rows := utils.SplitRows(input)

	bricks := make([]*Brick, len(rows))
	for i, row := range rows {
		splitByTilde := strings.Split(row, "~")
		p1, p2 := sortPoints(ParsePoint(splitByTilde[0]), ParsePoint(splitByTilde[1]))
		dir := p1.getDir(p2)
		bricks[i] = &Brick{*p1, *p2, dir, i, make([]*Brick, 0), make([]*Brick, 0)}
	}

	sort.Slice(bricks, func(i, j int) bool {
		return sortBricksByZXY(bricks[i], bricks[j]) > 0
	})

	for i, _ := range bricks {
		lowestPossibleZ := 1
		supportingBricks := make([]*Brick, 0)
		for j := 0; j < i; j++ {
			if bricks[i].cross(bricks[j]) {
				newLowestPossibleZ := bricks[j].p2.z + 1
				if lowestPossibleZ == newLowestPossibleZ {
					supportingBricks = append(supportingBricks, bricks[j])
				} else if lowestPossibleZ < newLowestPossibleZ {
					lowestPossibleZ = bricks[j].p2.z + 1
					supportingBricks = make([]*Brick, 0)
					supportingBricks = append(supportingBricks, bricks[j])
				}
			}
		}

		for _, sb := range supportingBricks {
			sb.supports = append(sb.supports, bricks[i])
			bricks[i].supportedBy = append(bricks[i].supportedBy, sb)
		}

		bricks[i].p2.z = lowestPossibleZ + (bricks[i].p2.z - bricks[i].p1.z)
		bricks[i].p1.z = lowestPossibleZ
	}

	canBeDisintegrated := 0
	for i := len(bricks) - 1; i >= 0; i-- {
		minSupportedBy := math.MaxInt
		for _, supportedBrick := range bricks[i].supports {
			supportCount := len((*supportedBrick).supportedBy)
			if supportCount < minSupportedBy {
				minSupportedBy = supportCount
			}
		}
		if minSupportedBy >= 2 {
			canBeDisintegrated++
		}
	}

	for _, b := range bricks {
		b.print()
	}

	fmt.Println("======================")
	fmt.Println("======================")

	//fmt.Println(bricks[6].cross(bricks[5]))

	fmt.Println("Part 1: ", canBeDisintegrated)
}

func (b *Brick) print() {
	fmt.Println("======================")
	fmt.Print("Brick: ", b.id, b.p1, b.p2)
	fmt.Println()
	fmt.Print("Supported by: ")
	for _, x := range b.supportedBy {
		fmt.Print(" ID: ", x.id, "pts:  ", x.p1, x.p2)
	}
	fmt.Println()

	fmt.Print("Supports: ")
	for _, x := range b.supports {
		fmt.Print(" ID:", x.id, "pts:", x.p1, x.p2)
	}
	fmt.Println()

	fmt.Println()
}
func (b1 *Brick) cross(b2 *Brick) bool {
	x1 := b1.p1.x <= b2.p1.x && b2.p1.x <= b1.p2.x
	x2 := b1.p1.x <= b2.p2.x && b2.p2.x <= b1.p2.x
	x3 := b2.p1.x <= b1.p1.x && b1.p1.x <= b2.p2.x
	x4 := b2.p1.x <= b1.p2.x && b1.p2.x <= b2.p2.x

	y1 := b1.p1.y <= b2.p1.y && b2.p1.y <= b1.p2.y
	y2 := b1.p1.y <= b2.p2.y && b2.p2.y <= b1.p2.y
	y3 := b2.p1.y <= b1.p1.y && b1.p1.y <= b2.p2.y
	y4 := b2.p1.y <= b1.p2.y && b1.p2.y <= b2.p2.y

	return (x1 || x2 || x3 || x4) && (y1 || y2 || y3 || y4)
}

func sortBricksByZXY(b1 *Brick, b2 *Brick) int {
	if b1.p1.z == b2.p1.z {
		if b1.p1.x == b2.p1.x {
			return b2.p1.y - b1.p1.y
		}
		return b2.p1.x - b1.p1.x
	}
	return b2.p1.z - b1.p1.z
}

func sortPoints(p *Point, p2 *Point) (*Point, *Point) {
	if p.z != p2.z {
		if p.z > p2.z {
			return p2, p
		}
	} else if p.x != p2.x {
		if p.x > p2.x {
			return p2, p
		}
	} else if p.y != p2.y {
		if p.y > p2.y {
			return p2, p
		}
	}
	return p, p2
}

func (p *Point) getDir(p2 *Point) Direction {
	if p.x != p2.x {
		return dirX
	} else if p.y != p2.y {
		return dirY
	} else if p.z != p2.z {
		return dirZ
	}
	return dirX
}

func ParsePoint(s string) *Point {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])

	return &Point{x, y, z}
}

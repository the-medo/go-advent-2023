package day22

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Brick struct {
	p1, p2             Point
	id                 int
	supportedBy        []*Brick
	supports           []*Brick
	totalFalls         int
	canBeDisintegrated bool
}

func Solve(input string) {
	rows := utils.SplitRows(input)

	// Parse bricks, sort their points right away
	bricks := make([]*Brick, len(rows))
	for i, row := range rows {
		splitByTilde := strings.Split(row, "~")
		p1, p2 := sortPoints(ParsePoint(splitByTilde[0]), ParsePoint(splitByTilde[1]))
		bricks[i] = &Brick{*p1, *p2, i, make([]*Brick, 0), make([]*Brick, 0), 0, false}
	}

	//sort bricks by Z and  X (Z is important, X is not)
	sort.Slice(bricks, func(i, j int) bool {
		return sortBricksByZXY(bricks[i], bricks[j]) > 0
	})

	//let the bricks fall and find their supporting bricks
	for i := range bricks {
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

	// Part 1 - check number of bricks, that can be removed
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
			bricks[i].canBeDisintegrated = true
		}
	}

	// Part 2 - total falls
	totalFallCount := 0
	for _, b := range bricks {
		totalFallCount += b.figureFallCount()

	}

	fmt.Println("Part 1: ", canBeDisintegrated)
	fmt.Println("Part 2: ", totalFallCount)
}

func (b *Brick) figureFallCount() int {
	if b.canBeDisintegrated {
		b.totalFalls = 0
		return 0
	}

	//prepare array of bricks to check - firstly, the ones that current brick supports
	bricksToCheck := make([]*Brick, 0)
	for _, sup := range b.supports {
		bricksToCheck = append(bricksToCheck, sup)
	}

	//prepare map of deleted bricks for faster lookup - actual brick is removed at the start
	removedBricksMap := make(map[int]*Brick)
	removedBricksMap[b.id] = b

	i := 0
	for i < len(bricksToCheck) { //while we are not at an end of bricks to check..
		rb := bricksToCheck[i]

		//check, if some of the support bricks are available
		hasValidSupport := false
		for _, supBrick := range rb.supportedBy {
			_, exists := removedBricksMap[supBrick.id]
			if !exists {
				hasValidSupport = true
				break
			}
		}

		//if there is no other available support brick, this brick falls
		//		=> add bricks that are supported by it to "bricks to check"
		//		=> add it to removed bricks
		if !hasValidSupport {
			for _, sup := range rb.supports {
				bricksToCheck = append(bricksToCheck, sup)
			}
			removedBricksMap[rb.id] = rb
		}
		i++
	}

	// decrease by one, because the brick itself doesn't count
	b.totalFalls = len(removedBricksMap) - 1
	return b.totalFalls
}

func (b *Brick) print() {
	fmt.Println("======================")
	fmt.Print("Brick: ", b.id, b.p1, b.p2, "FALLS: ", b.totalFalls)
	fmt.Println()
	fmt.Print("Supported by: ")
	for _, x := range b.supportedBy {
		fmt.Print(" ID: ", x.id)
	}
	fmt.Println()

	fmt.Print("Supports: ")
	for _, x := range b.supports {
		fmt.Print(" ID:", x.id)
	}
	fmt.Println()

	fmt.Println()
}

// cross checks if two bricks overlap in the X and Y axes.
func (b1 *Brick) cross(b2 *Brick) bool {
	overlapX := b1.p1.x <= b2.p2.x && b2.p1.x <= b1.p2.x
	overlapY := b1.p1.y <= b2.p2.y && b2.p1.y <= b1.p2.y

	return overlapX && overlapY
}

// Z axis is our main axis, lower should be first
func sortBricksByZXY(b1 *Brick, b2 *Brick) int {
	if b1.p1.z == b2.p1.z {
		if b1.p1.x == b2.p1.x {
			return b2.p1.y - b1.p1.y
		}
		return b2.p1.x - b1.p1.x
	}
	return b2.p1.z - b1.p1.z
}

// sort points in case they are reversed
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

func ParsePoint(s string) *Point {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])

	return &Point{x, y, z}
}

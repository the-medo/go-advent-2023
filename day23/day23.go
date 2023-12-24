package day23

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

type Point struct {
	x, y int
	t    rune
}

type Task struct {
	currentId int
	visited   []int
	steps     int
}

type Intersection struct {
	id        int
	point     *Point
	neighbors NeighborMap
}

type Neighbor struct {
	id             int
	startX, startY int
	distance       int
}

type NeighborMap = map[int]int

func Solve(input string) {
	rows := utils.SplitRows(input)

	m := make([][]*Point, len(rows))
	start, end := &Point{}, &Point{}

	for y, row := range rows {
		m[y] = make([]*Point, len(row))
		for x, c := range row {
			m[y][x] = &Point{x, y, c}
			if y == 0 && c == '.' {
				start = m[y][x]
			} else if y == len(rows)-1 && c == '.' {
				end = m[y][x]
			}
		}
	}
	fmt.Println(start, end)

	for i := 2; i <= 2; i++ {
		intersections := make(map[int]*Intersection)
		intersections[1] = &Intersection{
			id:        1,
			point:     start,
			neighbors: make(NeighborMap),
		}

		findIntersection(m, intersections, 1, start.x, start.y+1, start.x, start.y, start.x, start.y+1, 1, i, false)

		queue := []Task{{1, make([]int, 0), 0}}
		counter := 0
		maxSteps := 0
		for len(queue) > 0 {
			task := queue[0]
			queue = queue[1:]

			taskPoint := intersections[task.currentId].point
			if taskPoint.x == end.x && taskPoint.y == end.y && task.steps > maxSteps {
				maxSteps = task.steps
			}
			newTasks := task.check(intersections)
			queue = append(queue, newTasks...)

			counter++
			if counter%100000 == 0 {
				//fmt.Println("i: ", counter, " queue len: ", len(queue))
			}
		}

		// === PART 1 is broken
		fmt.Println("Part ", i, maxSteps)
	}

}

func findIntersection(m [][]*Point, i map[int]*Intersection, startIntersectionId int, branchStartX, branchStartY, lastX, lastY, x, y, distance, part int, throughSlope bool) {
	availablePoints := make([]*Point, 0)
	//fmt.Println(" =============== ", x, y, " ================= ")
	if checkPoint(x-1, y, &m, x, y, part) && !(x-1 == lastX && y == lastY) {
		availablePoints = append(availablePoints, m[y][x-1])
	}
	if checkPoint(x+1, y, &m, x, y, part) && !(x+1 == lastX && y == lastY) {
		availablePoints = append(availablePoints, m[y][x+1])
	}
	if checkPoint(x, y-1, &m, x, y, part) && !(x == lastX && y-1 == lastY) {
		availablePoints = append(availablePoints, m[y-1][x])
	}
	if checkPoint(x, y+1, &m, x, y, part) && !(x == lastX && y+1 == lastY) {
		availablePoints = append(availablePoints, m[y+1][x])
	}

	if len(availablePoints) == 1 {
		pt := m[y][x]
		if part == 1 && (pt.t == '>' || pt.t == 'v') {
			throughSlope = true
		}
		findIntersection(m, i, startIntersectionId, branchStartX, branchStartY, x, y, availablePoints[0].x, availablePoints[0].y, distance+1, part, throughSlope)
	} else {
		//fmt.Println(" . ", len(availablePoints), " . ")
		foundIntersectionId, _, existed := findIntersectionAtPoint(x, y, i)

		i[startIntersectionId].neighbors[foundIntersectionId] = distance
		if !throughSlope {
			i[foundIntersectionId].neighbors[startIntersectionId] = distance
		}

		if !existed {
			for _, pt := range availablePoints {
				findIntersection(m, i, foundIntersectionId, pt.x, pt.y, x, y, pt.x, pt.y, 1, part, false)
			}
		}

	}
}

func findIntersectionAtPoint(x, y int, i map[int]*Intersection) (int, *Intersection, bool) {
	newIntersectionId := len(i) + 1
	existed := false
	for k, v := range i {
		if v.point.x == x && v.point.y == y {
			newIntersectionId, existed = k, true
			break
		}
	}

	var intersection *Intersection

	if existed {
		intersection = i[newIntersectionId]
	} else {
		intersection = &Intersection{
			id:        newIntersectionId,
			point:     &Point{x: x, y: y},
			neighbors: make(NeighborMap),
		}
		i[newIntersectionId] = intersection
	}

	return newIntersectionId, intersection, existed
}

func checkPoint(x, y int, m *[][]*Point, lastX, lastY, part int) bool {
	if x < 0 || x >= len(*m) || y < 0 || y >= len((*m)[0]) {
		return false
	}
	pt := (*m)[y][x].t
	lastPt := (*m)[lastY][lastX].t
	if pt == '#' {
		return false
	}
	if part == 1 && lastPt == '>' && lastX-x != 1 {
		fmt.Println("FAIL", string(pt), string(lastPt), x, y, lastX, lastY)
		return false
	}
	if part == 1 && lastPt == 'v' && lastY-y != 1 {
		fmt.Println("FAIL", string(pt), string(lastPt), x, y, lastX, lastY)
		return false
	}

	return true
}

func (t *Task) check(i map[int]*Intersection) []Task {
	newTasks := make([]Task, 0)

	visited := make([]int, len(t.visited))
	copy(visited, t.visited)
	visited = append(visited, t.currentId)

	for key, distance := range i[t.currentId].neighbors {
		exists := false
		for _, e := range t.visited {
			if e == key {
				exists = true
				break
			}
		}
		if !exists {
			newTasks = append(newTasks, Task{
				currentId: key,
				visited:   visited,
				steps:     t.steps + distance,
			})
		}
	}

	return newTasks
}

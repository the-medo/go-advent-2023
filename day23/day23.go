package day23

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
)

type Point struct {
	x, y int
	max  int
	t    rune
}

type Task struct {
	x, y    int
	m       *[][]*Point
	visited map[string]int
	steps   int
}

func Solve(input string) {
	rows := utils.SplitRows(input)

	m := make([][]*Point, len(rows))
	start, end := &Point{}, &Point{}

	for y, row := range rows {
		m[y] = make([]*Point, len(row))
		for x, c := range row {
			m[y][x] = &Point{x, y, 0, c}
			if y == 0 && c == '.' {
				start = m[y][x]
			} else if y == len(rows)-1 && c == '.' {
				end = m[y][x]
			}
		}
	}
	fmt.Println(start, end)

	//visitedCache := []string{k(start.x, start.y)}
	visitedCache := make(map[string]int)

	queue := []Task{{start.x, start.y, &m, visitedCache, 0}}

	counter := 0
	maxSteps := 0
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]

		if task.x == end.x && task.y == end.y && task.steps > maxSteps {
			maxSteps = task.steps
		}
		//fmt.Println("Processing task: ", counter, " => ", task)
		newTasks := task.check()
		queue = append(queue, newTasks...)

		counter++
		if counter%1000 == 0 {
			fmt.Println("i: ", counter, " queue len: ", len(queue))
		}
	}

	fmt.Println("Part 1: ", maxSteps)

}

func (t *Task) check() []Task {
	newTasks := make([]Task, 0)
	key := k(t.x, t.y)
	t.visited[key] = t.steps

	slopeRight := false
	slopeDown := false

	if (*t.m)[t.y][t.x].t == '>' {
		slopeRight = true
	} else if (*t.m)[t.y][t.x].t == 'v' {
		slopeDown = true
	}

	if checkPoint(t.x-1, t.y, t.m, t.visited) && !slopeDown && !slopeRight {
		newTasks = append(newTasks, Task{t.x - 1, t.y, t.m, copyVisited(t.visited), t.steps + 1})
	}
	if checkPoint(t.x+1, t.y, t.m, t.visited) && !slopeDown {
		newTasks = append(newTasks, Task{t.x + 1, t.y, t.m, copyVisited(t.visited), t.steps + 1})
	}
	if checkPoint(t.x, t.y-1, t.m, t.visited) && !slopeDown && !slopeRight {
		newTasks = append(newTasks, Task{t.x, t.y - 1, t.m, copyVisited(t.visited), t.steps + 1})
	}
	if checkPoint(t.x, t.y+1, t.m, t.visited) && !slopeRight {
		newTasks = append(newTasks, Task{t.x, t.y + 1, t.m, copyVisited(t.visited), t.steps + 1})
	}

	return newTasks
}

func copyVisited(v map[string]int) map[string]int {
	r := make(map[string]int)
	for k, v := range v {
		r[k] = v
	}
	return r
}

func checkPoint(x, y int, m *[][]*Point, visited map[string]int) bool {
	//fmt.Print("Checking point ", x, y, " - ")
	if x < 0 || x >= len(*m) || y < 0 || y >= len((*m)[0]) {
		//fmt.Println("False 1")
		return false
	}
	if (*m)[y][x].t == '#' {
		//fmt.Println("False 2")
		return false
	}
	_, exists := visited[k(x, y)]
	if exists {
		//fmt.Println("False 3")
		return false
	}
	//fmt.Println("True")
	return true
}

func k(n1, n2 int) string {
	return fmt.Sprintf("%d-%d", n1, n2)
}

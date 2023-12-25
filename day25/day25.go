package day25

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"golang.org/x/exp/slices"
	"strings"
)

type Graph = map[string][]string

type Check struct {
	key                    string
	deleted                map[string]bool
	haveToVisit            map[string]bool
	visited                map[string]bool
	visitedFromHaveToVisit map[string]bool
	result                 bool
}

func Solve(input string) {
	rows := utils.SplitRows(input)
	g := make(Graph)
	for _, row := range rows {
		parseRow(g, row)
	}

	checks := make(map[string]*Check)

	for k, _ := range g {
		for _, l := range g[k] {
			for _, m := range g[l] {
				deleted := make(map[string]bool)
				haveToVisit := make(map[string]bool)
				for _, x := range g[k] {
					haveToVisit[x] = true
				}
				//for _, x := range g[l] {
				//	haveToVisit[x] = true
				//}
				//for _, x := range g[m] {
				//	haveToVisit[x] = true
				//}
				deleted[k] = true
				deleted[l] = true
				deleted[m] = true
				delete(haveToVisit, k)
				delete(haveToVisit, l)
				delete(haveToVisit, m)

				if m != k {
					checkKey := key([]string{k, l, m})
					_, exists := checks[checkKey]
					if !exists {
						checks[checkKey] = &Check{
							key:                    checkKey,
							deleted:                deleted,
							haveToVisit:            haveToVisit,
							visited:                make(map[string]bool),
							visitedFromHaveToVisit: make(map[string]bool),
						}
					}
				}
			}
		}
	}

	//fmt.Println(checks)

	//key := "bvb-ntq-xhk"
	//fmt.Println(checks[key])
	//checks[key].process(g)
	//
	//fmt.Println(checks[key])

	for _, c := range checks {
		//if strings.Contains(c.key, "bvb") {
		//	fmt.Println(c.key)
		//}
		c.process(g)
		if !c.result {
			fmt.Println(c)
		}
	}

	fmt.Println(checks["bvb-hfx-jqt"])
}

func (c *Check) process(g Graph) {
	queue := make([]string, 1)
	for k, _ := range c.haveToVisit {
		queue[0] = k
		break
	}
	//fmt.Print(" Processing... ")

	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]
		//fmt.Println(key)

		if c.visited[key] {
			//fmt.Println("Key visited already: ", key)
			continue
		}
		c.visited[key] = true
		if c.haveToVisit[key] {
			c.visitedFromHaveToVisit[key] = true
		}

		if len(c.visitedFromHaveToVisit) == len(c.haveToVisit) {
			c.result = true
			//fmt.Println(" [VFHTV] ", c.visitedFromHaveToVisit)
			//fmt.Println(" ======> ", c.haveToVisit)
			break
		}

		for _, neighborKey := range g[key] {
			if c.deleted[neighborKey] {
				//fmt.Println(" Neighboring key deleted: ", neighborKey)
				continue
			}
			if c.visited[neighborKey] {
				//fmt.Println(" Neighboring key visited: ", neighborKey)
				continue
			}

			//fmt.Println(" Add to queue => ", neighborKey)
			queue = append(queue, neighborKey)

		}
	}
	//fmt.Println("Visited: ", len(c.visited))
}

func parseRow(g Graph, row string) {
	fields := strings.Split(row, ": ")
	main := fields[0]
	other := strings.Split(fields[1], " ")

	_, exists := g[main]
	if !exists {
		g[main] = make([]string, 0)
	}

	for _, c := range other {
		_, exists := g[c]
		if !exists {
			g[c] = make([]string, 0)
		}
		g[c] = append(g[c], main)
		g[main] = append(g[main], c)
	}
}

func key(s []string) string {
	slices.Sort(s)
	return strings.Join(s, "-")
}

package day25

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"golang.org/x/exp/slices"
	"math"
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

	checkCount := 0

	resultCount := 0
	count := 0

	maxVisited := 0
	minVisited := math.MaxInt

	neighborLevel := 0

	for maxVisited == 0 {
		for k, _ := range g {
			for _, l := range g[k] {
				for _, m := range g[l] {
					keyPool := make(map[string]bool)
					keyPool[k] = true
					keyPool[l] = true
					keyPool[m] = true

					for i := 0; i < neighborLevel; i++ {
						for kpk, _ := range keyPool {
							for _, n := range g[kpk] {
								keyPool[n] = true
							}
						}
					}

					nodes := make([]string, 0)
					for k, _ := range keyPool {
						nodes = append(nodes, k)
					}

					for p := 0; p < len(nodes)-2; p++ {
						for q := p + 1; q < len(nodes)-1; q++ {
							for r := q + 1; r < len(nodes); r++ {

								deleted := make(map[string]bool)
								haveToVisit := make(map[string]bool)
								for _, x := range g[nodes[p]] {
									haveToVisit[x] = true
								}
								deleted[nodes[p]] = true
								deleted[nodes[q]] = true
								deleted[nodes[r]] = true
								delete(haveToVisit, nodes[p])
								delete(haveToVisit, nodes[q])
								delete(haveToVisit, nodes[r])

								checkKey := key([]string{nodes[p], nodes[q], nodes[r]})
								_, exists := checks[checkKey]
								if !exists {
									checkCount++
									checks[checkKey] = &Check{
										key:                    checkKey,
										deleted:                deleted,
										haveToVisit:            haveToVisit,
										visited:                make(map[string]bool),
										visitedFromHaveToVisit: make(map[string]bool),
									}

									c := checks[checkKey]
									c.process(g)
									if !c.result {
										resultCount++
										fmt.Print("FALSE: ", c.key, count)
										fmt.Println("  ===> ", len(c.visited))
										if len(c.visited) > maxVisited {
											maxVisited = len(c.visited)
										}
										if len(c.visited) < minVisited {
											minVisited = len(c.visited)
										}
									}

									if checkCount%100000 == 0 {
										fmt.Println("Check count: ", checkCount)
									}
									checks[checkKey] = &Check{}
								}
							}
						}
					}
				}
			}
		}
		neighborLevel++
	}

	fmt.Println("Total check resultCount: ", len(checks))
	diff := len(g) - maxVisited
	fmt.Println("Total found node combinations: ", resultCount, " ; max visited count: ", maxVisited, " ; min visited count: ", minVisited, " diff from max visited: ", diff)
	fmt.Println("Part 1: ", diff*maxVisited)

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

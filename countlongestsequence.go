/*
  Given an unsorted array, count the longest occuring sequence in the array. In the example below, the count is 4 => 1,2,3,4
  This algorithm builds a graph then counts the longest branch, it runs in O(N) time
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(buildGraph([]int{100, 4, 200, 1, 3, 2}))
}

func buildGraph(values []int) int {
	m := make(map[int]*Node)
	for _, v := range values {
		m[v] = &Node{value: v}
	}
	g := Graph{}
	for k := range m {
		if node, ok := m[k+1]; ok {
			m[k].child = node
		}
		if _, ok := m[k-1]; !ok {
			g.parents = append(g.parents, m[k])
		}
	}

	return g.findLongestSequence()
}

type Graph struct {
	parents []*Node
}

type Node struct {
	value int
	child *Node
}

func (g *Graph) findLongestSequence() int {
	max := math.MinInt64
	for _, p := range g.parents {
		counter := 1
		parent := p
		for parent.child != nil {
			counter++
			parent = parent.child
		}
		if counter > max {
			max = counter
		}
	}
	return max
}

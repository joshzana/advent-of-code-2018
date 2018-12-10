package main

import (
	"fmt"
	"helpers"
	"log"
	"sort"
	"strings"
)

type Step struct {
	Name   string
	Before string
}

type StepNodes []*StepNode

type StepNode struct {
	Name    string
	Prereqs StepNodes
	Done    bool
}

func (slice StepNodes) Len() int {
	return len(slice)
}

func (slice StepNodes) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice StepNodes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// https://adventofcode.com/2018/day/7
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := make([]Step, len(lines))
	ParseLines(lines, steps)

	// Make StepNodes map
	stepNodes := make(map[string]*StepNode, len(steps))
	for _, s := range steps {
		stepNodes[s.Name] = &StepNode{Name: s.Name}
		stepNodes[s.Before] = &StepNode{Name: s.Before}
	}

	// build a tree with the Prereqs array pointing to list of nodes that must come before this node
	for _, s := range steps {
		node := stepNodes[s.Before]
		node.Prereqs = append(node.Prereqs, stepNodes[s.Name])
	}

	// Find the roots
	var queue StepNodes
	for _, v := range stepNodes {
		if len(v.Prereqs) == 0 {
			queue = append(queue, v)
			v.Done = true
		}
	}

	// Now BFS
	var traversal StepNodes
	for len(queue) > 0 {
		// sort queue so we get the right first available
		sort.Sort(queue)

		// pop the top
		top := queue[0]
		queue = queue[1:]
		traversal = append(traversal, top)

		// Add all nodes which have no more unmet Prereqs
		for _, v := range stepNodes {

			if !v.Done {
				// Remove top from prereqs since its done
				for i, p := range v.Prereqs {
					if p.Name == top.Name {
						v.Prereqs = append(v.Prereqs[:i], v.Prereqs[i+1:]...)
						break
					}
				}

				// If v has no more prereqs, add it to queue and mark it
				if len(v.Prereqs) == 0 {
					queue = append(queue, v)
					v.Done = true
				}
			}
		}
	}

	// Traverse!
	for _, t := range traversal {
		fmt.Print(t.Name)
	}
}

func ParseLines(lines []string, steps []Step) {
	start := len("Step ")
	for i, l := range lines {
		name := l[start : start+1]
		beforeStart := strings.Index(l, "step ") + len("step ")
		before := l[beforeStart : beforeStart+1]
		steps[i] = Step{name, before}
	}
}

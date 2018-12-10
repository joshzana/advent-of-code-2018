package main

import (
	. "day7"
	"fmt"
	"helpers"
	"log"
	"sort"
)

// https://adventofcode.com/2018/day/7
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := make([]Step, len(lines))
	ParseLines(lines, steps)

	// Make StepNodes map
	stepNodes := MakeStepNodes(steps)

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
				v.RemovePrereq(top.Name)

				// If v has no more prereqs, add it to queue and mark it
				if v.IsUnblocked() {
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

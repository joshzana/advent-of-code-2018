package main

import (
	. "day7"
	"fmt"
	"helpers"
	"log"
	"sort"
)

type Worker struct {
	node      *StepNode
	startTime int
}

type Workers []*Worker

const WorkerCount = 5

// https://adventofcode.com/2018/day/7#part2
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

	var workers Workers
	for i := 0; i < WorkerCount; i++ {
		workers = append(workers, &Worker{})
	}

	for time := 0; ; time++ {
		// sort queue so we get the right first available
		sort.Sort(queue)

		// find all done and add to a worker
		for _, w := range workers {
			if w.node != nil && w.startTime+w.node.Cost() == time {

				// Enqueue all nodes which have no more unmet Prereqs
				for _, v := range stepNodes {

					if !v.Done {
						// Remove w from prereqs since its done
						v.RemovePrereq(w.node.Name)

						// If v has no more prereqs, add it to queue and mark it
						if v.IsUnblocked() {
							queue = append(queue, v)
							v.Done = true
						}
					}
				}

				// mark worker as available
				w.node = nil
				w.startTime = -1
			}
		}

		// Assign as many as available to a free worker with this startTime
		for _, w := range workers {
			if w.node == nil && len(queue) > 0 {
				// pop the top
				top := queue[0]
				queue = queue[1:]

				w.node = top
				w.startTime = time
			}
		}

		var workLeft = false
		for _, w := range workers {
			if w.node != nil {
				workLeft = true
			}
		}

		if !workLeft {
			fmt.Println("Done in time", time)
			break
		}
	}
}

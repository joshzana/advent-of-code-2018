package day7

import (
	"strings"
)

type Step struct {
	Name   string
	Before string
}

func (s *StepNode) IsUnblocked() bool {
	return len(s.Prereqs) == 0
}

func (s *StepNode) RemovePrereq(name string) {
	for i, p := range s.Prereqs {
		if p.Name == name {
			s.Prereqs = append(s.Prereqs[:i], s.Prereqs[i+1:]...)
			break
		}
	}
}

type StepNode struct {
	Name    string
	Prereqs StepNodes
	Done    bool
}

type StepNodes []*StepNode

func (slice StepNodes) Len() int {
	return len(slice)
}

func (slice StepNodes) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice StepNodes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func MakeStepNodes(steps []Step) map[string]*StepNode {
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
	return stepNodes
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

func (s *StepNode) Cost() int {
	return int(rune(s.Name[0])-'A') + 61
}

package day6

import (
	"helpers"
	"strings"
)

type Point struct {
	X int
	Y int
}

func ParseLines(lines []string, points map[int]Point) (int, int) {
	var maxX int
	var maxY int
	for i, l := range lines {
		strs := strings.Split(l, ", ")
		p := Point{helpers.MustAtoi(strs[0]), helpers.MustAtoi(strs[1])}
		points[i] = p
		if p.X > maxX {
			maxX = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxX, maxY
}

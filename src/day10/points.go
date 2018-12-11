package day10

import (
	"helpers"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Velocity struct {
	Dx int
	Dy int
}

func ParseLines(lines []string) ([]*Point, []*Velocity) {
	var points []*Point
	var velocities []*Velocity
	for _, l := range lines {
		f := strings.FieldsFunc(l, func(r rune) bool {
			return strings.ContainsRune("<,>", r)
		})
		points = append(points, &Point{helpers.MustAtoi(f[1]), helpers.MustAtoi(f[2])})
		velocities = append(velocities, &Velocity{helpers.MustAtoi(f[4]), helpers.MustAtoi(f[5])})
	}

	return points, velocities
}

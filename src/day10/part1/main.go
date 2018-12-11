package main

import (
	. "day10"
	"fmt"
	"helpers"
	"log"
	"math"
)

// https://adventofcode.com/2018/day/10
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	points, velocities := ParseLines(lines)

	for i := 0; ; i++ {

		minX := math.MaxInt32
		minY := math.MaxInt32
		maxX := math.MinInt32
		maxY := math.MinInt32
		for _, p := range points {
			if p.X < minX {
				minX = p.X
			}
			if p.Y < minY {
				minY = p.Y
			}
			if p.X > maxX {
				maxX = p.X
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}

		// totally arbitrary filter to only print when the range is small
		if maxY-minY < 50 {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					matched := false
					for _, p := range points {
						if p.X == x && p.Y == y {
							matched = true
							break
						}
					}

					if matched {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
			fmt.Println("Matched at iteration", i)
		}

		// Move the stars one unit
		for i, p := range points {
			p.X += velocities[i].Dx
			p.Y += velocities[i].Dy
		}
	}
}

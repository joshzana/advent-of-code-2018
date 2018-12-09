package main

import (
	"day6"
	"fmt"
	"helpers"
	"log"
)

// https://adventofcode.com/2018/day/6#part2
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Make array of points and calculate the size of the overall space
	points := make(map[int]day6.Point)
	maxX, maxY := day6.ParseLines(lines, points)

	// Make a grid of total distances to all coordinates
	grid := make([][]int, maxY)
	for y := 0; y < maxY; y++ {
		grid[y] = make([]int, maxX)
		for x := 0; x < maxX; x++ {

			// Sum the distances
			var totalDist int
			for _, p := range points {
				totalDist += day6.Dist(p.X, p.Y, x, y)
			}

			grid[y][x] = totalDist
		}
	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}

	// filter and count the values in the grid
	var totalArea int
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[y][x] < 10000 {
				totalArea++
			}
		}
	}

	fmt.Println("totalArea", totalArea)
}

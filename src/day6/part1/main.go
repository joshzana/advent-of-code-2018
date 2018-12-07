package main

import (
	"fmt"
	"helpers"
	"log"
	"math"
	"strings"
)

type point struct {
	x int
	y int
}

// https://adventofcode.com/2018/day/6
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Make array of points and calculate the size of the overall space
	points := make(map[int]point)
	var maxX int
	var maxY int

	for i, l := range lines {
		strs := strings.Split(l, ", ")
		p := point{helpers.MustAtoi(strs[0]), helpers.MustAtoi(strs[1])}
		points[i] = p
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// make grid of nearest point ids, with 0 for tie
	// Note the coords are flipped so the output matches the webpage - grid[y][x]
	grid := make([][]int, maxY)
	for y := 0; y < maxY; y++ {
		grid[y] = make([]int, maxX)
		for x := 0; x < maxX; x++ {

			// track all the distance values - key is distance, value is count at that distance
			distMap := make(map[int]int)

			var minI int
			var minDist = math.MaxInt32
			for i, p := range points {
				d := dist(p.x, p.y, x, y)

				distMap[d]++

				if d < minDist {
					minDist = d
					minI = i
				}
			}
			if distMap[minDist] == 1 {
				grid[y][x] = minI
			} else {
				grid[y][x] = -1
			}
		}
	}

	// find the ones on the edges and blacklist
	blackList := make(map[int]bool)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if y == 0 || y == maxY-1 || x == 0 || x == maxX-1 {
				blackList[grid[y][x]] = true
			}
		}
	}

	// count the values in the grid - key is point id, value is count of locations with that value
	distMap := make(map[int]int)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {

			// filter out the ones in the blacklist
			pid := grid[y][x]
			if !blackList[pid] {
				distMap[pid]++
			}
		}
	}

	// pick the highest in the map
	var maxI int
	var maxDist int
	for k, v := range distMap {
		if v > maxDist {
			maxI = k
			maxDist = v
		}
	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}

	fmt.Println("max I", maxI, "maxDist", maxDist)
}

func dist(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(dist int) int {
	if dist < 0 {
		return -dist
	} else {
		return dist
	}
}

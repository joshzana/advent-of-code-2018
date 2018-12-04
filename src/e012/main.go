package main

import (
	"../helpers"
	"fmt"
	"log"
	"strconv"
)

// https://adventofcode.com/2018/day/1#part2
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	var allSums = make(map[int]bool)
	for line := 0; ; line = (line + 1) % len(lines) {
		num, err := strconv.Atoi(lines[line])
		if err != nil {
			log.Fatal(err)
		}
		sum += num
		if allSums[sum] {
			fmt.Println(sum)
			return
		}
		allSums[sum] = true
	}
}

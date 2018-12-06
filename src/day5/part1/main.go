package main

import (
	"day5"
	"fmt"
	"helpers"
	"log"
)

// https://adventofcode.com/2018/day/5
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := lines[0]

	line = day5.React(line)

	fmt.Println(len(line))
}

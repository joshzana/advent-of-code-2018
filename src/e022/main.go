package main

import (
	"fmt"
	"log"
	"../helpers"
)

// https://adventofcode.com/2018/day/2#part2
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for offset := 0; offset < len(lines[0]); offset++ {
		set := make(map[string]bool)
		for _, line := range lines {
			joined := line[:offset] + line[offset+1:]
			if set[joined] {
				fmt.Println(joined)
			}
			set[joined] = true
		}
	}
}

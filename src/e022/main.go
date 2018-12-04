package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2018/day/2#part2
func main() {
	closer, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	scanner := bufio.NewScanner(closer)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
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

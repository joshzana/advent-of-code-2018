package main

import (
	"fmt"
	"helpers"
	"log"
	"unicode"
)

// https://adventofcode.com/2018/day/5
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := lines[0]
	lastLen := len(line)
	fmt.Println(lastLen)

	for {
		for i := 0; i < len(line)-1; i++ {
			a := rune(line[i])
			b := rune(line[i+1])

			if a != b && unicode.ToLower(a) == unicode.ToLower(b) {
				line = line[:i] + line[i+2:]
			}
		}

		newLen := len(line)
		fmt.Println(newLen)

		if newLen  == lastLen {
			// done, since no more mods were made
			break
		} else {
			lastLen = newLen
		}
	}

}

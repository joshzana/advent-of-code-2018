package main

import (
	"day5"
	"fmt"
	"helpers"
	"log"
	"strings"
)

// https://adventofcode.com/2018/day/5#part2
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var minLine string
	minLen := len(lines[0])

	for i := 0; i < 26; i++ {
		line := lines[0]
		a := string('a' + rune(i))
		b := strings.ToUpper(a)

		line = strings.Replace(line, a, "", -1)
		line = strings.Replace(line, b, "", -1)
		line = day5.React(line)

		if len(line) < minLen {
			minLen = len(line)
			minLine = line
		}

		fmt.Println(a, b, len(line), line)
	}

	fmt.Println("MIN:", minLen, minLine)
}

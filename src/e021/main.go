package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2018/day/2
func main() {
	closer, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	var doubles = 0
	var triples = 0
	scanner := bufio.NewScanner(closer)
	for scanner.Scan() {
		chars := make(map[rune]int)
		for _, c := range scanner.Text() {
			chars[c]++
		}

		for _, v := range chars {
			if v == 2 {
				doubles++
				break
			}
		}

		for _, v := range chars {
			if v == 3 {
				triples++
				break
			}
		}
	}

	fmt.Println(doubles * triples)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2018/day/1#part2
func main() {
	closer, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	var sum = 0
	var allSums = make(map[int]bool)
	for true {
		scanner := bufio.NewScanner(closer)
		for scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				break
			}
			sum += num
			if allSums[sum] {
				fmt.Println(sum)
				return
			}
			allSums[sum] = true
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		closer.Seek(0, 0)
	}
}

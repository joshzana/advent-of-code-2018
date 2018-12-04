package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2018/day/1
func main() {
	closer, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	scanner := bufio.NewScanner(closer)

	var sum = 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

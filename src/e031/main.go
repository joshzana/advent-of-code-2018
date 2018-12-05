package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"../helpers"
)

const Size = 1000

type claim struct {
	x, y          int
	width, height int
}

// https://adventofcode.com/2018/day/3
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fabric := make([][]int, Size, Size)
	for i := range fabric {
		fabric[i] = make([]int, Size)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return strings.ContainsRune("#@x,: ", r)
		})
		c := claim{helpers.MustAtoi(f[1]), helpers.MustAtoi(f[2]), helpers.MustAtoi(f[3]), helpers.MustAtoi(f[4])}

		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				fabric[i][j]++
			}
		}
	}

	var count int
	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

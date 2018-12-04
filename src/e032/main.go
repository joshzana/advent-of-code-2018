package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const Size = 1000

type claim struct {
	id            string
	x, y          int
	width, height int
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// https://adventofcode.com/2018/day/3#part2
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

	var claims []claim

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return strings.ContainsRune("#@x,: ", r)
		})

		c := claim{f[0], mustAtoi(f[1]), mustAtoi(f[2]), mustAtoi(f[3]), mustAtoi(f[4])}
		claims = append(claims, c)
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				fabric[i][j]++
			}
		}
	}

Loop:
	for _, c := range claims {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if fabric[i][j] > 1 {
					continue Loop
				}
			}
		}
		fmt.Println(c.id)
	}
}

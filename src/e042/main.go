package main

import (
	"../helpers"
	"fmt"
	"log"
	"sort"
	"strings"
)

// https://adventofcode.com/2018/day/4#part2
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// sort by date
	sort.Strings(lines)

	var sleepStartTime int

	// build map of guard id to array of minute slots asleep
	var guardCalendar = make(map[int][]int)

	// parse to structs
	var guardId int
	for _, l := range lines {
		f := strings.FieldsFunc(l, func(r rune) bool {
			return strings.ContainsRune("[]", r)
		})

		datetime := strings.Split(f[0], " ")
		time := helpers.MustAtoi(strings.Split(datetime[1], ":")[1])

		if strings.Contains(f[1], "begins") {
			f2 := strings.FieldsFunc(f[1], func(r rune) bool {
				return strings.ContainsRune("# ", r)
			})
			guardId = helpers.MustAtoi(f2[1])
		} else if strings.Contains(f[1], "asleep") {
			sleepStartTime = time
		} else {
			if guardCalendar[guardId] == nil {
				guardCalendar[guardId] = make([]int, 60)
			}
			for i := sleepStartTime; i < time; i++ {
				guardCalendar[guardId][i]++
			}
		}
	}

	// pick highest slot
	var maxGuard int
	var maxMinute int
	var maxCount int
	for k, v := range guardCalendar {
		fmt.Println(k, " ", v)
		for i, e := range v {
			if e > maxCount {
				maxMinute = i
				maxCount = e
				maxGuard = k
			}
		}
	}


	fmt.Println("ID * minute = ", maxGuard * maxMinute)
}

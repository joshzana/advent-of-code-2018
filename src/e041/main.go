package main

import (
	"../helpers"
	"fmt"
	"log"
	"sort"
	"strings"
)

type entry struct {
	id    int
	date  string
	time  int
	event int
}

const (
	BEGIN = 1
	SLEEP = 2
	WAKE  = 3
)

// https://adventofcode.com/2018/day/4
func main() {
	lines, err := helpers.ReadAllLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// sort by date
	sort.Strings(lines)

	var sleepStartTime int

	// build map of guard id to time asleep
	var guardTimes = make(map[int]int)

	// parse to structs
	var guardId int
	var entries []entry
	for _, l := range lines {
		f := strings.FieldsFunc(l, func(r rune) bool {
			return strings.ContainsRune("[]", r)
		})

		datetime := strings.Split(f[0], " ")
		date := datetime[0]
		time := helpers.MustAtoi(strings.Split(datetime[1], ":")[1])

		var event int
		if strings.Contains(f[1], "begins") {
			event = BEGIN
			f2 := strings.FieldsFunc(f[1], func(r rune) bool {
				return strings.ContainsRune("# ", r)
			})
			guardId = helpers.MustAtoi(f2[1])
		} else if strings.Contains(f[1], "asleep") {
			event = SLEEP
			sleepStartTime = time
		} else {
			event = WAKE
			sleepLength := time - sleepStartTime
			guardTimes[guardId] += sleepLength
		}

		entries = append(entries, entry{guardId, date, time, event})
	}

	// pick sleepiest guard
	var maxGuard int
	var maxTime int
	for k, v := range guardTimes {
		if v > maxTime {
			maxGuard = k
			maxTime = v
		}
	}

	// table of minute frequencies asleep
	sleepTable := make([]int, 60)
	for _, e := range entries {
		if e.id == maxGuard {
			if e.event == SLEEP {
				sleepStartTime = e.time
			} else if e.event == WAKE {
				for i := sleepStartTime; i < e.time; i++ {
					sleepTable[i]++
				}
			}
		}
	}

	// max minute
	var max int
	var minute int
	for i, s := range sleepTable {
		if s > max {
			max = s
			minute = i
		}
	}

	fmt.Println("ID * minute = ", minute*maxGuard)
}

package helpers

import (
	"log"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return i
}

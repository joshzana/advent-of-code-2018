package main

import "fmt"

func main() {

	var recipes = []int{3, 7}
	const input = 74501
	var a = 0 // index of first elf
	var b = 1 // index of second elf

	for len(recipes) < 30000000 {
		// combine recipes
		var sum = recipes[a] + recipes[b]
		var r1 = sum / 10 // high digit
		var r2 = sum % 10 // low digit

		// only append high digit if there is one
		if sum >= 10 {
			recipes = append(recipes, r1)
		}

		// always append low digit
		recipes = append(recipes, r2)

		// pick next index for each elf
		// move 1 plus the score of current recipe slots
		var aScore = 1 + recipes[a]
		var bScore = 1 + recipes[b]
		for i := 0; i < aScore; i++ {
			a = (a + 1) % len(recipes)
		}

		for i := 0; i < bScore; i++ {
			b = (b + 1) % len(recipes)
		}

		//printRecipes(recipes, a, b)
	}

	// Part 1: Print the ten recipes after the number of recipes in the input
	fmt.Print("Part 1: ")
	for i := input; i < input+10; i++ {
		fmt.Print(recipes[i])
	}
	fmt.Println()

	// Part 2:
	// Find the array inside the larger array
	var inputArray = []int{0, 7, 4, 5, 0, 1}
	for i := 0; i < len(recipes); i++ {
		var offset = 0
		for j := 0; j < len(inputArray); j++ {
			if recipes[i+offset] == inputArray[j] {
				offset++
			} else {
				break
			}
		}

		if offset == len(inputArray) {
			fmt.Println("Part 2: ", i)
			break
		}
	}
}

func printRecipes(recipes []int, a int, b int) {
	for i, v := range recipes {
		if i == a {
			fmt.Print("(", v, ")")
		} else if i == b {
			fmt.Print("[", v, "]")
		} else {
			fmt.Print(" ", v, " ")
		}
	}
}

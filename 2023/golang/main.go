package main

import "fmt"

func main() {

	answer, err := Day1Part1()
	if err != nil {
		fmt.Printf("Day 1 Part 1 Error: %v\n", err)
	} else {
		fmt.Printf("Day 1 Part 1 Answer: %v\n", answer)
	}

	answer, err = Day1Part2()
	if err != nil {
		fmt.Printf("Day 1 Part 2 Error: %v\n", err)
	} else {
		fmt.Printf("Day 1 Part 2 Answer: %v\n", answer)
	}

	answer, err = Day2Part1()
	if err != nil {
		fmt.Printf("Day 2 Part 1 Error: %v\n", err)
	} else {
		fmt.Printf("Day 2 Part 1 Answer: %v\n", answer)
	}

	answer, err = Day2Part2()
	if err != nil {
		fmt.Printf("Day 2 Part 2 Error: %v\n", err)
	} else {
		fmt.Printf("Day 2 Part 2 Answer: %v\n", answer)
	}

	answer, err = Day3Part1()
	if err != nil {
		fmt.Printf("Day 3 Part 1 Error: %v\n", err)
	} else {
		fmt.Printf("Day 3 Part 1 Answer: %v\n", answer)
	}

	answer, err = Day3Part2()
	if err != nil {
		fmt.Printf("Day 3 Part 2 Error: %v\n", err)
	} else {
		fmt.Printf("Day 3 Part 2 Answer: %v\n", answer)
	}
}

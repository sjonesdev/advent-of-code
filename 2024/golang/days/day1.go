package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const exampleInputPath = "../inputs/day1example.txt"
const realInputPath = "../inputs/day1.txt"

func getInput(useExample bool) ([]int, []int) {
	inputPath := realInputPath
	if useExample {
		inputPath = exampleInputPath
	}
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 0, 255)
	right := make([]int, 0, 255)
	for scanner.Scan() {
		text := scanner.Text()
		nums := strings.SplitN(text, "   ", 2)

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Error parsing int %s: %v", nums[0], err)
		}
		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Error parsing int %s: %v", nums[1], err)
		}
		right = append(right, rightNum)
	}

	return left, right
}

func Day1Pt1(useExample bool) int {
	left, right := getInput(useExample)

	slices.Sort(left)
	slices.Sort(right)

	diffSum := 0
	for idx, leftVal := range left {
		rightVal := right[idx]
		if rightVal > leftVal {
			diffSum += rightVal - leftVal
		} else {
			diffSum += leftVal - rightVal
		}
	}

	return diffSum
}

func Day1Pt2(useExample bool) int {
	left, right := getInput(useExample)

	rightOccurences := make(map[int]int)
	for _, val := range right {
		rightOccurences[val] += 1
	}
	similarity := 0
	for _, val := range left {
		similarity += val * rightOccurences[val]
	}
	return similarity
}

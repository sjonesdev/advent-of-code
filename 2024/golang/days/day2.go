package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const day2ExampleInputPath = "../inputs/day2example.txt"
const day2RealInputPath = "../inputs/day2.txt"

func getDay2Input(useExample bool) [][]int {
	inputPath := day2RealInputPath
	if useExample {
		inputPath = day2ExampleInputPath
	}
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]int, 0, 1023)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		text := scanner.Text()
		levelStrs := strings.Split(text, " ")
		levels := make([]int, len(levelStrs))
		for i, levelStr := range levelStrs {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				log.Fatalf("Error parsing int \"%s\" in position %d on line %d: %v\n", levelStr, i, lineNum, err)
			}
			levels[i] = level
		}
		input = append(input, levels)
	}
	return input
}

type Safety int

const (
	Unsafe Safety = iota
	Safe
)

func checkSafeSkipIdx(report []int, skipIdx int) Safety {
	direction := Unknown
	first := true
	last := 0
	for idx, level := range report {
		if idx == skipIdx {
			continue
		}

		diff := level - last
		switch direction {
		case Unknown:
			if !first {
				if diff > 0 && diff <= 3 {
					direction = Increasing
				} else if diff < 0 && diff >= -3 {
					direction = Decreasing
				} else {
					return Unsafe
				}
			} else {
				first = false
			}
		case Increasing:
			if diff <= 0 || diff > 3 {
				return Unsafe
			}
		case Decreasing:
			if diff >= 0 || diff < -3 {
				return Unsafe
			}
		}
		last = level
	}
	return Safe
}

type Direction int

const (
	Unknown Direction = iota
	Increasing
	Decreasing
)

func Day2Pt1(useExample bool) int {
	numSafe := 0
	for _, report := range getDay2Input(useExample) {
		numSafe += int(checkSafeSkipIdx(report, -1))
	}
	return numSafe
}

func Day2Pt2(useExample bool) int {
	numSafe := 0
reports:
	for _, report := range getDay2Input(useExample) {
		if checkSafeSkipIdx(report, -1) == Safe {
			numSafe++
			continue reports
		}
		for idx := range report {
			safety := checkSafeSkipIdx(report, idx)
			if safety == Safe {
				numSafe++
				continue reports
			}
		}
	}
	return numSafe
}

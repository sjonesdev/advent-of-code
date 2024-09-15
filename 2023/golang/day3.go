package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// attempt 1: 529673

var symbols = []rune{'*', '-', '=', '&', '#', '$', '/', '+', '@', '%'}

func findSymbols() {
	file, err := os.Open("../inputday3.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	var symbols []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
	line:
		for _, char := range text {
			if char == '.' {
				continue
			}
			for _, symbol := range symbols {
				if char == symbol {
					continue line
				}
			}
			symbols = append(symbols, char)
		}
	}
	fmt.Printf("Symbols: %v\n", string(symbols))
}

var count = 0

func checkLine(prev string, cur string, next string, linelen int) int {
	count++
	numrunes := make([]rune, 0, 3)
	innum := false
	shouldadd := false
	linesum := 0
	for i, char := range cur {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numrunes = append(numrunes, char)
			if !shouldadd {
				if i > 0 {
					if slices.Contains(symbols, []rune(cur)[i-1]) ||
						slices.Contains(symbols, []rune(prev)[i-1]) ||
						slices.Contains(symbols, []rune(next)[i-1]) {
						shouldadd = true
					}
				}
				if slices.Contains(symbols, []rune(prev)[i]) ||
					slices.Contains(symbols, []rune(next)[i]) {
					shouldadd = true
				}
				if i < linelen-1 {
					if slices.Contains(symbols, []rune(cur)[i+1]) ||
						slices.Contains(symbols, []rune(prev)[i+1]) ||
						slices.Contains(symbols, []rune(next)[i+1]) {
						shouldadd = true
					}
				}
			}
			innum = true
			break
		case '.':
			if innum {
				num, err := strconv.Atoi(string(numrunes))
				if err != nil {
					fmt.Printf("Error converting %v to int: %v\n", string(numrunes), err)
				} else if shouldadd {
					// if there's a neighboring symbol
					linesum += num
					if count < 100 {
						fmt.Printf("Line %v, adding %v\n", count, num)
					}
				}
				numrunes = numrunes[:0] // empties slice but leaves capacity unchanged
			}
			innum = false
			shouldadd = false
			break
		default:
			if innum {
				num, err := strconv.Atoi(string(numrunes))
				if err != nil {
					fmt.Printf("Error converting %v to int: %v\n", string(numrunes), err)
				} else {
					// this is necessarily a symbol since it's not a number or period
					linesum += num
					if count < 100 {
						fmt.Printf("Line %v, adding %v\n", count, num)
					}
				}
				numrunes = numrunes[:0] // empties slice but leaves capacity unchanged
			}
			innum = false
			shouldadd = false
			break
		}
	}
	return linesum
}

// assumptions: all lines of the input file are the same length, there are no numbers longer than 3 digits
func Day3Part1() (int, error) {
	file, err := os.Open("../inputday3.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()

	// naive solution: read entire file into a 2d array of characters and for every number check all around it for symbols
	// potentially better: read file line by line and have 3 buffers for the previous, current, and next lines. Any time a number is encountered in a line, check for neighboring symbols and if there is one, construct the rest of the number without checking for symbols and add it to the total, if not, add the number to the number construction string and continue to the next character

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("Empty file")
	}
	cur := scanner.Text()
	linelen := len(cur)
	prev := strings.Repeat(".", len(cur))
	next := ""

	sum := 0
	// check first
	for scanner.Scan() {
		next = scanner.Text()

		sum += checkLine(prev, cur, next, linelen)

		prev = cur
		cur = next
	}
	// check last
	next = strings.Repeat(".", linelen)
	sum += checkLine(prev, cur, next, linelen)

	return sum, nil
}

func Day3Part2() (int, error) {
	file, err := os.Open("../inputday3.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()
	return 0, nil
}

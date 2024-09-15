package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func Day1Part1() (int, error) {
	file, err := os.Open("../inputday1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()

	twodigits, err := regexp.Compile(`\d.*\d`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	onedigit, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		idxs := twodigits.FindStringSubmatchIndex(text)

		if len(idxs) < 2 {
			num := onedigit.FindStringSubmatch(text)[0]
			num += num

			numval, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", num, err)
			} else {
				sum += numval
			}
		} else {
			num := text[idxs[0]:idxs[0]+1] + text[idxs[1]-1:idxs[1]]

			numval, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", num, err)
			} else {
				sum += numval
			}
		}
	}

	return sum, nil
}

func Day1Part2() (int, error) {
	file, err := os.Open("../inputday1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()

	twodigits, err := regexp.Compile(`(?:\d|one|two|three|four|five|six|seven|eight|nine).*(?:\d|one|two|three|four|five|six|seven|eight|nine)`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
	}

	onedigit, err := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
	}

	sum := 0

	digitchars := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		idxs := twodigits.FindStringSubmatchIndex(text)

		if len(idxs) < 2 {
			num := onedigit.FindStringSubmatch(text)[0]
			if !slices.Contains(digitchars, num[0]) {
				if num[0] == 'o' {
					num = "1"
				} else if num[0] == 't' {
					if num[1] == 'w' {
						num = "2"
					} else {
						num = "3"
					}
				} else if num[0] == 'f' {
					if num[1] == 'o' {
						num = "4"
					} else {
						num = "5"
					}
				} else if num[0] == 's' {
					if num[1] == 'i' {
						num = "6"
					} else {
						num = "7"
					}
				} else if num[0] == 'e' {
					if num[1] == 'i' {
						num = "8"
					}
				} else {
					num = "9"
				}
			}

			num += num

			numval, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", num, err)
			} else {
				sum += numval
			}
		} else {
			num := text[idxs[0]:idxs[1]]
			num1 := num[0:1]
			if !slices.Contains(digitchars, num[0]) {
				if num[0] == 'o' {
					num1 = "1"
				} else if num[0] == 't' {
					if num[1] == 'w' {
						num1 = "2"
					} else {
						num1 = "3"
					}
				} else if num[0] == 'f' {
					if num[1] == 'o' {
						num1 = "4"
					} else {
						num1 = "5"
					}
				} else if num[0] == 's' {
					if num[1] == 'i' {
						num1 = "6"
					} else {
						num1 = "7"
					}
				} else if num[0] == 'e' {
					if num[1] == 'i' {
						num1 = "8"
					}
				} else {
					num1 = "9"
				}
			}

			// "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"
			last := len(num) - 1
			num2 := num[last:]
			if !slices.Contains(digitchars, num[last]) {
				if num[last] == 'e' {
					if num[last-2] == 'o' {
						num2 = "1"
					} else if num[last-2] == 'r' {
						num2 = "3"
					} else if num[last-3] == 'f' {
						num2 = "5"
					} else {
						num2 = "9"
					}
				} else if num[last] == 'o' {
					num2 = "2"
				} else if num[last] == 'r' {
					num2 = "4"
				} else if num[last] == 'x' {
					num2 = "6"
				} else if num[last] == 'n' {
					num2 = "7"
				} else {
					num2 = "8"
				}
			}

			numval, err := strconv.Atoi(num1 + num2)
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", num, err)
			} else {
				sum += numval
			}
		}
	}

	return sum, nil
}

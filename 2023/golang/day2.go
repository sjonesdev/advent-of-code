package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day2Part1() (int, error) {
	redcubes := 12
	greencubes := 13
	bluecubes := 14

	file, err := os.Open("../inputday2.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()

	sum := 0
	game := 0

	numre, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	gamere, err := regexp.Compile(`Game \d+:`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game++
		text := scanner.Text()
		gamematch := gamere.FindStringSubmatchIndex(text)
		if len(gamematch) < 2 {
			continue
		}
		text = text[gamematch[1]:]
		subsets := strings.Split(text, ";")
		valid := true
	checkgame:
		for _, subset := range subsets {
			colors := strings.Split(subset, ",")
			for _, color := range colors {
				if strings.Contains(color, "red") {
					numredstr := numre.FindStringSubmatch(color)[0]
					numred, err := strconv.Atoi(numredstr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numredstr, err)
					} else if numred > redcubes {
						valid = false
						break checkgame
					}
				} else if strings.Contains(color, "green") {
					numgreenstr := numre.FindStringSubmatch(color)[0]
					numgreen, err := strconv.Atoi(numgreenstr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numgreenstr, err)
					} else if numgreen > greencubes {
						valid = false
						break checkgame
					}
				} else if strings.Contains(color, "blue") {
					numbluestr := numre.FindStringSubmatch(color)[0]
					numblue, err := strconv.Atoi(numbluestr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numbluestr, err)
					} else if numblue > bluecubes {
						valid = false
						break checkgame
					}
				}
			}
		}
		if valid {
			sum += game
		}
	}

	return sum, nil
}

func Day2Part2() (int, error) {
	file, err := os.Open("../inputday2.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return 0, err
	}
	defer file.Close()

	sum := 0
	game := 0

	numre, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	gamere, err := regexp.Compile(`Game \d+:`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game++
		text := scanner.Text()
		gamematch := gamere.FindStringSubmatchIndex(text)
		if len(gamematch) < 2 {
			continue
		}
		text = text[gamematch[1]:]
		subsets := strings.Split(text, ";")
		minred := 0
		mingreen := 0
		minblue := 0
		for _, subset := range subsets {
			colors := strings.Split(subset, ",")
			for _, color := range colors {
				if strings.Contains(color, "red") {
					numredstr := numre.FindStringSubmatch(color)[0]
					numred, err := strconv.Atoi(numredstr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numredstr, err)
					} else if numred > minred {
						minred = numred
					}
				} else if strings.Contains(color, "green") {
					numgreenstr := numre.FindStringSubmatch(color)[0]
					numgreen, err := strconv.Atoi(numgreenstr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numgreenstr, err)
					} else if numgreen > mingreen {
						mingreen = numgreen
					}
				} else if strings.Contains(color, "blue") {
					numbluestr := numre.FindStringSubmatch(color)[0]
					numblue, err := strconv.Atoi(numbluestr)
					if err != nil {
						fmt.Printf("Error converting %v to int: %v\n", numbluestr, err)
					} else if numblue > minblue {
						minblue = numblue
					}
				}
			}
		}
		sum += minred * mingreen * minblue
	}

	return sum, nil
}

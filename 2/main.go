package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	string
	int
}

const (
	Forward = "forward"
	Down = "down"
	Up = "up"
)

type instructions []pair

func main()  {
	test := "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"
	first := "forward 1\nforward 2\ndown 5\ndown 5\ndown 4\ndown 9\nup 6\nup 7\ndown 2\nforward 9\nup 4\nforward 7\nforward 9\nup 3\ndown 1\nforward 5\ndown 3\nforward 3\nforward 3\nforward 3\ndown 2\ndown 5\nforward 7\ndown 7\nup 7\ndown 9\ndown 1\ndown 4\ndown 9\ndown 2\nforward 2\nforward 4\nup 7\nup 1\nforward 3\nforward 8\nforward 9\nforward 6\nforward 9\nforward 1\nforward 5\ndown 9\nup 7\ndown 9\nforward 2\nforward 9\nforward 1\nforward 5\nup 8\ndown 5\nforward 4\nup 6\nup 9\nforward 2\nup 8\ndown 1\nup 5\nforward 3\ndown 1\nforward 6\nup 6\nforward 9\nforward 1\nforward 3\ndown 4\ndown 9\ndown 8\nup 9\ndown 9\ndown 2\ndown 4\nforward 2\ndown 4\ndown 2\ndown 8\nup 3\nup 9\nforward 3\ndown 5\ndown 1\nup 6\nup 6\ndown 4\nup 3\nforward 1\ndown 2\ndown 7\nforward 1\ndown 4\nforward 5\ndown 5\nforward 3\nforward 8\ndown 4\nforward 3\nforward 2\ndown 4\nforward 6\nforward 6\ndown 9\ndown 3\nup 7\nup 6\ndown 8\ndown 4\ndown 4\ndown 8\ndown 4\nforward 5\nup 7\ndown 8\ndown 4\ndown 5\ndown 3\nforward 1\nup 1\nforward 9\nforward 4\ndown 9\nforward 5\nup 4\ndown 3\nup 7\nup 2\ndown 5\ndown 2\nforward 8\nup 1\ndown 8\nforward 2\nforward 8\ndown 9\nforward 3\nforward 7\nforward 1\ndown 2\ndown 8\nforward 1\nforward 9\nforward 9\nup 3\nup 7\nforward 9\ndown 4\nup 6\nforward 2\ndown 2\ndown 4\ndown 4\ndown 6\ndown 5\nforward 2\ndown 8\ndown 1\nup 6\nup 1\nup 4\ndown 5\nforward 5\nforward 4\nforward 1\nforward 9\nup 9\ndown 9\ndown 5\ndown 7\nup 6\nup 2\nforward 5\ndown 5\ndown 3\ndown 8\ndown 6\nforward 4\ndown 9\nup 3\nforward 2\nforward 9\nforward 6\nforward 5\ndown 5\ndown 1\ndown 2\nforward 9\ndown 2\ndown 2\ndown 3\nforward 3\nforward 9\nforward 1\ndown 3\ndown 8\nforward 7\ndown 9\nforward 4\nup 3\nup 7\nup 4\ndown 5\nforward 9\nforward 2\nforward 2\ndown 3\nup 5\ndown 5\ndown 4\nforward 2\nforward 7\nup 2\ndown 8\nup 2\nup 2\nforward 2\ndown 2\nforward 3\ndown 3\nup 8\nforward 7\nup 5\nforward 4\ndown 6\ndown 8\nforward 4\nforward 3\nup 9\nup 2\ndown 2\ndown 5\ndown 8\ndown 1\nforward 9\ndown 6\nforward 1\ndown 9\ndown 4\nup 6\nup 3\nforward 1\ndown 8\nup 3\ndown 6\ndown 7\ndown 2\nforward 5\ndown 6\nforward 8\nup 7\ndown 5\ndown 3\nforward 4\nforward 5\nforward 3\nforward 4\nforward 6\nforward 2\nforward 1\ndown 3\ndown 5\ndown 3\ndown 5\nforward 4\ndown 7\nforward 8\nforward 5\nup 7\nup 3\nforward 9\nup 1\nforward 9\nup 8\ndown 3\nforward 1\nforward 6\nforward 9\ndown 1\nup 9\nforward 5\ndown 6\nforward 8\ndown 7\ndown 3\nup 4\ndown 6\nforward 5\nforward 6\nforward 5\nup 2\ndown 5\nup 7\nup 4\nup 5\nforward 3\ndown 9\nup 4\nforward 9\nforward 8\nforward 6\ndown 5\ndown 4\ndown 2\nup 5\nup 7\nup 2\nforward 9\nforward 9\ndown 9\ndown 4\nup 2\nforward 3\nup 3\nup 2\ndown 9\nforward 8\nforward 6\ndown 6\nforward 3\ndown 1\nforward 4\nforward 9\nforward 5\ndown 2\ndown 7\nup 6\ndown 3\nforward 7\ndown 3\nup 3\nup 8\nforward 3\nup 7\nforward 5\ndown 7\nforward 7\nforward 3\ndown 6\ndown 3\nforward 5\nforward 9\nup 8\ndown 7\ndown 3\ndown 7\ndown 4\ndown 1\ndown 7\nup 6\nforward 8\nup 7\ndown 9\nforward 6\ndown 4\nforward 6\nup 9\nforward 4\ndown 5\nup 3\nforward 5\nforward 6\ndown 8\nup 9\nforward 4\nup 5\nforward 4\nforward 2\nforward 8\ndown 7\nforward 1\ndown 8\nforward 8\nforward 4\ndown 4\nforward 5\ndown 2\ndown 5\nforward 9\ndown 7\nforward 1\ndown 1\nforward 9\nforward 3\nforward 9\nforward 8\ndown 5\ndown 6\nforward 8\nup 9\nforward 7\ndown 1\nforward 9\nup 7\nforward 2\nforward 6\nforward 1\ndown 8\ndown 6\ndown 7\ndown 6\nup 5\ndown 5\nforward 9\ndown 6\ndown 9\nforward 9\ndown 7\nup 7\nforward 1\ndown 5\ndown 8\nup 5\ndown 6\nup 5\nup 7\nforward 3\nforward 2\ndown 5\ndown 6\nforward 3\ndown 4\ndown 5\nup 4\nforward 5\ndown 4\ndown 5\nforward 4\ndown 1\nforward 1\ndown 1\nforward 4\nforward 2\ndown 3\nforward 1\ndown 1\nforward 2\nforward 2\nforward 6\nup 5\nforward 5\ndown 9\ndown 1\nforward 7\nup 9\ndown 2\ndown 1\nforward 3\nup 5\ndown 8\nforward 2\nforward 1\ndown 7\nforward 5\nup 6\nforward 2\nup 5\nforward 8\nup 6\ndown 6\ndown 3\nforward 1\nforward 7\ndown 7\ndown 1\nup 7\nforward 2\nup 5\nforward 4\nforward 9\nforward 4\nforward 2\nforward 4\ndown 3\ndown 7\nforward 1\nup 9\nup 2\nforward 1\ndown 5\nup 9\nforward 6\ndown 7\ndown 2\ndown 7\ndown 2\ndown 1\ndown 7\ndown 6\nup 1\nup 4\ndown 9\nup 3\nforward 1\ndown 2\nforward 4\nup 4\nup 9\ndown 4\nforward 6\ndown 1\ndown 1\ndown 8\nup 5\nforward 1\nup 6\ndown 5\nforward 4\nup 8\ndown 4\nforward 4\nforward 3\ndown 7\ndown 1\nforward 3\nforward 1\nup 6\ndown 1\ndown 8\nforward 6\ndown 4\ndown 6\nforward 5\nforward 3\nforward 5\ndown 4\nforward 7\ndown 6\ndown 6\ndown 9\nup 9\nforward 5\nup 9\nup 4\nup 6\ndown 4\nforward 3\nup 2\ndown 7\ndown 8\nforward 7\ndown 4\ndown 3\ndown 5\ndown 1\nforward 5\nup 4\ndown 3\ndown 3\ndown 6\nforward 9\ndown 1\nforward 4\ndown 9\nforward 1\nforward 4\ndown 1\nup 5\ndown 6\nforward 5\nup 5\nforward 5\ndown 8\ndown 1\ndown 8\nup 1\ndown 1\nforward 8\nforward 3\nup 2\nforward 9\nforward 1\nforward 3\ndown 2\ndown 7\ndown 2\nup 4\nup 3\ndown 2\nforward 2\nforward 9\nforward 8\ndown 8\nforward 3\nup 9\nup 6\ndown 9\ndown 1\nup 3\nup 2\nforward 2\ndown 6\nup 2\nup 1\ndown 9\ndown 3\ndown 6\nup 7\nup 5\nforward 8\ndown 1\nforward 7\ndown 6\ndown 1\nup 9\nforward 9\nforward 8\ndown 3\ndown 9\ndown 5\nforward 7\nup 1\nup 4\nup 4\ndown 7\ndown 1\nup 2\ndown 2\nforward 8\nforward 7\nup 8\ndown 1\ndown 8\nup 7\nforward 1\ndown 9\nforward 7\nforward 1\ndown 4\ndown 8\ndown 1\nforward 5\nforward 8\nforward 5\ndown 8\ndown 7\nup 5\nforward 8\ndown 5\nup 9\ndown 5\ndown 9\nforward 2\nforward 6\nforward 2\nup 1\nforward 4\nforward 9\nforward 7\ndown 7\ndown 3\nforward 9\nforward 6\nup 5\nforward 5\nforward 7\ndown 9\nforward 6\ndown 7\nforward 5\nforward 5\nforward 4\nforward 1\nforward 1\nup 7\nforward 3\nup 3\nforward 6\nup 3\ndown 9\nforward 9\nup 6\nup 3\nforward 2\ndown 2\nforward 9\ndown 7\nup 7\nforward 6\nforward 2\ndown 2\ndown 4\nforward 1\nforward 4\ndown 4\nup 9\ndown 4\ndown 4\ndown 3\nforward 6\nforward 3\ndown 3\nforward 5\nforward 7\nup 4\nforward 1\nforward 5\nforward 2\nforward 5\nforward 5\nforward 2\nup 8\ndown 7\nup 7\ndown 7\nforward 1\nforward 5\nforward 3\nforward 7\nforward 5\nforward 6\nup 7\nforward 3\ndown 7\ndown 2\nup 9\nforward 6\ndown 7\nforward 9\nup 8\ndown 1\nup 8\nup 2\ndown 7\ndown 6\ndown 5\nup 7\nforward 6\ndown 5\nforward 7\ndown 6\ndown 2\nup 3\nup 7\nup 5\nforward 9\nforward 2\ndown 1\ndown 5\nup 9\nforward 8\ndown 7\nforward 1\nup 6\ndown 1\ndown 3\nforward 3\nforward 6\ndown 4\nforward 8\nup 2\ndown 8\nup 4\nup 9\nforward 8\ndown 9\nforward 3\nforward 7\ndown 5\nforward 4\nup 3\nup 1\nforward 7\ndown 6\nup 4\ndown 3\nforward 8\ndown 9\nforward 2\ndown 8\nforward 9\nup 7\nforward 2\nup 1\ndown 7\ndown 1\nforward 4\nforward 5\ndown 4\ndown 9\nforward 2\ndown 9\ndown 5\nup 2\ndown 6\nforward 8\nup 6\ndown 3\ndown 5\ndown 1\nup 7\ndown 5\nforward 8\nup 4\ndown 7\ndown 4\ndown 4\ndown 2\ndown 3\ndown 1\ndown 7\nforward 4\ndown 4\ndown 7\ndown 7\ndown 9\nup 1\nup 7\nforward 4\nup 7\nforward 6\nforward 8\nforward 2\nforward 8\nup 8\ndown 3\nforward 7\ndown 9\nforward 9\nforward 6\nup 3\ndown 4\ndown 3\nforward 6\nforward 2\nforward 3\ndown 8\nforward 1\nforward 5\nup 3\ndown 8\nforward 7\nforward 4\ndown 3\nforward 2\ndown 9\ndown 9\nforward 9\ndown 7\nforward 6\nforward 4\ndown 5\nforward 9\nforward 3\ndown 1\ndown 1\ndown 7\nforward 8\ndown 3\nforward 7\nforward 8\nup 3\nforward 8\nforward 8\nup 6\nforward 2\ndown 3\ndown 4\nforward 9\nup 8\ndown 9\nforward 5\ndown 3\nup 7\nforward 5\ndown 2\nforward 2\nforward 1\ndown 6\ndown 1\nup 4\nforward 4\ndown 7\nup 3\ndown 3\nforward 4\nforward 2\nforward 1\nforward 9\nforward 7\nforward 9\ndown 1\nforward 6\ndown 5\nup 7\ndown 9\ndown 2\nup 4\ndown 2\ndown 2\ndown 1\ndown 2\ndown 5\ndown 4\ndown 6\nforward 4\nforward 9\nforward 6\nforward 7\nup 9\nup 2\nforward 7\nforward 9\nup 5\nforward 2\nup 5\nup 9\nforward 9\ndown 8\nforward 6\ndown 8\nforward 4\ndown 1\ndown 4\nforward 4\ndown 3\nforward 4\nforward 3\nforward 9\ndown 5\nforward 3\nforward 2\nforward 9\nforward 8\ndown 2\nforward 5\nup 2\nforward 1\ndown 3\nforward 9\nforward 8\ndown 9\nforward 5\nforward 1\ndown 1\ndown 9\ndown 6\nforward 8\nforward 1\ndown 5\nup 5\nforward 8\nup 4\ndown 6\nforward 1\nup 2\ndown 3\ndown 1\ndown 8\ndown 2\nup 6\ndown 5\ndown 2\ndown 3\nforward 2\nup 7\ndown 9\nup 1\nup 1\nforward 7\nforward 4\ndown 7\nup 7\ndown 7\nforward 5\nup 2\ndown 7\nforward 9\ndown 7\nup 4\nforward 2\nforward 1\nup 6\ndown 8\nup 6\ndown 2\ndown 4\nup 8\nup 8\nup 5\ndown 6\nup 6\ndown 5\nup 1\ndown 1\nforward 6\nup 7\nforward 8\nup 9\ndown 8\nup 7\nforward 9\nup 4\ndown 5\nforward 3\nforward 6\nforward 4\nforward 4\ndown 7\nforward 9\ndown 6\ndown 2\nforward 9\nforward 3"

	inputFile := os.Args[1]
	if len(inputFile) == 0 {
		inputFile = "/test.text"
	}

	inputFromFile, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Could not read file", inputFile)
		return
	}
	inputStr := string(inputFromFile)

	test1 := calculatePosition(test)
	fmt.Println(test1)

	test2 := calculatePositionWithAim(test)
	fmt.Println(test2)

}

func buildDirectionList(input string) instructions {
	inputSlice := strings.Split(input, "\n")
	result := instructions{}

	for _, val := range inputSlice {
		currStep := strings.Split(val, " ")
		dir := currStep[0]
		dist, err := strconv.Atoi(currStep[1])
		if err != nil {
			fmt.Println("Failed to convert string to number", currStep)
			return result
		}
		result = append(result, pair{dir, dist})
	}
	return result
}

func calculatePosition(input string) string {
	insts := buildDirectionList(input)
	if len(insts) < 1 {
		return "Could not get instructions"
	}
	x := 0
	y := 0

	for _, currPair := range insts {
		switch currPair.string {
		case Forward:
			x += currPair.int
		case Down:
			y += currPair.int
		case Up:
			y -= currPair.int
		}
	}

	return fmt.Sprintf("Result\nx:%d\ny:%d\nprod:%d\n", x, y, x * y)
}

func calculatePositionWithAim(input string) string {
	insts := buildDirectionList(input)
	if len(insts) < 1 {
		return "Could not get instructions"
	}
	x := 0
	y := 0
	aim := 0

	for _, currPair := range insts {
		switch currPair.string {
		case Forward:
			x += currPair.int
			y += currPair.int * aim
		case Down:
			aim += currPair.int
		case Up:
			aim -= currPair.int
		}
	}

	return fmt.Sprintf("Result\nx:%d\ny:%d\naim:%d\nprod:%d\n", x, y, aim, x * y)
}
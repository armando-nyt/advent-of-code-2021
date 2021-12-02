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

	res1 := calculatePosition(inputStr)
	fmt.Println(res1)

	res2 := calculatePositionWithAim(inputStr)
	fmt.Println(res2)

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
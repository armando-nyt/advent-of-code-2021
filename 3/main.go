package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/armando-nyt/advent-of-code-2021/utils"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("Provide input file")
		return
	}
	inputFilename := os.Args[1]
	input, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	inputList, err := utils.SplitIntoList(string(input))
	if err != nil {
		fmt.Println("Could not get string input: ", err)
	}
	zerosCount, onesCount := tallyAllNumbers(inputList)
	gama := findRelation("gama", zerosCount, onesCount)
	epsylon := findRelation("epsylon", zerosCount, onesCount)
	gamaNum, err := strconv.ParseInt(strings.Trim(strings.Replace(fmt.Sprint(gama), " ", "", -1), "[]"), 2, 32)
	epsylonNum, err := strconv.ParseInt(strings.Trim(strings.Replace(fmt.Sprint(epsylon), " ", "", -1), "[]"), 2, 32)

	fmt.Printf("Result\nZeros: %v\nOnes: %v\nGama: %v\nGamaNum: %v\nEpsylon: %v\nRate: %v\n", zerosCount, onesCount, gama, gamaNum, epsylon, gamaNum * epsylonNum)
}

// take the input and turn it into a list
// have two list tracking count of zero and one
// visit every element and count the values for each position
// update the count for each position
func tallyAllNumbers(input []string) (zeros []int, ones[]int){
	inputsLength := len(input[0])
	fmt.Println("length: ", inputsLength)
	zeros = make([]int, inputsLength)
	ones = make([]int, inputsLength)

	for _, number := range input{
		for key, digit := range number {
			if string(digit) == "0" {
				zeros[key] += 1
			} else {
				ones[key] += 1
			}
		}
	}

	return zeros, ones
}

func findRelation(t string, zeros []int, ones []int) []int {
	result := make([]int, len(zeros))
	compare := greater
	if t != "gama" {
		compare = lesser
	}

	for i := 0; i < len(zeros); i++ {
		if compare(zeros[i], ones[i]) {
			result[i] = 0
		} else {
			result[i] = 1
		}
	}
	return result
}

func greater(num1 int, num2 int) bool {
	return num1 > num2
}

func lesser(num1 int, num2 int) bool {
	return num1 < num2
}
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
	o2Num, err := strconv.ParseInt(reduceListByCriteria("o2", inputList), 2, 64)
	co2Num, err := strconv.ParseInt(reduceListByCriteria("co2", inputList), 2, 64)
	fmt.Println("o2", o2Num)
	fmt.Println("co2", co2Num)

	o2, _ := strconv.ParseInt(filterListByCriteria("o2", inputList), 2, 64)
	co2, _ := strconv.ParseInt(filterListByCriteria("co2", inputList), 2, 64)
	fmt.Printf("Result Using Filter\nO2:%d\nCo2:%d\nProduct:%d\n", o2, co2, o2 * co2)
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

// use tally allNums
// o2 get the most common, if equal favour 1
// co2 get the least common, if equal favour 0

// countListByPosition - traverses a list and returns a list of elements that match the comparison
func reduceListByCriteria(criteria string, list []string) string {
	// get count of 0s and 1s
	zerosCount, onesCount := tallyAllNumbers(list)
	compare := greater
	if criteria != "o2" {
		compare = lesser
	}
	prioritize := "1"
	//delegate := "0"
	if criteria != "o2" {
		prioritize = "0"
		//delegate = "1"
	}

	// create the values that are expected
	expected := ""
	for i, _ := range zerosCount {
		if zerosCount[i] == onesCount[i] {
			expected += prioritize
		} else if compare(zerosCount[i], onesCount[i]) {
			expected += "0"
		} else {
			expected += "1"
		}
	}
	fmt.Println("This is expected", expected)

	filtered := make(map[string]string)
	for _, num := range list {
		filtered[num] = num
	}

	for key, char := range expected {
		for currNum := range filtered {
			if string(currNum[key]) != string(char) {
				delete(filtered, currNum)
			}
			if len(filtered) == 1 {
				fmt.Println(filtered, len(filtered))
				for _, res := range filtered {
					return res
				}
			}
		}
	}

	return ""
}

// countListByPosition - traverses a list and returns a list of elements that match the comparison
func filterListByCriteria(criteria string, list []string) string {
	// get count of 0s and 1s
	compare := greater
	if criteria != "o2" {
		compare = lesser
	}

	// go through the list and separate the numbers based on 1 and 0 for given digit
	filteredList := list
	for i := 0; i < len(list[0]); i++ {
		zeroList := []string{}
		oneList := []string{}
		for _, num := range filteredList {
			if string(num[i]) == "0" {
				zeroList = append(zeroList, num)
			} else {
				oneList = append(oneList, num)
			}
		}
		if len(zeroList) == 0 {
			filteredList = oneList
		} else if len(oneList) == 0 {
			filteredList = zeroList
		} else if len(zeroList) == len(oneList) {
			if criteria == "o2" {
				filteredList = oneList
			} else {
				filteredList = zeroList
			}
		} else if compare(len(zeroList), len(oneList)) {
			filteredList = zeroList
		} else {
			filteredList = oneList
		}
	}

	//fmt.Println("list", list, "  filetered", filteredList)
	return filteredList[0]
}
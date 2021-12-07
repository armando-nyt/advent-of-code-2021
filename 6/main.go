package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fishCount map[int]int

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
	fList := createFishList(string(input))

	fmt.Println(fList)
	//for i:= 0; i < 256; i++ {
	//	fList = calculateNewFish(fList)
	//}
	fCount := make([]int, 9)
	for _, f := range fList {
		//if _, ok := fCount[f.day]; !ok {
		//	fCount[f.day] = 0
		//}
		fCount[f.day] += 1
	}
	fmt.Printf("InitialCount: %v\n", fCount)

	for i:= 0; i < 256; i++ {
		tempCount := make([]int, 9)
		sixOverFlow := 0
		for day := 0; day <= 8; day++ {
			if day == 0 {
				tempCount[8] = fCount[day]
				sixOverFlow = fCount[day]
			} else {
				tempCount[day - 1] = fCount[day]
			}
		}
		tempCount[6] += sixOverFlow
		//fmt.Println(tempCount)
		fCount = tempCount
	}

	tot := 0
	for _, c := range fCount {
		tot += c
	}
	fmt.Println(fCount, tot)
}

type fish struct {
	day int
}

type fishList []fish

func createFishList(input string) fishList {
	tempList := strings.Split(input, ",")

	var list []fish
	for _, age := range tempList {
		day, err := strconv.Atoi(age)
		if err != nil {
			panic("Could not convert fish age to int")
		}
		list = append(list, fish{day: day})
	}
	return list
}

func (f *fish) resetDay()  {
	f.day = 6
}

func (f *fish) decrementDay()  {
	f.day -= 1
}

func calculateNewFish(l fishList) fishList {
	var newList fishList
	var newFish fishList
	for _, f := range l {
		if f.day == 0 {
			f.resetDay()
			newFish = append(newFish, fish{day: 8})
		} else {
			f.decrementDay()
		}
		newList = append(newList, f)
	}
	newList = append(newList, newFish...)
	return newList
}
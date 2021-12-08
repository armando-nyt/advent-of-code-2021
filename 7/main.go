package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("Provide input file")
		return
	}
	fileName := os.Args[1]
	puzInput, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	puzList := strings.Split(string(puzInput), ",")
	var crabList []int
	for _, puzItem := range puzList {
		currInt, err := strconv.Atoi(puzItem)
		if err != nil {
			fmt.Println("Failed to convert puzItem to int")
			return
		}
		crabList = append(crabList, currInt)
	}
	least, max := findExtremes(crabList)
	midPoint := (max + least) / 2
	fuel := findBestFuel(crabList, midPoint, math.MaxInt32)
	fmt.Println("TotalFuel: ", fuel)
}

// given 16,1,2,0,4,2,7,1,2,14
// brute force
// get the two extremes
// recurse through the different permutations
// choose the midpoint between extremes
// calculate cost of moving all pieces to that point
// recurse on both ends
// if greater exit otherwise continue

func findExtremes(list []int) (int, int) {
	least := list[0]
	max := list[0]

	for _, curr := range list {
		if least > curr {
			least = curr
		} else if max < curr {
			max = curr
		}
	}
	return least, max
}

var steps = map[int]int{}

// findBestFuel recurses until it finds the smallest values
func findBestFuel(list []int, mid int, fuel int) int {
	if mid >= len(list) || mid < 0 {
		return math.MaxInt32
	}

	totFuel := 0
	for _, curr := range list {
		step := abs(curr - mid)
		if _, ok := steps[step]; !ok {
			steps[step] = calcExpense(step)
		}
		totFuel += steps[step]
	}

	if totFuel > fuel {
		return fuel
	}
	up := findBestFuel(list, mid + 1, totFuel)
	down := findBestFuel(list, mid - 1, totFuel)

	return min(up, down)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcExpense(n int) int {
	tot := 0
	for i := 1; i <= n; i++ {
		tot += i
	}
	return tot
}
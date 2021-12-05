package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	list := strings.Split(string(input),"\n\n")
	drawnNums := strings.Split(list[0], ",")
	boards := buildBoards(list[1:])
	//fmt.Printf("Nums: %s\nBoards: %v\n", drawnNums, boards)

	var boardStructs []BoardMap
	for _, board := range boards {
		boardStructs = append(boardStructs, buildBoardMap(board))
	}
	fmt.Printf("%+v\n  %d", boardStructs[0], len(boardStructs[0].Nums))
	type winners struct {
		Board BoardMap
		Rem int
		CurrNum int
		Prod int
	}
	winingBoards := map[int]winners{}
	lastKey := -1
	for _, drawnNum := range drawnNums {
		for key, bStruct := range boardStructs {
			if _, ok := winingBoards[key]; ok {
				continue
			}
			checkAndAdd(bStruct, drawnNum)
			bingo := checkBingo(bStruct, drawnNum)
			delete(bStruct.Nums, drawnNum)
			if bingo {
				rems := calcAllRemaining(bStruct)
				currNum, err := strconv.Atoi(drawnNum)
				if err != nil {
					panic("Failed to convert drawn num")
				}
				fmt.Printf("Result\nRems: %v\nCurrNum: %v\nProd: %v\nBoard: %+v\n\n", rems, drawnNums, rems * currNum, bStruct)
				//return // ignore for second step
				winingBoards[key] = winners{
					Board: bStruct,
					Rem: rems,
					CurrNum: currNum,
					Prod: rems * currNum,
				}
				lastKey = key
			}
		}
	}

	fmt.Printf("Last Board: %+v\n", winingBoards[lastKey])

}

func buildBoards(list []string) [][][]string {
	var res [][][]string

	for _, board := range list {
		temp := strings.Split(board, "\n")
		var currBoard [][]string
		for _, row := range temp {
			currBoard = append(currBoard, strings.Fields(row))
		}
		res = append(res, currBoard)
	}
	return res
}



// take the input and separate it into the drawn nums and boards
// for each board we create a map of the numbers and their position
// each row and column get a counter once that counter meets 5 it's a bingo

// create a map of each board
	// contains maps for the nums-> position, col count, row count
// when a number is present we get coords and update the row and col counters if it is complete we call bingo
type BoardMap struct {
	Nums map[string][]int // num: [x, y] coords
	Rows map[int]int // row position and count of found numbers
	Cols map[int]int
}

func buildBoardMap(board [][]string) BoardMap {
	nums := map[string][]int{}
	rows := map[int]int{}
	cols := map[int]int{}

	for rowPos, rowList := range board {
		rows[rowPos] = 0
		cols[rowPos] = 0 // being a square, using the row position to initialize the column map as well

		for col, colVal := range rowList {
			nums[colVal] = []int{rowPos, col}
		}
	}

	return BoardMap{
		Nums: nums,
		Rows: rows,
		Cols: cols,
	}
}

func checkAndAdd(board BoardMap, num string) {
	if board.Nums[num] != nil {
		coords := board.Nums[num]
		board.Rows[coords[0]] += 1
		board.Cols[coords[1]] += 1
	}
}

func checkBingo(board BoardMap, num string) bool {
	if board.Nums[num] != nil {
		coords := board.Nums[num]
		if board.Rows[coords[0]] == 5 || board.Cols[coords[1]] == 5 {
			return true
		}
	}
	return false
}

func calcAllRemaining(board BoardMap) int {
	tot := 0
	for num, _ := range board.Nums {
		curr, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		tot += curr
	}
	return tot
}
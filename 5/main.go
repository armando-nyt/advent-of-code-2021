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

	dirList := strings.Split(string(input), "\n")
	dirs := createOrderedDirs(dirList, true) // update for step two
	//fmt.Println(dirs)
	overLaps := countOverlaps(dirs)
	//fmt.Println(overLaps)

	counter := 0
	for _, count := range overLaps {
		if count > 1 {
			counter += 1
		}
	}
	fmt.Println("Count of overlaps is: ", counter)
}

type coords struct {
	x int
	y int
}

type line struct {
	start coords
	end coords
}

func buildDir(dir string) line {
	points := strings.Split(dir, " -> ")
	tempCoords := []coords{}

	for _, currPoint := range points {
		pointVals := strings.Split(currPoint, ",")
		x, err := strconv.Atoi(pointVals[0])
		if err != nil {
			panic("could not convert x")
		}
		y, err := strconv.Atoi(pointVals[1])
		if err != nil {
			panic("could not convert y")
		}
		tempCoords = append(tempCoords, coords{x: x, y: y})
	}

	res := line{}
	if tempCoords[0].x != tempCoords[1].x && tempCoords[0].y != tempCoords[1].y {
		// handle diags
		if tempCoords[0].x < tempCoords[1].x {
			res.start = tempCoords[0]
			res.end = tempCoords[1]
		} else {
			res.start = tempCoords[1]
			res.end = tempCoords[0]
		}
	} else if tempCoords[0].x == tempCoords[1].x {
		if tempCoords[0].y > tempCoords[1].y {
			res.start = tempCoords[1]
			res.end = tempCoords[0]
		} else {
			res.start = tempCoords[0]
			res.end = tempCoords[1]
		}
	} else {
		if tempCoords[0].x > tempCoords[1].x {
			res.start = tempCoords[1]
			res.end = tempCoords[0]
		} else {
			res.start = tempCoords[0]
			res.end = tempCoords[1]
		}
	}
	return res
}

func createOrderedDirs(input []string, diags bool) []line {
	var res []line

	for _, currLine := range input {
		dir := buildDir(currLine)
		if !diags {
			if dir.start.x != dir.end.x && dir.start.y != dir.end.y {
				continue
			}
		}
		res = append(res, dir)
	}
	return res
}

type count map[string]int

// iterate over dirs and add counts for each point in a line
func countOverlaps(dirs []line) count {
	overlaps := count{}

	for _, line := range dirs {
		// handle diags
		if line.start.x != line.end.x && line.start.y != line.end.y {
			// x always increments, so we need to check the direction of y
			difY := line.end.y - line.start.y
			stepY := 1
			if difY < 0 {
				stepY = -1
			}
			for i, j := line.start.x, line.start.y; i <= line.end.x; i, j = i + 1, j + stepY {
				strI := strconv.FormatInt(int64(i), 10)
				strJ := strconv.FormatInt(int64(j), 10)
				overlaps[strI + "," + strJ] += 1
			}

		} else {
			start := line.start.x
			end := line.end.x
			cons := line.start.y
			isX := true
			if start == end {
				start = line.start.y
				end = line.end.y
				cons = line.start.x
				isX = false
			}
			for i := start; i <= end; i++ {
				strI := strconv.FormatInt(int64(i), 10)
				strCons := strconv.FormatInt(int64(cons), 10)
				if isX {
					overlaps[strI + "," + strCons] += 1
				} else {
					overlaps[strCons + "," + strI] += 1
				}
			}
		}
	}
	return overlaps
}
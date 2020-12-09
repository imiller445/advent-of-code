package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	boardingPassesRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	boardingPasses := strings.Split(string(boardingPassesRaw), "\r\n")
	highestId := 0
	existingBoardingPasses := make([]int, 971)
	for _, boardingPass := range boardingPasses {
		id := getBoardingPassId(boardingPass, 127, 7)

		if id >= highestId {
			highestId = id
		}

		existingBoardingPasses[id] = id
	}

	var missingPasses []int
	for i := 0; i < 971; i++ {
		if existingBoardingPasses[i] != i {
			missingPasses = append(missingPasses, i)
		}
	}
	mySeat := 0
	for _, missingPass := range missingPasses {
		if existingBoardingPasses[missingPass-1] != 0 && existingBoardingPasses[missingPass+1] != 0 {
			mySeat = missingPass
			break
		}
	}
	fmt.Println("My seat is: ", mySeat)
	fmt.Println("Highest ID is", highestId)
}

func getBoardingPassId(boardingPass string, numRows int, numCols int) (id int) {
	curRowSectionLow := 0
	curRowSectionHigh := numRows
	curColSectionLow := 0
	curColSectionHigh := numCols
	curIndex := 0
	for curRowSectionHigh != curRowSectionLow || curColSectionHigh != curColSectionLow {
		//fmt.Println(string(boardingPass[curIndex]), curRowSectionHigh, curRowSectionLow, curColSectionHigh, curColSectionLow)
		if string(boardingPass[curIndex]) == "F" {
			curRowSectionHigh = curRowSectionHigh - ((curRowSectionHigh - curRowSectionLow) / 2) - 1
		} else if string(boardingPass[curIndex]) == "B" {
			curRowSectionLow = (curRowSectionLow + (curRowSectionHigh-curRowSectionLow)/2) + 1
		} else if string(boardingPass[curIndex]) == "L" {
			curColSectionHigh = curColSectionHigh - ((curColSectionHigh - curColSectionLow) / 2) - 1
		} else if string(boardingPass[curIndex]) == "R" {
			curColSectionLow = (curColSectionLow + (curColSectionHigh-curColSectionLow)/2) + 1
		} else {
			fmt.Println("YIKES")
		}
		curIndex++
	}
	return (curRowSectionHigh * 8) + curColSectionHigh
}

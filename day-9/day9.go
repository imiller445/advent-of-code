package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	xmasValsRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	xmasValsStrings := strings.Split(string(xmasValsRaw), "\r\n")
	xmasVals := make([]int, len(xmasValsStrings))
	for i, xmasValString := range xmasValsStrings {
		xmasVal, _ := strconv.Atoi(xmasValString)
		xmasVals[i] = xmasVal
	}
	currentValue := xmasVals[25]
out:
	for {
		for i := 25; i < len(xmasVals); i++ {
			currentValue = xmasVals[i]
			isValid := valid(xmasVals[i-25:i], currentValue)
			if !isValid {
				break out
			}
		}

	}
	fmt.Println("First value was", currentValue)
	fmt.Println("Sum high low is", findValidContiguousSubsetForSum(currentValue, xmasVals))
}

func valid(xmasVals []int, xmasValToCheck int) (valid bool) {
	for i := 0; i < len(xmasVals); i++ {
		for j := 0; j < len(xmasVals); j++ {
			if i != j && xmasVals[i]+xmasVals[j] == xmasValToCheck {
				return true
			}
		}
	}
	return false
}

func sumSlice(slice []int) (sum int) {
	sum = 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func findHighAndLowValInSlice(slice []int) (high int, low int) {
	high = slice[0]
	low = slice[0]
	for _, v := range slice {
		if v > high {
			high = v
		} else if v < low {
			low = v
		}
	}
	return high, low
}

func findValidContiguousSubsetForSum(sum int, xmasVals []int) (highLowSum int) {
out:
	for i := 0; i < len(xmasVals); i++ {
		for j := i + 1; j < len(xmasVals); j++ {
			curSum := sumSlice(xmasVals[i:j])
			if curSum == sum {
				high, low := findHighAndLowValInSlice(xmasVals[i:j])
				return high + low
			} else if curSum > sum {
				continue out
			}
		}
	}
	return 0
}

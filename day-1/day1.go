package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	expenseReportBinary, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	expenseReport, err2 := ReadInts(strings.NewReader(string(expenseReportBinary)))
	if err2 != nil {
		fmt.Println("Error reading file", err)
		return
	}
	part_1(expenseReport)
	part_2(expenseReport)
}

func part_1(expenseReport []int) {
out:
	for i := 0; i < len(expenseReport); i++ {
		for j := 0; j < len(expenseReport); j++ {
			if i != j && expenseReport[i]+expenseReport[j] == 2020 {
				answer := expenseReport[i] * expenseReport[j]
				fmt.Println("Part 1 answer is", answer)
				break out
			}
		}

	}
}

func part_2(expenseReport []int) {
out:
	for i := 0; i < len(expenseReport); i++ {
		for j := 0; j < len(expenseReport); j++ {
			for k := 0; k < len(expenseReport); k++ {
				if i != j && expenseReport[i]+expenseReport[j]+expenseReport[k] == 2020 {
					answer := expenseReport[i] * expenseReport[j] * expenseReport[k]
					fmt.Println("Part 2 answer is", answer)
					break out
				}
			}
		}
	}
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

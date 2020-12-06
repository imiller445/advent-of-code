package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	passwordBytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	passwordInputArray := strings.Split(string(passwordBytes), "\n")
	part1ValidCount := 0
	part2ValidCount := 0
	for _, passwordRule := range passwordInputArray {
		firstCut := strings.Split(passwordRule, "-")
		firstValue, _ := strconv.Atoi(firstCut[0])
		secondCut := strings.Split(firstCut[1], " ")
		secondValue, _ := strconv.Atoi(secondCut[0])
		letter := strings.Split(secondCut[1], ":")[0]
		password := secondCut[2]
		//fmt.Println(firstValue, secondValue, letter,password)
		if isValidPart1(firstValue, secondValue, letter, password) {
			part1ValidCount++
		}
		if isValidPart2(firstValue, secondValue, letter, password) {
			part2ValidCount++
		}
	}

	fmt.Println("The number of valid passwords for part 1 is: ", part1ValidCount)
	fmt.Println("The number of valid passwords for part 2 is: ", part2ValidCount)
}

func isValidPart1(low int, high int, letter string, password string) (valid bool) {
	occurrences := strings.Count(password, letter)
	return occurrences >= low && occurrences <= high
}

func isValidPart2(firstIndex int, secondIndex int, letter string, password string) (valid bool) {
	i := password[firstIndex-1]
	j := password[secondIndex-1]
	isValid := ((i == letter[0] && j != letter[0]) ||
		(i != letter[0] && j == letter[0])) &&
		(i == letter[0] || j == letter[0])
	return isValid
}

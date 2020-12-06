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
	validCount := 0
	for _, passwordRule := range passwordInputArray {
		lowCut := strings.Split(passwordRule, "-")
		low, _ := strconv.Atoi(lowCut[0])
		highCut := strings.Split(lowCut[1], " ")
		high, _ := strconv.Atoi(highCut[0])
		letter := strings.Split(highCut[1], ":")[0]
		password := highCut[2]
		if isValid(low, high, letter, password) {
			validCount++
		}
	}

	fmt.Println("The number of valid passwords is: ", validCount)
}

func isValid(low int, high int, letter string, password string) (valid bool) {
	occurrences := strings.Count(password, letter)
	return occurrences >= low && occurrences <= high
}

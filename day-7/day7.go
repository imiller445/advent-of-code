package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	baggageRulesRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	baggageRulesSlice := strings.Split(string(baggageRulesRaw), ".\r\n")
	var baggageReverseMap = map[Bag][]Bag{}
	var baggageRuleMap = map[Bag][]BagRule{}
	for _, baggageRule := range baggageRulesSlice {
		baggageRuleSplitForContain := strings.Split(baggageRule, "contain")
		outerBag := extractBag(baggageRuleSplitForContain[0])
		baggageRuleSplitForMultiples := strings.Split(baggageRuleSplitForContain[1], ",")
		for _, bagString := range baggageRuleSplitForMultiples {
			bagString = strings.TrimSpace(bagString)
			bagSlice := strings.SplitN(bagString, " ", 2)
			count, _ := strconv.Atoi(bagSlice[0])
			bag := extractBag(bagSlice[1])
			baggageRuleMap[outerBag] = append(baggageRuleMap[outerBag], BagRule{bag: bag, count: count})
			baggageReverseMap[bag] = append(baggageReverseMap[bag], outerBag)
		}
	}
	goldBag := Bag{color: "gold", pattern: "shiny"}
	bagsContainingBag := findAllBagsAllowedToContain(goldBag, baggageReverseMap)
	fmt.Println("Bags that can contain bag: ", bagsContainingBag.Cardinality()-1)
	fmt.Println("num of bags in bag: ", findNumberOfBagsRequiredForBagToContain(goldBag, baggageRuleMap)-1)
}

type Bag struct {
	color   string
	pattern string
}

type BagRule struct {
	bag   Bag
	count int
}

func extractBag(bagString string) (bag Bag) {
	bagSlice := strings.Split(bagString, " ")
	return Bag{
		color:   bagSlice[1],
		pattern: bagSlice[0],
	}
}

func findNumberOfBagsRequiredForBagToContain(bag Bag, bagMap map[Bag][]BagRule) (numberOfBags int) {
	count := 1
	for _, bagRule := range bagMap[bag] {
		count += bagRule.count * findNumberOfBagsRequiredForBagToContain(bagRule.bag, bagMap)
	}
	return count
}

func findAllBagsAllowedToContain(bag Bag, bagMap map[Bag][]Bag) (bags mapset.Set) {
	var bagsSet = mapset.NewSet()
	bagsSet.Add(bag)
	for _, deepBag := range bagMap[bag] {
		deepBagsSet := findAllBagsAllowedToContain(deepBag, bagMap)
		bagsSet = bagsSet.Union(deepBagsSet)
	}
	return bagsSet
}

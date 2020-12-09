package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"io/ioutil"
	"strings"
)

func main() {
	customsFormsRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	customsForms := strings.Split(string(customsFormsRaw), "\r\n\r\n")
	sumOfYes := 0
	sumAllAnsweredYes := 0
	for _, customsForm := range customsForms {
		answers := strings.Split(customsForm, "\r\n")
		var answerSetsList []mapset.Set
		for _, answer := range answers {
			answerAsSet := mapset.NewSet()
			for _, a := range answer {
				answerAsSet.Add(string(a))
			}
			answerSetsList = append(answerSetsList, answerAsSet)
		}
		var allAnsweredYes = answerSetsList[0]
		var anyAnsweredYes = answerSetsList[0]
		for i := 1; i < len(answerSetsList); i++ {
			allAnsweredYes = allAnsweredYes.Intersect(answerSetsList[i])
			anyAnsweredYes = anyAnsweredYes.Union(answerSetsList[i])
		}
		sumOfYes += anyAnsweredYes.Cardinality()
		sumAllAnsweredYes += allAnsweredYes.Cardinality()
	}
	fmt.Println("number of yes: ", sumOfYes)
	fmt.Println("All answered yes: ", sumAllAnsweredYes)
}

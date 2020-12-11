package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instructionsRaw, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	instructionsSlice := strings.Split(string(instructionsRaw), "\r\n")
	instructionsList := make([]Instruction, len(instructionsSlice))
	for i, instructionString := range instructionsSlice {
		r := regexp.MustCompile("^([a-z]+) (-|\\+)([0-9]+)$")
		matched := r.FindStringSubmatch(instructionString)
		command := matched[1]
		sign := matched[2]
		value, _ := strconv.Atoi(matched[3])
		instruction := Instruction{command: command, sign: sign, value: value}
		instructionsList[i] = instruction
	}
	accPart1, _ := traverseInstructionSetStopAtInfinite(instructionsList, 0, make([]bool, len(instructionsList)), true, 0)
	fmt.Println("Part 1 acc is", accPart1)
	acc, _ := traverseInstructionSetStopAtInfinite(instructionsList, 0, make([]bool, len(instructionsList)), false, 0)
	fmt.Println("final acc is", acc)
}

type Instruction struct {
	command string
	sign    string
	value   int
}

func traverseInstructionSetStopAtInfinite(
	instructions []Instruction,
	curIndex int,
	visitedIndexes []bool,
	flipped bool,
	startingAcc int,
) (acc int, foundTheEnd bool) {
	instruction := instructions[curIndex]
	for {
		if curIndex == len(visitedIndexes) {
			break
		}
		if visited := visitedIndexes[curIndex]; visited {
			break
		}
		visitedIndexes[curIndex] = true
		switch instruction.command {
		case "nop":
			if !flipped {
				newInstructions := make([]Instruction, len(instructions))
				copy(newInstructions, instructions)
				newInstructions[curIndex] = Instruction{command: "jmp", value: instruction.value, sign: instruction.sign}
				newVisitedIndexes := make([]bool, len(visitedIndexes))
				copy(newVisitedIndexes, visitedIndexes)
				newVisitedIndexes[curIndex] = false
				a, end := traverseInstructionSetStopAtInfinite(newInstructions, curIndex, newVisitedIndexes, true, startingAcc)
				if end {
					return a, end
				}
			}
			curIndex += 1
			instruction = instructions[curIndex]
		case "acc":
			if instruction.sign == "+" {
				startingAcc += instruction.value
			} else if instruction.sign == "-" {
				startingAcc -= instruction.value
			}
			curIndex += 1
			instruction = instructions[curIndex]
		case "jmp":
			if !flipped {
				newInstructions := make([]Instruction, len(instructions))
				copy(newInstructions, instructions)
				newInstructions[curIndex] = Instruction{command: "nop", value: instruction.value, sign: instruction.sign}
				newVisitedIndexes := make([]bool, len(visitedIndexes))
				copy(newVisitedIndexes, visitedIndexes)
				newVisitedIndexes[curIndex] = false
				a, end := traverseInstructionSetStopAtInfinite(newInstructions, curIndex, newVisitedIndexes, true, startingAcc)
				if end {
					return a, end
				}
			}
			if instruction.sign == "+" {
				curIndex += instruction.value
			} else if instruction.sign == "-" {
				curIndex -= instruction.value
			}
			if curIndex == len(instructions) {
				break
			}
			instruction = instructions[curIndex]
		}
	}
	if curIndex == len(instructions) {
		return startingAcc, true
	} else {
		return startingAcc, false
	}
}

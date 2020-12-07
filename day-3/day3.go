package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	treeMap, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	answer := 1
	traversals := []slopeTraversal{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, traversal := range traversals {
		answer *= getNumberOfTreesInSlopeTraversal(traversal.right, traversal.down, treeMap)
	}
	fmt.Println("Answer is", answer)
}

type slopeTraversal struct {
	right int
	down  int
}

func getNumberOfTreesInSlopeTraversal(
	right int,
	down int,
	treeMap []uint8,
) (numTrees int) {
	tree := []uint8("#")[0]
	rowLength := bytes.IndexByte(treeMap, 10) + 1
	stopLength := len(treeMap) - rowLength
	currentLocation := 0
	currentCol := 0
	treeCount := 0
	for currentLocation <= stopLength {
		currentCol += right
		if currentCol >= (rowLength - 2) {
			currentCol -= rowLength - 2
			currentLocation -= rowLength - 2
		}
		currentLocation += (down * rowLength) + right
		if treeMap[currentLocation] == tree {
			treeCount++
		}
	}
	return treeCount
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_BRANCH = 10
const MAX_DEPTH = 10

type TreeNode struct {
	value    int
	branches []*TreeNode
}

func randomTreeProperties() {

	rand.Seed(time.Now().UnixNano())
	depth := rand.Intn(MAX_DEPTH-1) + 1

	// fmt.Println("Generating Random Tree...")
	var tree = generateRandomTree(depth)

	// fmt.Println("Calculating Random Tree Properties...")
	mean := calculateTreeMean(tree)
	fmt.Println("Random Tree Mean: ", mean)
}

func generateRandomTree(depth int) *TreeNode {

	// fmt.Println("Depth: ", depth)
	if depth == 0 {
		return nil
	}

	var node = &TreeNode{value: rand.Intn(MAX_VALUE)}
	// fmt.Println("Value: ", node.value)
	if depth == 1 {
		return node
	}

	nBranches := rand.Intn(MAX_BRANCH-1) + 1
	// fmt.Println("Branch Size: ", nBranches)
	node.branches = make([]*TreeNode, nBranches)
	for b := 0; b < nBranches; b++ {
		node.branches[b] = generateRandomTree(depth - 1)
	}

	return node
}

func value_property(node *TreeNode) int {
	if node != nil {
		return node.value
	} else {
		return 0
	}
}

func count_property(node *TreeNode) int {
	if node != nil {
		return 1
	} else {
		return 0
	}
}

func traverseTree(node *TreeNode, property func(node *TreeNode) int) int {

	if node == nil {
		return 0
	}

	// fmt.Println("Node: ", node.value)
	prop_val := property(node)

	// fmt.Println("Branch Size: ", len(node.branches))
	for b := 0; b < len(node.branches); b++ {
		prop_val += traverseTree(node.branches[b], property)
	}

	return prop_val
}

func calculateTreeMean(tree *TreeNode) float64 {

	sum := traverseTree(tree, value_property)
	count := traverseTree(tree, count_property)
	fmt.Println("Random Tree Count: ", count)

	if count != 0 {
		return float64(sum) / float64(count)
	}

	return 0
}

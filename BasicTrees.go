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
	var tree = generateRandomTree(depth)

	depth = getTreeDepth(tree)
	fmt.Println("Random Tree Depth: ", depth)
}

func generateRandomTree(depth int) *TreeNode {
	fmt.Println("Depth: ", depth)

	if depth == 0 {
		return nil
	}

	var node = &TreeNode{value: rand.Intn(MAX_VALUE)}
	fmt.Println("Value: ", node.value)
	if depth == 1 {
		fmt.Println("Branch Size: ", 0)
		return node
	}

	nBranches := rand.Intn(MAX_BRANCH-1) + 1
	fmt.Println("Branch Size: ", nBranches)
	node.branches = make([]*TreeNode, nBranches)

	for b := 0; b < nBranches; b++ {
		node.branches[b] = generateRandomTree(depth - 1)
	}

	return node
}

func traverseTree(node *TreeNode) {

	if node == nil {
		return
	}

	fmt.Println("Node: ", node.value)
	fmt.Println("Branch Size: ", len(node.branches))

	for b := 0; b < len(node.branches); b++ {
		traverseTree(node.branches[b])
	}
}

func getTreeDepth(tree *TreeNode) int {
	return 0
}

func getRandomTreePath(tree *TreeNode) []int {
	return nil
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
}

func randomListProperties() {
	// fmt.Println("Generating Random List...")

	var list = generateRandomList()

	len := getListLength(list)
	fmt.Println("Random List Length: ", len)

	start := time.Now()
	mean := calculateListMean(list)
	fmt.Println("Random List Mean: ", mean)
	elapsed := time.Since(start)
	fmt.Println("Normal Calculate Random List Mean in: ", elapsed.String())
}

func generateRandomList() *Node {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(MAX_SIZE)

	var list = &Node{value: rand.Intn(MAX_VALUE)}
	iter := list
	for i := 1; i < size; i++ {
		iter.next = &Node{value: rand.Intn(MAX_VALUE)}
		// fmt.Println(iter.value)
		iter = iter.next
	}
	iter.next = nil

	return list
}

func getListLength(list *Node) int {
	iter := list
	len := 0
	for ; iter != nil; len++ {
		iter = iter.next
	}
	return len
}

func calculateListMean(list *Node) float64 {
	iter := list
	sum := 0
	len := 0
	for iter != nil {
		sum += iter.value
		len++
		iter = iter.next
	}
	mean := float64(sum) / float64(len)
	return mean
}

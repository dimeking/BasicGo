package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateRandomIntArray() []int {
	const MAX_SIZE = 1000
	const MAX_VALUE = 100

	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(MAX_SIZE)
	intArray := make([]int, size)

	for i := 0; i < len(intArray); i++ {
		intArray[i] = rand.Intn(MAX_VALUE)
	}

	return intArray[:len(intArray)]
}

func calculateMean(intArray []int) float64 {

	sum := 0
	for _, v := range intArray {
		sum += v
	}

	var mean float64 = float64(sum) / float64(len(intArray))
	return mean
}

func calculateMedian(array []int) float64 {
	median := 0.0
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	if len(array)%2 == 1 {
		median = float64(array[len(array)/2])
	} else {
		median = float64(array[len(array)/2]+array[len(array)/2+1]) / 2
	}
	return median
}

func calculateMode(array []int) int {
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	// fmt.Println("Sorted Array: ", array)

	mode := array[0]
	highFreq := 1
	changeIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			if i-changeIndex > highFreq {
				mode = array[i-1]
				// fmt.Println("New Mode: ", mode)
				highFreq = i - changeIndex
				// fmt.Println("New Frequency: ", highFreq)
			}
			changeIndex = i
		}
	}
	return mode
}

func main() {
	intArray := generateRandomIntArray()
	// fmt.Println("Random Array: ", intArray)
	fmt.Println("Random Array Size: ", len(intArray))

	mean := calculateMean(intArray)
	fmt.Println("Random Array Mean: ", mean)

	median := calculateMedian(intArray)
	fmt.Println("Random Array Median: ", median)

	mode := calculateMode(intArray)
	fmt.Println("Random Array Mode: ", mode)

}

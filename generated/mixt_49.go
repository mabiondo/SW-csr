// BEGIN
package main

import (
"fmt"
)

func countSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	maxVal := arr[0]
	for _, el := range arr {
		if el > maxVal {
			maxVal = el
		}
	}

	freq := make([]int, maxVal+1)
	sortedArr := make([]int, len(arr))

	// Determine frequencies of each element
	for _, el := range arr {
		freq[el]++
	}

	// Construct cumulative distribution table
	runningTotal := 0
	for idx := 0; idx < len(freq); idx++ {
		prevRunningTotal := runningTotal
		runningTotal += freq[idx]
		freq[idx] = prevRunningTotal
	}

	// Populate the sorted array according to the cumulative distribution table
	for _, el := range arr {
		sortedArr[freq[el]] = el
		freq[el]++
	}

	return sortedArr
}

func main() {
	arr := []int{5, 3, 7, 2, 8, 4, 9, 1, 6}
	fmt.Println("Original Array:", arr)
	fmt.Println("Sorted Array Using Count Sort:", countSort(arr))
}
// END
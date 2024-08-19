//BEGIN
package main

import (
"fmt"
)

func countSort(arr []int) []int {
	max := getMax(arr)
	count := make([]int, max+1)
	for _, num := range arr {
		count[num]++
	}
	sortedIndex := 0
	for i, freq := range count {
		for freq > 0 {
			arr[sortedIndex] = i
			sortedIndex++
			freq--
		}
	}
return arr
}

func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	arr := []int{3, 2, 5, 100, 2, 4, 2, 3}
	fmt.Println(countSort(arr))
}
//END
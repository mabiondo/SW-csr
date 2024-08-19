// BEGIN
package main

import "sort"

func oddSwaps(arr []int) []int {
	sort.Ints(arr)
	for i := 1; i < len(arr); i += 2 {
		if i+1 < len(arr) {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return arr
}

//END

//BEGIN
package main

import (
"fmt"
)

func areElementsUnique(arr []int) bool {
	elemSet := make(map[int]bool)
	for _, elem := range arr {
	if elemSet[elem] {
		return false
	}
		elemSet[elem] = true
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 2, 3, 4}
	fmt.Println(areElementsUnique(arr1)) // Expected output: true
	fmt.Println(areElementsUnique(arr2)) // Expected output: false
}
//END
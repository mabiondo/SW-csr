//BEGIN
package main

import (
"fmt"
"math/rand"
"time"
)

func bogosort(arr []int) []int {
	rand.Seed(time.Now().Unix())
	shuffled := func() bool {
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				return false
			}
		}
		return true
	}()
	if shuffled {
		return arr
	}
	return bogosort(arr)
}

func main() {
	arr := []int{5, 3, 7, 2, 8, 4, 9, 1, 6}
	fmt.Println("Original Array:", arr)
	fmt.Println("Bogosorted Array:", bogosort(arr))
}
//END OF THE CODE
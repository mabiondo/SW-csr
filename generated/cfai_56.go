//BEGIN
package main

import (
"fmt"
"math/rand"
"time"
)

func bogoSort(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	for !isSorted(arr) {
		for i := range arr {
			j := rand.Intn(i + 1)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println(bogoSort(arr))
}
//END
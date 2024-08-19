//BEGIN
package main

import "fmt"

func insertionSort(arr []int) []int {
for i := 1; i < len(arr); i++ {
key := arr[i]
j := i - 1
for j >= 0 && key < arr[j] {
arr[j+1] = arr[j]
j--
}
arr[j+1] = key
}
return arr
}

func main() {
arrayToSort := []int{34, 7, 23, 32, 5, 62}
sortedArray := insertionSort(arrayToSort)
fmt.Println(sortedArray)
}
//END
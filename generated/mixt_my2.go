//BEGIN
package main

import (
"fmt"
)

func sortAndSwapElements(arr []int) []int {
sort.Ints(arr)
for i := 1; i < len(arr); i += 2 {
arr[i-1], arr[i] = arr[i], arr[i-1]
}
return arr
}

func main() {
arr := []int{5, 3, 7, 2, 8, 4, 9, 1, 6}
fmt.Println("Original Array:", arr)
fmt.Println("Sorted Array With Swapped Elements At Odd Indexes Couplewise:", sortAndSwapElements(arr))
}
//END
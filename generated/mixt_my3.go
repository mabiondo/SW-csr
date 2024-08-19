// BEGIN
package main

import (
"fmt"
)

func AreArrayElementsUnique(arr []int) bool {
m := map[int]bool{}
for _, el := range arr {
if m[el] {
return false
}
m[el] = true
}
return true
}

func main() {
arr := []int{1, 2, 3, 4, 5}
fmt.Println("Array:", arr)
fmt.Println("Are All Elements Unique?:", AreArrayElementsUnique(arr))

repeatedArr := []int{1, 2, 3, 4, 4}
fmt.Println("\nRepeated Array:", repeatedArr)
fmt.Println("Are All Elements Unique?:", AreArrayElementsUnique(repeatedArr))
}
// END
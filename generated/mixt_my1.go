
//BEGIN
package main

import (
"fmt"
)

func CheckElementAppearanceNumber(arr []int, element, num int) bool {
counter := 0
for _ , el := range arr {
if el == element {
counter++
if counter == num {
return true
}
}
}
return false
}

func main() {
arr := []int{1, 2, 3, 2, 4, 2, 5, 2, 6, 2, 7, 2}
element := 2
num := 4
fmt.Println("Array:", arr)
fmt.Println("Element:", element)
fmt.Println("Target Number:", num)
fmt.Println("CheckElementAppearanceNumber:", CheckElementAppearanceNumber(arr, element, num))
}
//END
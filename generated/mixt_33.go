//BEGIN
package main

import "fmt"

func fillArray(value int, length int) []int {
arr := make([]int, length)
for i := 0; i < length; i++ {
arr[i] = value
}
return arr
}

func main() {
filledArr := fillArray(5, 10)
fmt.Println(filledArr)
}
//END
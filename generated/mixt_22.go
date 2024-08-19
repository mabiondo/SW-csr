//BEGIN
package main

import "fmt"

func evenOddSums(x []int) (evenSum, oddSum int) {
for i, value := range x {
if i % 2 == 0 {
evenSum += value
} else {
oddSum += value
}
}
return
}

func main() {
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
e, o := evenOddSums(numbers)
fmt.Printf("Even indexed sum: %d\n", e)
fmt.Printf("Odd indexed sum: %d\n", o)
}
//END

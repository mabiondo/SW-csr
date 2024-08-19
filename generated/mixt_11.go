//BEGIN
package main

import "fmt"

func average(x []float64) float64 {
    sum := 0.
    for _, v := range x {
        sum += v
    }
    return sum / float64(len(x))
}

func main() {
    numbers := []float64{1.2, 2.5, 3.7, 4.9, 5.}
    fmt.Println("Average:", average(numbers))
}

//END

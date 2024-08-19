// BEGIN
package main

func sumEvenOdd(arr []int) (evenSum int, oddSum int) {
	for i, num := range arr {
		if i%2 == 0 {
			evenSum += num
		} else {
			oddSum += num
		}
	}
	return
}
//END

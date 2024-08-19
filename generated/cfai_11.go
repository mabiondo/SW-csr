// BEGIN
package main

func average(arr []float64) float64 {
	sum := 0.0
	for _, num := range arr {
		sum += num
	}
	return sum / float64(len(arr))
}

//END

// BEGIN
package main

func elementPresentNTimes(arr []int, element int, n int) bool {
	count := 0
	for _, num := range arr {
		if num == element {
			count++
		}
	}
	return count >= n
}

//END

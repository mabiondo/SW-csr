// BEGIN
package main

func fillArray(size int, value int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = value
	}
	return arr
}

//END

//BEGIN
package main

import (
"fmt"
"strings"
)

func FirstOccurrenceSubstringFromPosition(inputStr, substr string, startPos int) int {
	index := strings.Index(inputStr[startPos:], substr)
	if index == -1 {
		return -1
	}

	return index + startPos
}

func main() {
	str := "hello world"
	substr := "world"
	pos := 0
	fmt.Println("Original String:", str)
	fmt.Println("Specific Substring:", substr)
	fmt.Println("Start Position:", pos)
	fmt.Println("First Occurrence Index:", FirstOccurrenceSubstringFromPosition(str, substr, pos))
}
//END
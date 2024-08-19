//BEGIN
package main

import (
"fmt"
)

func RemoveOnlyOnceCharacters(inputStr string) string {
	charCountMap := make(map[rune]int)
	uniqueChars := ""

	for _, c := range inputStr {
		if charCountMap[c] == 0 {
			charCountMap[c]++
		} else if charCountMap[c] == 1 {
			uniqueChars += string(c)
		}
	}

	for _, c := range uniqueChars {
		if charCountMap[rune(c)]&^1 == 0 {
			inputStr = strings.ReplaceAll(inputStr, string(c), "")
		}
	}

	return inputStr
}

func main() {
	str := "hello world"
	fmt.Println("Original String:", str)
	fmt.Println("Resultant String:", RemoveOnlyOnceCharacters(str))
}
//END
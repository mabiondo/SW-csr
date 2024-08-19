// BEGIN
package main

func removeUniqueChars(s string) string {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	result := make([]rune, 0)
	for char, count := range charCount {
		if count > 1 {
			result = append(result, char)
		}
	}
	return string(result)
}

//END

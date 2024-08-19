// BEGIN
package main

func findSubstring(s string, sub string, startPos int) int {
	subLen := len(sub)
	sLen := len(s)
	endPos := sLen - len(sub)

	for i := startPos; i <= endPos; i++ {
		if s[i:i+subLen] == sub {
			return i
		}
	}

	return -1
}

//END

//BEGIN
package main

import (
"fmt"
)

func StringLength(str string) int {
return len(str)
}

func main() {
str := "abc"
fmt.Println("String Length:", StringLength(str))
}
//END
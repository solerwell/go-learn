package main

import (
	"C"
	"fmt"
	"os"
	"unicode"
)

//export ContainsChinese
func ContainsChinese(str string) bool {
	if len(str) == 0 {
		os.Exit(1)
	}
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("string is %s\n", os.Args[1])
	fmt.Println("contains han char: ", ContainsChinese(os.Args[1]))
}

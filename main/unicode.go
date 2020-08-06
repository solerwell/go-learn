package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Jimmy: hello,李明博，이명박!"
	fmt.Println(str)

	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			fmt.Print(string(r))
		}
	}
}

package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	reversedRunes := make([]rune, len(runes))

	initIndex := len(runes) - 1

	for i := initIndex; i >= 0; i-- {
		reversedRunes[initIndex-i] = runes[i]
	}

	return string(reversedRunes)
}

func main() {
	s := "главр😀ыба"

	fmt.Println(reverseString(s))
}

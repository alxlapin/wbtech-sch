package main

import (
	"fmt"
	"unicode"
)

func checkUniqueness(s string) bool {
	m := map[rune]int{}

	for _, r := range s {
		if m[unicode.ToLower(r)] > 0 {
			return false
		}

		m[r]++
	}

	return true
}

func main() {
	var input string
	fmt.Println("Введите строку для анализа")
	fmt.Print()
	fmt.Scan(&input)

	if checkUniqueness(input) {
		fmt.Println("Строка соответствует требованиям")
	} else {
		fmt.Println("Строка НЕ соответствует требованиям")
	}
}

package main

import (
	"fmt"
	"strings"
)

func reverseSentence(s string) string {
	parts := strings.Split(s, " ")

	i := 0
	j := len(parts) - 1

	for i < j {
		parts[i], parts[j] = parts[j], parts[i]
		i++
		j--
	}

	return strings.Join(parts, " ")
}

func main() {
	fmt.Println(reverseSentence("snow dog sun"))
	fmt.Println(reverseSentence("snow dog sun beach"))
}

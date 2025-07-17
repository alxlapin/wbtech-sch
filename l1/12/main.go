package main

import "fmt"

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	inputMap := map[string]bool{}

	for _, str := range input {
		inputMap[str] = true
	}

	result := []string{}
	for key := range inputMap {
		result = append(result, key)
	}

	fmt.Println(result)
}

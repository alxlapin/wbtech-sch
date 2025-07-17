package main

import (
	"fmt"
	"reflect"
)

func getVarType(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}

		return "unknown"
	}
}

func main() {
	i := 1
	j := 1.1
	s := "Golang"
	b := true
	ch := make(chan int)

	fmt.Printf("%v is type of %s\n", i, getVarType(i))
	fmt.Printf("%v is type of %s\n", j, getVarType(j))
	fmt.Printf("%v is type of %s\n", s, getVarType(s))
	fmt.Printf("%v is type of %s\n", b, getVarType(b))
	fmt.Printf("%v is type of %s\n", ch, getVarType(ch))
}

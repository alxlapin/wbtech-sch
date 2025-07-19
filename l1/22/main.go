package main

import (
	"fmt"
	"math/big"
)

func main() {
	var op1, op2, operator string
	fmt.Print("Enter operand 1: ")
	fmt.Scan(&op1)

	fmt.Print("Enter operand 2: ")
	fmt.Scan(&op2)

	fmt.Println("Select operator:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Print()
	fmt.Scan(&operator)

	bigInt_1, _ := new(big.Int).SetString(op1, 10)
	bigInt_2, _ := new(big.Int).SetString(op2, 10)

	result := new(big.Int)

	operatorCodes := map[string]string{
		"1": "+",
		"2": "-",
		"3": "*",
		"4": "/",
	}

	switch operator {
	case "1":
		result.Add(bigInt_1, bigInt_2)
	case "2":
		result.Sub(bigInt_1, bigInt_2)
	case "3":
		result.Mul(bigInt_1, bigInt_2)
	case "4":
		result.Div(bigInt_1, bigInt_2)
	}

	fmt.Printf("%s %s %s = %s\n", bigInt_1.String(), operatorCodes[operator], bigInt_2.String(), result.String())
}

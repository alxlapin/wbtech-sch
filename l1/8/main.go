package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

//Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
//
//Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
//
//Подсказка: используйте битовые операции (|, &^).

func main() {
	i, _ := strconv.ParseInt(os.Args[1], 10, 64)

	bitNum, _ := strconv.ParseInt(os.Args[2], 10, 0)

	fmt.Printf("Input = %d (%b), bitNum = %d \n", i, i, bitNum)

	bitNumDec := int64(math.Pow(2, float64(bitNum-1)))

	j := i | bitNumDec
	if j == i {
		j = i &^ bitNumDec
	}

	fmt.Printf("Output = %d (%b)\n", j, j)
}

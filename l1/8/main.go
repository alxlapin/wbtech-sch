package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

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

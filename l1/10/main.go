package main

import (
	"fmt"
	"math"
)

func main() {
	data := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	step := 10.0

	groups := make(map[int][]float64)

	for _, v := range data {
		group := int(math.Floor(math.Abs(v)/step) * 10)

		if v < 0 {
			group = group * -1
		}

		groups[group] = append(groups[group], v)
	}

	fmt.Println(groups)
}

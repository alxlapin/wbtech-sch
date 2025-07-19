package main

import (
	"fmt"
	"wbtech-sch-l1/24/domain"
)

func main() {
	var p1X, p1Y, p2X, p2Y float64
	fmt.Println("Введите координаты первой точки")
	fmt.Print("x=")
	fmt.Scan(&p1X)
	fmt.Print("y=")
	fmt.Scan(&p1Y)

	fmt.Println("Введите координаты второй точки")
	fmt.Print("x=")
	fmt.Scan(&p2X)
	fmt.Print("y=")
	fmt.Scan(&p2Y)

	p1 := domain.NewPoint(p1X, p1Y)
	p2 := domain.NewPoint(p2X, p2Y)

	fmt.Println(p1.Distance(p2))
}

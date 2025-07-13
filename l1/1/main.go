package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) GetInfo() string {
	return fmt.Sprintf("Human name is %s, age is %d", h.Name, h.Age)
}

type Action struct {
	Human
	ID string
}

func main() {
	action := Action{
		Human: Human{Name: "Alexei Lapin", Age: 30},
		ID:    "get-info",
	}

	fmt.Println("Name and age:", action.Name, action.Age)
	fmt.Println("Action.ID:", action.ID)
	fmt.Println("Action.GetInfo():", action.GetInfo())
}

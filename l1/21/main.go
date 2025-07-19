package main

import (
	"fmt"
)

type EuroSocket interface {
	ConnectToEuroSocket()
}

type EuroPlug struct {
}

func (ep EuroPlug) ConnectToEuroSocket() {
	fmt.Println("Connected euro plug to euro socket")
}

type AmericanPlug struct {
}

type AmericanPlugToEuroSocketAdapter struct {
	ap AmericanPlug
}

func (ap AmericanPlugToEuroSocketAdapter) ConnectToEuroSocket() {
	fmt.Println("Connected american plug to euro socket adapter -> Connected adapter to euro socket")
}

func connectPlugToEuroSocket(es EuroSocket) {
	es.ConnectToEuroSocket()
}

func main() {
	// Пример примитивный, отображает решение бытовой задачи с адаптацией различных типов вилок под евро розетку.

	ep := EuroPlug{}

	fmt.Println("-----Euro plug -> euro socket")
	connectPlugToEuroSocket(ep)
	fmt.Println("-----Euro plug -> euro socket")

	fmt.Println()

	adapter := AmericanPlugToEuroSocketAdapter{
		ap: AmericanPlug{},
	}

	fmt.Println("-----American plug -> euro socket")
	connectPlugToEuroSocket(adapter)
	fmt.Println("-----American plug -> euro socket")

	// Паттерн адаптер применим, например, в следующих сценариях:
	// 1. если нет возможности изменить/расширить код сторонней библиотеки, но требуется адаптировать его под конкретные нужды;
	// 2. если требуется адаптировать старый формат данных для новой версии API и т.д.

	// Из минусов, как и по многим другим паттернам, можно выделить усложнение логики
	// из-за "раздувания" количества классов/сущностей в системе.
}

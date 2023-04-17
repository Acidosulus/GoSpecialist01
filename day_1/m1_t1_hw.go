/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

const iCost = 48

var iDistance, iConsumption float64

func main() {

	fmt.Print("Длина пути (50 - 10000 км): ")
	fmt.Scan(&iDistance)
	if !(50 <= iDistance && iDistance <= 10000) {
		fmt.Println("Длина пути вне интервала 50 - 10000 км")
		return
	}
	fmt.Println("")

	fmt.Print("Расход в литрах (5-25 литров на 100 км): ")
	fmt.Scan(&iConsumption)
	if !(5 <= iConsumption && iConsumption <= 25) {
		fmt.Println("Расход в литрах вне интервала (5-25 литров на км)")
		return
	}
	fmt.Println("")

	fmt.Println("Стоимость бензина: ", iCost)
	fmt.Println("")

	fmt.Println("=========================================")
	fmt.Println("")
	fmt.Print("Стоимость поездки: ", iDistance*iConsumption*iCost/100, " рублей")
}

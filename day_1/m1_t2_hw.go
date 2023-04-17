/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import (
	"fmt"
	"strconv"
)

var cNumber string
var iNumber int

func main() {

	fmt.Print("Введите трёхзначное число: ")
	fmt.Scan(&iNumber)

	if iNumber < 100 || iNumber > 999 {
		fmt.Print("Введённое число не является трёхзначным")
		return
	}
	fmt.Println()

	cNumber = strconv.Itoa(iNumber)
	fmt.Println("Реверсная запись числа: " + string(cNumber[2]) + string(cNumber[1]) + string(cNumber[0]))
}

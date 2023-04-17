/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import (
	"fmt"
)

var iNumber, iFirstDigit, iSecondDigit, iThirdDigit, iFourthDigit int

func main() {

	fmt.Print("Введите четырёзначное: ")
	fmt.Scan(&iNumber)
	fmt.Println()
	if iNumber < 1000 || iNumber > 9999 {
		fmt.Print("Введённое число не является четырёхзначным")
		return
	}

	iFirstDigit = iNumber % 10
	println("1й разряд числа: ", iFirstDigit)

	iSecondDigit = ((iNumber % 100) - iFirstDigit) / 10
	println("2й разряд числа: ", iSecondDigit)

	iThirdDigit = ((iNumber % 1000) - iSecondDigit*10 - iFirstDigit) / 100
	println("3й разряд числа: ", iThirdDigit)

	iFourthDigit = ((iNumber % 10000) - iThirdDigit*100 - iSecondDigit*10 - iFirstDigit) / 1000
	println("4й разряд числа: ", iFourthDigit)

	fmt.Println()
	if (iFirstDigit == iFourthDigit) && (iSecondDigit == iThirdDigit) {
		fmt.Println("Введённое число - Палиндром")
	} else {
		fmt.Println("Введённое число - НЕ Палиндром")
	}
}

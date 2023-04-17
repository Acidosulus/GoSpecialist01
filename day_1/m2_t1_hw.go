/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/
package main

import (
	"fmt"
)

var iFirst, iSecond, iThird int

func main() {

	fmt.Println("Введите три числа: ")
	fmt.Scan(&iFirst, &iSecond, &iThird)
	fmt.Println()

	if iFirst <= iSecond && iFirst <= iThird {
		if iSecond <= iThird {
			fmt.Println(iFirst, iSecond, iThird)
		} else {
			fmt.Println(iFirst, iThird, iSecond)
		}
	}

	if iSecond <= iFirst && iSecond <= iThird {
		if iFirst <= iThird {
			fmt.Println(iSecond, iFirst, iThird)
		} else {
			fmt.Println(iSecond, iThird, iFirst)
		}
	}

	if iThird <= iSecond && iThird <= iFirst {
		if iSecond <= iFirst {
			fmt.Println(iThird, iSecond, iFirst)
		} else {
			fmt.Println(iThird, iFirst, iSecond)
		}
	}

}

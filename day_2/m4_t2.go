/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/

package main

import "fmt"

func bottles(iCount int) (string, bool) {
	if iCount < 0 || iCount > 200 {
		return fmt.Sprintf("Количество бутылок %v за пределами диапазона от 0 до 200", iCount), false
	}

	var iRemain int
	var cBottlesEnding string
	/*
		0,5-20,25-30,100,200,11,12,13,14 бутылок
		1,21 бутылка
		2-4,22-24 бутылки
	*/
	iRemain = iCount % 10

	if iCount%100 >= 10 && iCount%100 <= 14 {
		iRemain = 0
	}

	switch iRemain {
	case 0, 5, 6, 7, 8, 9:
		cBottlesEnding = "ок"
	case 1:
		cBottlesEnding = "ка"
	case 2, 3, 4:
		cBottlesEnding = "ки"
	}

	return fmt.Sprintf("%v бутыл%v", iCount, cBottlesEnding), true
}

func main() {
	var sResult string

	for i := 0; i <= 201; i++ {

		sResult, _ = bottles(i)
		fmt.Println(sResult)

	}
}

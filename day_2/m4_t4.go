/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import "fmt"

func DrawChessBoard(nSideLength int) {
	for i := 1; i <= nSideLength; i++ {
		for j := 1; j <= nSideLength; j++ {
			fmt.Print((j+(i%2))%2, " ")
		}
		fmt.Println()
	}
}

var nSideLength int

func main() {
	fmt.Println()
	fmt.Print("Размер шахматной доски (0 - 20): ")
	fmt.Scan(&nSideLength)
	if !(0 <= nSideLength && nSideLength <= 20) {
		fmt.Println("Размер шахматной доски вне интервала (0 - 20)")
		return
	}

	fmt.Println()

	DrawChessBoard(nSideLength)

	fmt.Println()
}

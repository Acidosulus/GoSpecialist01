/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package main

import (
	"fmt"
	"math"
)

// возвращает длину отрезка ограниченного точками с заданными координатами
func SegmentLength(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// проверяет треугольник на возможность существования на плоскоскти
func MayBeSoTrianle(x1, y1, x2, y2, x3, y3 float64) bool {
	var fSideLength1, fSideLength2, fSideLength3 float64
	fSideLength1 = SegmentLength(x1, y1, x2, y2)
	fSideLength3 = SegmentLength(x1, y1, x3, y3)
	fSideLength2 = SegmentLength(x2, y2, x3, y3)
	return ((fSideLength1+fSideLength2 > fSideLength3) &&
		(fSideLength2+fSideLength3 > fSideLength1) &&
		(fSideLength1+fSideLength3 > fSideLength2))
}

// вычисляет прощадь треугольника
func TriangleSquare(x1, y1, x2, y2, x3, y3 float64) float64 {
	return math.Abs((x2-x1)*(y3-y1)-(x3-x1)*(y2-y1)) / 2
}

// проверяет треугольник на прямоугольность
func IsRectangularTriangle(x1, y1, x2, y2, x3, y3 float64) bool {
	var fSideLength1, fSideLength2, fSideLength3 float64
	fSideLength1 = SegmentLength(x1, y1, x2, y2)
	fSideLength3 = SegmentLength(x1, y1, x3, y3)
	fSideLength2 = SegmentLength(x2, y2, x3, y3)

	if fSideLength1 >= fSideLength2 && fSideLength1 >= fSideLength3 {
		return math.Round(math.Pow(fSideLength1, 2)*1000000)/1000000 == math.Round((math.Pow(fSideLength2, 2)+math.Pow(fSideLength3, 2))*1000000)/1000000
	}
	if fSideLength2 >= fSideLength1 && fSideLength2 >= fSideLength3 {
		return math.Round(math.Pow(fSideLength2, 2)*1000000)/1000000 == math.Round((math.Pow(fSideLength1, 2)+math.Pow(fSideLength3, 2))*1000000)/1000000
	}
	if fSideLength3 >= fSideLength2 && fSideLength3 >= fSideLength1 {
		return math.Round(math.Pow(fSideLength3, 2)*1000000)/1000000 == math.Round((math.Pow(fSideLength2, 2)+math.Pow(fSideLength1, 2))*1000000)/1000000
	}
	return false
}

// вызывает функции проверок, выводит результаты в консоль
func RunThreeFuncs(x1, y1, x2, y2, x3, y3 float64) {
	fmt.Printf("Точки (%v,%v) (%v,%v) (%v,%v)\n", x1, y1, x2, y2, x3, y3)
	bMayBe := MayBeSoTrianle(x1, y1, x2, y2, x3, y3)
	if bMayBe {
		fmt.Println("Треугольник с такими точками существовать может")
		fmt.Printf("Площадь треугольника %v\n", TriangleSquare(x1, y1, x2, y2, x3, y3))
		if IsRectangularTriangle(x1, y1, x2, y2, x3, y3) {
			fmt.Println("Это прямоугольный треугольник")
		} else {
			fmt.Println("Это НЕ прямоугольный треугольник")
		}
	} else {
		fmt.Println("Треугольник с такими точками существовать НЕ может")
	}
	fmt.Println()
}

func main() {
	fmt.Println()

	RunThreeFuncs(0, 1, 3, 4, 5, 6)
	RunThreeFuncs(2, 1, 8, 6, 7, 1)
	RunThreeFuncs(3, 2, 3, 6, 7, 6)

}

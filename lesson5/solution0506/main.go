package main

import (
	"fmt"
	"math"
)

func main() {
	var a, c int
	var b float64
	fmt.Println("ax^2 + bx + c = 0")
	fmt.Println("Введите 3 числа")
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*(float64(a*c))

	if d > 0 { // если дескрименант больше 0
		x1 := (-b - math.Sqrt(d)) / (2 * float64(a))
		x2 := (-b + math.Sqrt(d)) / (2 * float64(a))
		fmt.Printf("Уровнение имеет 2 корня: х1 = %.2f и х2 = %.2f", x1, x2)
	} else if d == 0 { // если дескрименант равен 0
		x := -b / (2 * float64(a))
		fmt.Printf("Уровнение имеет 1 корень: х = %.2f", x)
	} else {
		fmt.Println("Уровнение не имеет корней")
	}

}

func init() {
	fmt.Println("**********Решение квадратного уравнения*************")
	fmt.Println("Пользователь вводит коэффициенты a, b и c квадратного уравнения." +
		"Программа должна вывести корни уравнения.")
}

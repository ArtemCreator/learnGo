// Сумма двух чисел по единице
package main

import "fmt"

func main() {
	fmt.Println("Введите 2 целых числа")
	var a, b int
	fmt.Scan(&a, &b)

	c := a + b

	for a <= c{
		fmt.Println(a)
		a++
	}
}

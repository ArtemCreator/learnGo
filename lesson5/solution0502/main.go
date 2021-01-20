package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите 3 целых числа")
	fmt.Scan(&a, &b, &c)

	if a > 0 || b > 0 || c > 0 {
		if a > 0 {
			fmt.Printf("Первое число %d положительно", a)
		} else if b > 0 {
			fmt.Printf("Второе число %d положительно", b)
		} else if c > 0 {
			fmt.Printf("Третье число %d положительно", c)
		}
	} else {
		fmt.Println("Положительных чисел не введено")
	}
}

func init() {
	fmt.Println("**************Проверить, что одно из чисел положительное***************")
	fmt.Println("Ввести 3 целых числа, проверить является ли хотя бы одно число положительное\n")
}

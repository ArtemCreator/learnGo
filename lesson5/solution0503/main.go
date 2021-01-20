package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите 3 числа, проверим есть ли совпадающие из них")
	fmt.Scan(&a, &b, &c)

	if a == b || a == c || b == c {
		fmt.Println("Есть совпадения чисел")
	} else {
		fmt.Println("Нету совпадений")
	}

	comparing()
}

func comparing() {
	var a, b, c int
	var count int

	fmt.Println("Введите 3 числа, проверим есть ли совпадающие из них")
	fmt.Scan(&a, &b, &c)

	if a == b || a == c || b == c {
		fmt.Print("Совпадают числа: ")
		if a == b {
			count++
			fmt.Printf("%d и %d; ", a, b)
		}
		if a == c {
			count++
			fmt.Printf("%d и %d; ", a, c)
		}
		if b == c {
			count++
			fmt.Printf("%d и %d; ", b, c)
		}
		fmt.Println(" всего совпадений ", count)
	} else {
		fmt.Println("Нету совпадений")
	}
}

func init() {
	fmt.Println("*************Проверить, что есть совпадающие числа***************")
}

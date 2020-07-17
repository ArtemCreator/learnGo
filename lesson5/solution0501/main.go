package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Введите координаты точки")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0{
		fmt.Print("Точка находится в I четверти")
	}else if x < 0 && y > 0 {
		fmt.Print("Точка находится в II четверти")
	}else if x < 0 && y < 0 {
		fmt.Print("Точка находится в III четверти")
	}else if x > 0 && y < 0 {
		fmt.Print("Точка находится в IV четверти")
	}else{
		fmt.Print("Точка находится на нулевой оси координат")
	}
}

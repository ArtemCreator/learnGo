package main

import "fmt"

func main() {
	//chess(6, 8)
	//triangle1(10)
	//triangle4(10)
	//triangle2(10)
	//triangle3(10)
	//triangle5(10)
	triangle6(10)
	//triangle7(10)
}

// chess выводит на экран шахматную доску
func chess(row, column int) {
	//for i := 0; i < row; i++ {
	//	for j := 0; j < column; j++ {
	//		if i % 2 == 0 {
	//			if j % 2 == 0 {
	//				fmt.Print(" ")
	//			} else{
	//				fmt.Print("*")
	//			}
	//		}
	//		if i % 2 != 0 {
	//			if j % 2 != 0{
	//				fmt.Print(" ")
	//			}else {
	//				fmt.Print("*")
	//			}
	//		}
	//	}
	//	fmt.Println()
	//}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

// 1ый вариант треугольника
func triangle1(row int) {
	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("(")
		}
		fmt.Println()
	}
}

// 2ый вариант треугольника
func triangle2(row int) {
	for i := 0; i <= row; i++ {
		for j := row; j > i; j-- {
			fmt.Print("(")
		}
		fmt.Println()
	}
}

// Обратный вариант 1го треугольника
func triangle3(row int) {
	for i := 0; i <= row; i++ {
		for j := row; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("(")
		}
		fmt.Println()
	}
}

// Обратный вариант 2го треугольника
func triangle4(row int) {
	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := row; j > i; j-- {
			fmt.Print("(")
		}
		fmt.Println()
	}
}

func triangle5(row int) {
	for i := 0; i <= row; i++ {
		for j := row; j > i/2; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

// Равностороний треугольник
func triangle6(row int) {
	for i := 0; i <= row; i++ {
		// левое пустое пространство
		for j := row; j > i; j-- {
			fmt.Print(" ")
		}
		// левая часть треугольника
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		// правая часть треугольника
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		// правое пустое место
		for j := row; j > i; j-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

// Reverse равностороний треугольник
func triangle7(row int) {
	for i := row; i > 0; i-- {
		// левое пустое пространство
		for j := row; j > i; j-- {
			fmt.Print(" ")
		}
		// левая часть треугольника
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		// правая часть треугольника
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		// правое пустое место
		for j := row; j > i; j-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

// Заполнение трех переменных
package main

import "fmt"

func main() {
	var a, b, c int
	for c != 1000 {
		a++; b++; c++
		if a <= 10 {
			fmt.Print("a = ", a)
		}
		if b <= 100 {
			fmt.Print("\tb = ", b)
		}
		fmt.Println("\tc = ", c)
	}
}
package main

import (
	"fmt"
)

func main() {
	var numberTicket int
	fmt.Println("Введите четырехзначный номер")
	fmt.Scan(&numberTicket)

	ticket1 := numberTicket / 100 // первая половина билета
	ticket2 := numberTicket % 100 // вторая половина билета

	n1 := ticket1 / 10
	n2 := ticket1 % 10
	n3 := ticket2 / 10
	n4 := ticket2 % 10

	//fmt.Print(n1, n2, n3, n4)

	if n1 + n2 == n3 + n4{
		if n1 == n4 && n2 == n3{
			fmt.Println("Зеркалный билет")
		}
		fmt.Println("Счастливый билет")
	}else if numberTicket < 1000 || numberTicket > 9999{
		fmt.Println("Введено не четерехзначное число")
	}else {
		fmt.Println("Обычный билет")
	}
}

func init() {
	fmt.Println("*********Счастливый билет**********")
	fmt.Println("Написать программу, в которую пользователь будет вводить четырехзначный номер билета," +
		"а программа выводит, является ли он счастливым (сумма старших двух разрядов равна сумме двух младших" +
		"разрядов) или даже зеркальным.")
}
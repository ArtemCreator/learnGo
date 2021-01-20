package main

import "fmt"

func main() {
	var profit int
	var tax int
	maxRange := 50000
	minRange := 10000

	fmt.Println("Введите размер прибыли")
	fmt.Scan(&profit)

	// начинать надо с большей ставки
	if profit >= 50000 {
		tax = ((profit - maxRange) * 30) / 100
		profit1 := profit - maxRange
		profit -= profit1
		if profit >= 10000 {
			tax1 := ((profit-minRange)*20)/100 + (13*minRange)/100
			tax += tax1
			fmt.Println("Размер налога 30% = ", tax)
		} else {
			tax1 := (13 * minRange) / 100
			tax += tax1
			fmt.Println("Размер налога 30% = ", tax)
		}
	} else if profit >= 10000 {
		tax := ((profit-minRange)*20)/100 + (13*minRange)/100
		fmt.Println("Размер налога 20% = ", tax)
	} else if profit > 0 {
		tax = profit * 13 / 100
		fmt.Println("Размер налога 13% = ", tax)
	} else if profit == 0 {
		fmt.Println("Нету? А если найдем?")
	} else if profit < 0 {
		fmt.Println("Некорректная сумма")
	}
}
func init() {
	fmt.Println("*********Прогрессивный налог***********")
}

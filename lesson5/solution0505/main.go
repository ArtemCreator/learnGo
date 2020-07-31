package main

import "fmt"

func main() {
	var firstD, secondD, thirdD int
	fmt.Println("Введите 3 процентные ставки по депозитам")
	fmt.Scan(&firstD, &secondD, &thirdD)
	isLessZero := firstD < 0 || secondD < 0 || thirdD < 0

	if isLessZero == false{
		if firstD > secondD && firstD > thirdD{
			fmt.Print("Наиболее выгодная ставка: ", firstD)
			if secondD > thirdD{
				fmt.Println("и ", secondD)
			}else {
				fmt.Println("и ", thirdD)
			}
		}else if secondD > firstD && secondD > thirdD{
			fmt.Print("Наиболее выгодная ставка: ", secondD)
			if firstD > thirdD{
				fmt.Println("и ", firstD)
			}else {
				fmt.Println("и ", thirdD)
			}
		}else if thirdD > firstD && thirdD > secondD {
			fmt.Print("Наиболее выгодная ставка: ", thirdD)
			if firstD > secondD {
				fmt.Println(" и ", firstD)
			} else {
				fmt.Println(" и ", secondD)
			}
		}
	}else {
		fmt.Print("Введены ставки меньше нуля")
	}
}

func init() {
	fmt.Println("*****Определение максимальных процентов*****")
	fmt.Print("Определить какие 2 ставки по депозитам являются максимальными" +
		"из предложенных 3\n")
}
package main

import "fmt"

func main() {
	var firstM, secondM, thirdM int
	fmt.Println("Введите номинал монет имеющихся у вас")
	fmt.Scan(&firstM, &secondM, &thirdM)

	var money int
	fmt.Println("Введите сумму которую надо оплатить")
	fmt.Scan(&money)

	if firstM+secondM+thirdM == money {
		fmt.Println("Денежных средств хватает для оплаты без сдачи")
	} else if firstM+secondM == money || secondM+thirdM == money || firstM+secondM == money {
		fmt.Println("Денежных средств хватает для оплаты без сдачи")
	} else if firstM == money || secondM == money || thirdM == money {
		fmt.Println("Денежных средств хватает для оплаты без сдачи")
	} else {
		fmt.Println("Денежных средств не хватить для оплаты без сдачи")
	}
}
func init() {
	fmt.Println("*********Сумма без сдачи**********")
}

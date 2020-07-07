package main

import "fmt"

func main() {
	var day int
	var totalAmount int
	var guests int
	var delivery int
	var service int
	countDelivery := 10000

	fmt.Println("Сумма счета составила ")
	fmt.Scan(&totalAmount)

	fmt.Println("Сколько обслуживалось человек?")
	fmt.Scan(&guests)
	if guests >= 5{
		service = totalAmount / 10
		totalAmount += service
	}

	fmt.Println("Какой сегодня день недели, введите число " +
		"от 1 до 7, где 1 это понедельник, 7 - воскресенье")
	fmt.Scan(&day)

	if day == 1{
		delivery = totalAmount / 10
		totalAmount -= delivery
	}
	if day == 5{
		if totalAmount > countDelivery{
			delivery = (totalAmount / 100) * 5
			totalAmount -= delivery
		}
	}

	fmt.Printf("Обслуживание: %d\nСкидка: %d\nИтоговая сумма: %d",
		service, delivery, totalAmount)
}

func init() {
	fmt.Println("*********Ресторан*********")
}

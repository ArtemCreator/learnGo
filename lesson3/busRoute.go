package main

import (
	"fmt"
)

func main() {
	stantion1 := "Будапештская"
	stantion2 := "ул. Славы 4"
	stantion3 := "Вечный огонь"
	stantion4 := "Аэропорт"
	passagir := 0
	ticketAmount := 20
	var totalSum int

	fmt.Printf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion1, passagir)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&passagir)
	passagirBus := -passagir
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&passagir)
	passagirBus += passagir
	totalSum += (passagir * ticketAmount)
	fmt.Printf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion2, passagirBus)
	fmt.Println("----------Едем-------------")

	fmt.Printf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion2, passagirBus)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&passagir)
	passagirBus -= passagir
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&passagir)
	passagirBus += passagir
	totalSum += (passagir * ticketAmount)
	fmt.Printf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion3, passagirBus)
	fmt.Println("----------Едем-------------")

	fmt.Printf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion3, passagirBus)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&passagir)
	passagirBus -= passagir
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&passagir)
	passagirBus += passagir
	totalSum += (passagir * ticketAmount)
	fmt.Printf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion4, passagirBus)
	fmt.Println("----------Едем-------------")

	fmt.Printf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion4, passagirBus)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&passagir)
	passagirBus -= passagir
	driverSalary := totalSum / 4
	taxes := totalSum / 5
	carRepair := totalSum / 5
	fuelConsumption := totalSum / 5
	income := totalSum - (driverSalary + taxes + carRepair + fuelConsumption)
	fmt.Printf("Всего заработанно %d руб.\n", totalSum)
	fmt.Printf("Зарплата водителя: %d руб.\n", driverSalary)
	fmt.Printf("Расходы на топливо: %d руб.\n", fuelConsumption)
	fmt.Printf("Налоги %d руб.\n", taxes)
	fmt.Printf("Расходы на ремонт машины: %d руб.\n", carRepair)
	fmt.Printf("Итого доход: %d руб.\n", income)

}
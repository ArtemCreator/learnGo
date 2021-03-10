package main

import fmt "fmt"

func main() {
	fmt.Println("Введите название месяца")
	var month string
	_, _ = fmt.Scan(&month)
	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("зима")
	case "март", "апрель", "май":
		fmt.Println("весна")
	case "июнь", "июль", "август":
		fmt.Println("лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("осень")
	default:
		fmt.Println("Не выдумывайте. Нету такого месяца.")
	}
}

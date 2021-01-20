//Заполняем корзины яблоками
package main

import "fmt"

var capacityBasket, capacityBasket2, capacityBasket3 int
var apples int

func main() {
	fmt.Println("Какая вместимость у 3 корзин, введите 3 числа")
	fmt.Scan(&capacityBasket, &capacityBasket2, &capacityBasket3)
	fmt.Println("Сколько яблок надо разложить по корзинам")
	fmt.Scan(&apples)
	var index int

	if capacityBasket > 0 && capacityBasket2 > 0 && capacityBasket3 > 0 && apples > 0 {
		for index = 0; index < capacityBasket; index++ {
			if apples == 0 {
				break
			}
			apples--
		}
		fmt.Printf("В корзине №1 находиться %d яблока, у вас осталось %d яблок\n", index, apples)

		for index = 0; index < capacityBasket2; index++ {
			if apples == 0 {
				break
			}
			apples--
		}
		fmt.Printf("В корзине №2 находиться %d яблока, у вас осталось %d яблок\n", index, apples)

		for index = 0; index < capacityBasket3; index++ {
			if apples == 0 {
				break
			}
			apples--
		}
		fmt.Printf("В корзине №3 находиться %d яблока, у вас осталось %d яблок\n", index, apples)
	} else {
		fmt.Println("Проверьте данные")
	}
}

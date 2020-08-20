//Алгоритм лифта
package main

import "fmt"

const (
	minFloor     = 1  // минимальный этаж
	floors       = 24 // максимальный этаж
	liftCapacity = 2  // вместимость лифта
)

func main() {

	passFloor4 := 3  // пассажиров на 4 этаже
	passFloor7 := 3  // пассажиров на 7 этаже
	passFloor10 := 3 // пассажиров на 10 этаже
	totalPass := passFloor4 + passFloor7 + passFloor10 // общее количество пассажиров

	currentLocation := 1 // текущее положение лифта
	passanger := 0       // колличество людей в лифте
	duration := 1        // направление лифта

	for totalPass != 0 {
		fmt.Printf("Лифт находится на этаже: %d, в нем находится %d пассажир\n", currentLocation, passanger)
		currentLocation += duration                     // лифт едет вверх
		if duration == 1 && currentLocation == floors { // если лифт доехал до верха
			duration = -1 // лифт едет вниз
		} else if currentLocation == minFloor { // если лифт приехал на 1 ый этаж
			duration = 1  // лифт вновь поедет вверх
			passanger = 0 // высадить людей
		} else if duration == -1 && currentLocation == 10 && passFloor10 > 0 && passanger < liftCapacity { // иначе если лифт поехал вниз то
			fmt.Println("10 этаж")
			// если есть люди на этаже ии пасажиров < максВместимости то
			if passFloor10%liftCapacity == 0 { // если людей кратно вместимостиЛифта
				passFloor10 -= liftCapacity // зашли люди
				totalPass -= liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.\n", liftCapacity, passFloor10)
			} else if passFloor10 == 1 {
				passFloor10 = 0
				totalPass--
				passanger++
				fmt.Printf("Забрали с этажа 1 пассажирa. Осталось %d пассажир на этаже.\n", passFloor10)
			} else {
				passFloor10 -= liftCapacity
				totalPass -= liftCapacity
				passanger = liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.", liftCapacity, passFloor10)
			}
		}else if duration == -1 && currentLocation == 7 && passFloor7 > 0 && passanger < liftCapacity { // иначе если лифт поехал вниз то
			fmt.Println("7 этаж")
			if passFloor7 == 1 {
				passFloor7 = 0
				totalPass--
				passanger++
				fmt.Printf("Забрали с этажа 1 пассажирa. Осталось %d пассажир на этаже.\n", passFloor7)
			}else if passFloor7%liftCapacity == 0 { // если людей кратно вместимостиЛифта
				passFloor7 -= liftCapacity // зашли люди
				totalPass -= liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.\n", liftCapacity, passFloor7)
			} else {
				passFloor7 -= liftCapacity
				totalPass -= liftCapacity
				passanger = liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.\n", liftCapacity, passFloor7)
			}
		}else if duration == -1 && currentLocation == 4 && passFloor4 > 0 && passanger < liftCapacity { // иначе если лифт поехал вниз то
			fmt.Println("4 этаж")
			//break
			if passFloor4 == 1 {
				passFloor4 = 0
				totalPass--
				passanger++
				fmt.Printf("Забрали с этажа 1 пассажирa. Осталось %d пассажир на этаже.\n", passFloor4)
			}else if passFloor4%liftCapacity == 0 { // если людей кратно вместимостиЛифта
				passFloor4 -= liftCapacity // зашли люди
				totalPass -= liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.\n", liftCapacity, passFloor4)
			} else {
				passFloor4 -= liftCapacity
				totalPass -= liftCapacity
				passanger = liftCapacity
				fmt.Printf("Забрали с этажа %d пассажиров. Осталось %d пассажир на этаже.\n", liftCapacity, passFloor4)
			}
		}
	}

}

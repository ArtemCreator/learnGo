// Написание простого цикла
package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Введите число до которого будет вестись отсчет")
	var value int
	fmt.Scan(&value)
	index := 0

	for index <= value{
		fmt.Println(index)
		time.Sleep(time.Second)
		index++
	}
}

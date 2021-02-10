package main

import (
	"fmt"
)

func main() {
	ticMinCount := 100000
	ticMaxCount := 999999
	var mirrorCount int

	for i := ticMinCount; i <= ticMaxCount; i++ {

		tic1 := i / 1000 // первая половина билета
		tic2 := i % 1000 // вторая половина билета

		n1 := tic1 / 100
		n2 := (tic1 % 100) / 10
		n3 := (tic1 % 100) % 10
		n4 := tic2 / 100
		n5 := (tic2 % 100) / 10
		n6 := (tic2 % 100) % 10
		if n1 == n6 && n2 == n5 && n3 == n4 {
			mirrorCount++
		}
		//fmt.Println(n1, n2, n3, n4, n5, n6)
	}
	fmt.Printf("Количество зеркальных билетов %d\n", mirrorCount)
	var count int
	mirrorCount = 0
	for i := 100; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if i/100+i/10%10+i%10 ==
				j/100+j/10%10+j%10 {
				count++
			}
			if i/100 == j%10 && i/10%10 == j/10%10 && i%10 == j/100 {
				mirrorCount++
			}

		}
	}
	fmt.Printf("Количество счастливых билетов: %d,\nколичество зеркальных билетов %d\n", count, mirrorCount)

}

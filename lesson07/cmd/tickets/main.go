package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	countTickets := mirrorTicket(999)
	fmt.Printf("Total number of mirror tikets: %v\n", countTickets)

	ticMinCount := 100000
	ticMaxCount := 999999
	var minHappyTicket = 100000
	var maxHappyTicket = 100000
	var minusTic int // разница
	var maxBuyTicket int
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
		if n1+n2+n3 == n4+n5+n6 {
			maxHappyTicket = i
			minusTic = maxHappyTicket - minHappyTicket
			minHappyTicket = maxHappyTicket
			if minusTic > maxBuyTicket {
				maxBuyTicket = minusTic
			}
			//fmt.Println(maxHappyTicket, minusTic)
		}
		//fmt.Println(n1, n2, n3, n4, n5, n6)
	}
	log.Printf("Max count a buy tickets: %v", maxBuyTicket)
	log.Printf("Количество зеркальных билетов %d", mirrorCount)

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

// mirrorTicket returns the number of mirror tickets and an error
func mirrorTicket(countTickets int) int {
	var mirrorCountTickets int
	for i := countTickets; i >= 100; i-- {
		strNumber := strconv.Itoa(i)
		reverseStrNumber := ""
		for length := len(strNumber); length > 0; length-- {
			reverseStrNumber += string(strNumber[length-1])
		}
		reversNum, err := strconv.Atoi(reverseStrNumber)
		if err != nil {
			log.Println("Failure to cast String to int")
		}
		if i/100 == reversNum%10 && i/10%10 == reversNum/10%10 && i%10 == reversNum/100 {
			mirrorCountTickets++
		}
		// fmt.Println(i, reversNum, mirrorCountTickets)
	}
	return mirrorCountTickets
}

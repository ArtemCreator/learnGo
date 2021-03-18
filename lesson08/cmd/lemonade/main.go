package main

import "fmt"

func main() {
	var bills = []int{5, 5, 5, 20, 10, 5}
	//fmt.Print(bills)
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, val := range bills {
		switch {
		case val%5 != 0, val < 0, val > 20:
			return false
		case val == 5:
			five++
		case val == 10:
			five--
			ten++
		case val == 20:
			switch {
			case ten > 0:
				ten--
				five--
			default:
				five -= 3
			}
		}

	}
	if five < 0 || ten < 0 {
		return false
	}
	if bills == nil {
		return false
	}
	return true
}

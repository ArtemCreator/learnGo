package main

import "fmt"

func main() {
	var bills = []int{5, 10, 5, 20}
	//fmt.Print(bills)
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, val := range bills {
		switch val {
		case 5:
			five++
		case 10:
			five--
			ten++
		case 20:
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if five < 0 || ten < 0 {
			return false
		}
	}
	if bills == nil {
		return false
	}
	return true
}

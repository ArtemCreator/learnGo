package main

import "fmt"

func main() {
	fmt.Println("Введите день недели в сокращенной форме\n (пн, вт, ср, чт, пт)")
	var dayweek string
	_, _ = fmt.Scan(&dayweek)
	switch dayweek {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("This word not found")
	}
}

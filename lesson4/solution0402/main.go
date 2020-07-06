package main

import "fmt"

func main() {
	var val, val2, val3 int
	fmt.Println("Введите 3 целых числа")
	fmt.Scan(&val, &val2, &val3)

	//fmt.Print(val, val2, val3)

	if val > 5{
	fmt.Printf("Значение %d больше 5\n", val)
	}
	if val2 > 5{
		fmt.Printf("Значение %d больше 5\n", val2)
	}
	if val3 > 5{
		fmt.Printf("Значение %d больше 5\n", val3)
	}
}

func init(){
	fmt.Println("****Три числа, какое больше 5****")
}
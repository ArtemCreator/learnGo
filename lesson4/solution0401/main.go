package main

import "fmt"

func main() {
	var examScores, examScores2, examScores3  int
	var totalScores int
	const passingScore int = 275

	fmt.Println("Результат первого экзамена")
	fmt.Scan(&examScores)
	fmt.Println("Результат второго экзамена")
	fmt.Scan(&examScores2)
	fmt.Println("Результат первого экзамена")
	fmt.Scan(&examScores3)

	totalScores = examScores + examScores2 + examScores3

	if totalScores >= passingScore {
		fmt.Printf("Вы поступили. Количество набранных баллов %d\n", totalScores)
	}else {
		fmt.Printf("Вы не поступили. Количество набранных баллов %d\n", totalScores)
	}
}

func init() {
	fmt.Println("****Вступительные баллы****")
}

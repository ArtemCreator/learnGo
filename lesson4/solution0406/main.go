package main

import "fmt"

func main() {
	var numberStudents int
	var countGroup int
	var serialNumStudent int
	var numberGroup int

	fmt.Println("Cколько студентов числится на курсе ")
	fmt.Scan(&numberStudents)
	fmt.Println("На сколько групп надо разбить учащихся ")
	fmt.Scan(&countGroup)

	fmt.Println("Сколько человек будет в каждой группе")
	studentsCountGroup := numberStudents / countGroup
	fmt.Println(studentsCountGroup)

	fmt.Println("Порядковый номер студента")
	fmt.Scan(&serialNumStudent)

	if serialNumStudent % studentsCountGroup == 0{
		numberGroup = serialNumStudent / studentsCountGroup
	}else {
		numberGroup  = (serialNumStudent / studentsCountGroup) + 1
		if numberGroup > countGroup{
			numberGroup = countGroup
		}
	}
	fmt.Printf("Студент порядковый № %d зачислен в группу №%d\n", serialNumStudent, numberGroup)
}

func init() {
	fmt.Println("***********Стденты**********")
	fmt.Println("Задача разделить весь курс, состоящий из N студентов, на K групп.\n" +
		"Написать программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента,\n" +
		"а программа определяет, в какую группу он попадает.\n\n" +
		"Подсказка: в одну группу могут попадать студенты из разных частей списка\n345.")
}

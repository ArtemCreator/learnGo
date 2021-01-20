package main

import "fmt"

func main() {
	fmt.Println("Давай сыграем в игру. Ты загадаешь число от 1 до 10, а я попробую угадать его за 4 попытки.")

	var count int
	minValue := 1
	maxValue := 10
	isValueFound := false
	fmt.Println("Загадывай...")
	value := (minValue + maxValue) / 2

	fmt.Printf("Ты загадал число %d. Если Загаданное число больше напиши >, если меньше то напиши <,"+
		" а если я угадал то напиши =\n", value)

	var answerUser string
	fmt.Scan(&answerUser)

	if answerUser == ">" {
		count++
		minValue = value + 1
	} else if answerUser == "<" {
		count++
		maxValue = value - 1
	} else if answerUser == "=" {
		count++
		isValueFound = true
		fmt.Printf("Ответ: %d. Я угадал, мне потребовалось %d попытки", value, count)
	} else {
		count++
		fmt.Println("Ты ответил не корректно. Попробуй еще раз.")
	}

	if isValueFound == false {
		value := (minValue + maxValue) / 2
		fmt.Printf("Ты загадал число %d. Если Загаданное число больше напиши >, если меньше то напиши <,"+
			" а если я угадал то напиши =\n", value)
		fmt.Scan(&answerUser)
		if answerUser == ">" {
			count++
			minValue = value + 1
		} else if answerUser == "<" {
			count++
			maxValue = value - 1
		} else if answerUser == "=" {
			count++
			isValueFound = true
			fmt.Printf("Ответ: %d. Я угадал, мне потребовалось %d попытки", value, count)
		} else {
			count++
			fmt.Println("Ты ответил не корректно. Попробуй еще раз.")
		}
	}

	if isValueFound == false {
		value := (minValue + maxValue) / 2
		fmt.Printf("Ты загадал число %d. Если Загаданное число больше напиши >, если меньше то напиши <,"+
			" а если я угадал то напиши =\n", value)
		fmt.Scan(&answerUser)
		if answerUser == ">" {
			count++
			minValue = value + 1
		} else if answerUser == "<" {
			count++
			maxValue = value - 1
		} else if answerUser == "=" {
			count++
			isValueFound = true
			fmt.Printf("Ответ: %d. Я угадал, мне потребовалось %d попытки", value, count)
		} else {
			count++
			fmt.Println("Ты ответил не корректно. Попробуй еще раз.")
		}
	}

	if isValueFound == false {
		value := (minValue + maxValue) / 2
		fmt.Printf("Ты загадал число %d. Если Загаданное число больше напиши >, если меньше то напиши <,"+
			" а если я угадал то напиши =\n", value)
		fmt.Scan(&answerUser)
		if answerUser == ">" {
			count++
			minValue = value + 1
		} else if answerUser == "<" {
			count++
			maxValue = value - 1
		} else if answerUser == "=" {
			count++
			isValueFound = true
			fmt.Printf("Ответ: %d. Я угадал, мне потребовалось %d попытки", value, count)
		} else {
			count++
			fmt.Println("Ты ответил не корректно. Это была последняя попытка.")
		}

		if count == 4 {
			fmt.Println("Число не отгадано.")
			fmt.Println("ТЫ ПОБЕДИЛ")
		}
	}
}

func init() {
	fmt.Println("*****Игра “Угадывание числа” (дополнительно)*****")
	fmt.Println()
}

// Расчет суммы скидки
package main

import "fmt"

const maxDiscount = 30
const maxAmountDiscount = 2000

func main() {
	//discountFor()
	discount()
}

func discountFor(){
	var sum, discount int
	var amountDiscount float64

	for {
		fmt.Println("Введите сумму товара и скидку")
		fmt.Scan(&sum, &discount)
		if sum > 0 && discount > 0{
			amountDiscount = (float64(sum) / 100) * float64(discount)
			if discount < maxDiscount || amountDiscount < maxAmountDiscount {
				fmt.Println("Ваша скидка составит ", amountDiscount)
			} else if amountDiscount > float64(maxAmountDiscount){
				fmt.Println("Ваша скидка составит ", maxAmountDiscount)
			} else if discount > maxDiscount{
				amountDiscount = (float64(sum) / 100) * float64(maxDiscount)
				fmt.Println("Ваша скидка составит ", amountDiscount)
			}
			break
		}else if sum < 0 || discount < 0{
			fmt.Println("Проверьте данные")
		}
	}
}

func discount(){
	var sum, discount int
	var amountDiscount float64
	isSum := sum > 0 && discount > 0

	for {
		fmt.Println("Введите сумму товара и скидку")
		fmt.Scan(&sum, &discount)
		if isSum {
			amountDiscount = (float64(sum) / 100) * float64(discount)
			if discount < maxDiscount || amountDiscount < maxAmountDiscount {
				fmt.Println("Ваша скидка составит ", amountDiscount)
			} else if amountDiscount > float64(maxAmountDiscount){
				fmt.Println("Ваша скидка составит ", maxAmountDiscount)
			} else if discount > maxDiscount{
				amountDiscount = (float64(sum) / 100) * float64(maxDiscount)
				fmt.Println("Ваша скидка составит ", amountDiscount)
			}
			break
		}else if sum < 0 || discount < 0{
			fmt.Println("Проверьте данные")
		}
	}
}

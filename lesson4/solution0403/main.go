package main

import "fmt"

func main() {
	var amount int
	denomination := 100
	maxWithdrawalAmount := 100000

	fmt.Println("Введите сумму снятия со счета")
	fmt.Println("Минимум ", denomination)
	fmt.Println("Максимальная сумма снятия ", maxWithdrawalAmount)
	fmt.Scan(&amount)

	if amount < denomination{
		fmt.Println("Отказано.\nСумма меньше минимального снятия")
	}else if amount > maxWithdrawalAmount {
		fmt.Println("Отказано.\nПревышена сумма максимального снятия")
	}else {
		fmt.Printf("Вы точно хотите снять %d", amount)
	}


}
func init() {
	fmt.Println("*****Банкомат*****")
}

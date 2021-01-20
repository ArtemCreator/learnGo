package main

import (
	"fmt"
)

func main() {
	var growth int
	var ate int
	var day int
	var startHeight int
	var targetHeight int

	fmt.Println("\tНа бамбуковой плантации завелись гусеницы. Они спят днем и едят бамбук ночью.")
	fmt.Println("Введите целое число, на какую высоту вырастает бамбук за день?")
	fmt.Scan(&growth)
	fmt.Println("Введите целое число, сколько съедят гусеницы за ночь?")
	fmt.Scan(&ate)
	fmt.Println("Нам важно знать какой высоты достигнет бамбук в середине дня. Введите целое число, сколько\n" +
		"дней расти?")
	fmt.Scan(&day)
	fmt.Println("Введите целое число, какая стартовая высота саженца бамбука в сантиметрах?")
	fmt.Scan(&startHeight)
	fmt.Printf("Бамбук генно-модифицированный, растет только при свете дня, зато  очень быстро - по %d см\n"+
		"ежедневно! Гусеницы съедают %d из них каждую ночь. Бамбуковые саженцы при высадке утром имеют высоту "+
		"%d сантиметров\n", growth, ate, startHeight)
	fmt.Printf("Какой высоты бамбук будет в середине %d дня?\n", day)
	height := startHeight + (growth-ate)*day + growth/2
	fmt.Printf("Высота бамбука на %d день состовит %d сантиметров\n\n", day, height)

	fmt.Println("Введите целое число, любую целевую высоту взрослого бамбука в сантиметрах?")
	fmt.Scan(&targetHeight)
	fmt.Printf("\tCколько полных дней понадобится бамбуку, чтобы его можно было срубить и продать. Для этого\n"+
		"он должен вырасти до %d сантиметров\n", targetHeight)
	days := (targetHeight - startHeight) / (growth - ate)
	fmt.Printf("Что бы бамбуку достичь %d сантиметров понадобится %d полных суток\n", targetHeight, days)
	fmt.Println("Если считать по формуле из задания будет другой результат")
	dayss := (targetHeight-startHeight-growth)/(growth-ate) + 1
	fmt.Printf("Что бы бамбуку достичь %d сантиметров понадобится %d полных суток\n", targetHeight, dayss)
}

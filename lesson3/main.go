package main

import "fmt"

func main() {
	fmt.Println("Задание 4")
	//solution0304()
	//solution030401()
	fmt.Println("Задание 5")
	//solution0305()
	solution03051()
}

func solution0304() {
	growth := 50
	ate := 20
	const day int = 3
	startHeight := 100
	targetHeight := 300

	fmt.Printf("\tНа бамбуковой плантации завелись гусеницы. Они спят днем и едят бамбук ночью. Бамбук\n"+
		"генно-модифицированный, растет только при свете дня, зато  очень быстро - по %d см ежедневно! Гусеницы\n"+
		"съедают %d из них каждую ночь. Бамбуковые саженцы при высадке утром имеют высоту 1 метр\n", growth, ate)
	fmt.Println("Какой высоты бамбук будет в середине третьего дня?")
	height := startHeight + (growth-ate)*day + growth/2
	fmt.Printf("Высота бамбука на %d день состовит %d сантиметров\n\n", day, height)

	fmt.Println("\tCколько полных дней понадобится бамбуку, чтобы его можно было срубить и продать. Для этого\n" +
		"он должен вырасти до 3 метров")
	days := (targetHeight - startHeight) / (growth - ate)
	fmt.Printf("Что бы бамбуку достичь 3 метров понадобится %d полных суток\n", days)
}

func solution030401() {
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

func solution0305() {
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	c := 345
	a, b, c = c, a, b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

}

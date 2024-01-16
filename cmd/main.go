package main

import (
	"fmt"
)

func main() {
	var numberOfPeople int

	// Ввод количества людей
	fmt.Print("Введите количество людей: ")
	fmt.Scan(&numberOfPeople)

	// Создаем карту для хранения взносов каждого человека
	individualShares := make(map[string]float64)

	// Ввод индивидуальных блюд и их цен для каждой персоны
	for i := 1; i <= numberOfPeople; i++ {
		person := fmt.Sprintf("Человек %d", i)
		fmt.Printf("\nВведите информацию для %s:\n", person)

		// Ввод количества индивидуальных блюд и их цен
		var numIndividualItems int
		fmt.Print("Введите количество индивидуальных блюд: ")
		fmt.Scan(&numIndividualItems)

		individualTotal := 0.0

		for j := 1; j <= numIndividualItems; j++ {
			var itemName string
			var itemPrice float64

			fmt.Printf("Введите название индивидуального блюда %d: ", j)
			fmt.Scan(&itemName)

			fmt.Printf("Введите цену индивидуального блюда %d для %s: ", j, person)
			fmt.Scan(&itemPrice)

			individualTotal += itemPrice
		}

		individualShares[person] = individualTotal
	}

	// Ввод общих блюд и их общей стоимости
	fmt.Print("\nВведите количество общих блюд: ")
	var numSharedItems int
	fmt.Scan(&numSharedItems)

	var sharedTotal float64

	for i := 1; i <= numSharedItems; i++ {
		var itemName string
		var itemPrice float64

		fmt.Printf("Введите название общего блюда %d: ", i)
		fmt.Scan(&itemName)

		fmt.Printf("Введите цену общего блюда %d: ", i)
		fmt.Scan(&itemPrice)

		sharedTotal += itemPrice
	}

	// Распределяем общую стоимость между всеми людьми
	sharedPerPerson := sharedTotal / float64(numberOfPeople)

	// Добавляем общую стоимость к индивидуальным взносам
	for person := range individualShares {
		individualShares[person] += sharedPerPerson
	}

	// Выводим результат
	fmt.Println("\nРезультат разделения счета:")
	for person, share := range individualShares {
		fmt.Printf("%s должен(-на) заплатить: %.2f\n", person, share)
	}
}
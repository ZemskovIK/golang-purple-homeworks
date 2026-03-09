package main

import "fmt"

func readAmount() float64 {
	var amount float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	return amount
}

func convert(amount float64, fromCurrency string, toCurrency string) {
	// pass
}

func main() {
	const USDToEUR = 0.86
	const USDToRUB = 79.15
	const EURToRUB = USDToRUB / USDToEUR
}
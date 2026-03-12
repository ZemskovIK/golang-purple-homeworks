package main

import "fmt"

const (
	USDToEUR = 0.86
	USDToRUB = 79.15
	EURToRUB = USDToRUB / USDToEUR
)

func readAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}
		return amount
	}
}

func readCurrency(query string, exclude string) string {
	var currency string
	fmt.Println(query)
	for {
		if exclude == "" {
			fmt.Println("Доступные валюты: USD EUR RUB")
		}
		if exclude == "USD" {
			fmt.Println("Доступные валюты: EUR RUB")
		}
		if exclude == "EUR" {
			fmt.Println("Доступные валюты: USD RUB")
		}
		if exclude == "RUB" {
			fmt.Println("Доступные валюты: USD EUR")
		}
		fmt.Scan(&currency)
		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			if currency != exclude {
				return currency
			}
		}
		fmt.Println("Неверная валюта. Попробуйте снова.")
	}
}

func convert(amount float64, fromCurrency string, toCurrency string) float64 {
	switch fromCurrency {
	case "USD":
		switch toCurrency {
		case "RUB":
			return amount * USDToRUB
		case "EUR":
			return amount * USDToEUR
		}
	case "RUB":
		switch toCurrency {
		case "USD":
			return amount / USDToRUB
		case "EUR":
			return amount / EURToRUB
		}
	case "EUR":
		switch toCurrency {
		case "RUB":
			return amount * EURToRUB
		case "USD":
			return amount / USDToEUR
		}
	}
	return 0
}


func main() {
	fmt.Println("Калькулятор валют")

	fromCurrency := readCurrency("Введите исходную валюту:", "")
	amount := readAmount()
	toCurrency := readCurrency("Введите целевую валюту:", fromCurrency)

	result := convert(amount, fromCurrency, toCurrency)
	fmt.Printf("Результат: %.2f %s", result, toCurrency)
}
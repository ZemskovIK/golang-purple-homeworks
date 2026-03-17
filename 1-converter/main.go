package main

import (
	"fmt"
	"strconv"
)

const (
	USDToEUR = 0.86
	USDToRUB = 79.15
	EURToRUB = USDToRUB / USDToEUR
)

var rates = map[string]map[string]float64{
	"USD": {
		"EUR": USDToEUR,
		"RUB": USDToRUB,
	},
	"EUR": {
		"USD": 1 / USDToEUR,
		"RUB": EURToRUB,
	},
	"RUB": {
		"USD": 1 / USDToRUB,
		"EUR": 1 / EURToRUB,
	},
}

func getAmount() (float64, error) {
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, err
	}
	amount, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func readAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		amount, err := getAmount()
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
		switch exclude {
		case "":
			fmt.Println("Доступные валюты: USD EUR RUB")
		case "USD":
			fmt.Println("Доступные валюты: EUR RUB")
		case "EUR":
			fmt.Println("Доступные валюты: USD RUB")
		case "RUB":
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
	rate := rates[fromCurrency][toCurrency]
	return amount * rate
}

func main() {
	fmt.Println("Калькулятор валют")

	fromCurrency := readCurrency("Введите исходную валюту:", "")
	amount := readAmount()
	toCurrency := readCurrency("Введите целевую валюту:", fromCurrency)

	result := convert(amount, fromCurrency, toCurrency)
	fmt.Printf("Результат: %.2f %s", result, toCurrency)
}

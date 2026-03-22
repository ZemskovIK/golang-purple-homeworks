package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readValue() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	return strings.TrimSpace(str), err
}

func readOperation() string {
	for {
		fmt.Println("Выберите операцию: AVG (среднее), SUM (сумма), MED (медиана)")
		operation, _ := readValue()
		operation = strings.ToUpper(operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		}
		fmt.Println("Неверная операция. Попробуйте снова.")
	}
}

func readNumbers() []float64 {
	for {
		fmt.Printf("Введите числа через запятую: ")
		str, _ := readValue()
		trimmedStr := strings.TrimSpace(str)
		if trimmedStr != "" {
			parts := strings.Split(trimmedStr, ",")
			numbers := []float64{}
			validInput := true
			for _, part := range parts {
				part = strings.TrimSpace(part)
				number, err := strconv.ParseFloat(part, 64)
				if err != nil {
					fmt.Println("Ошибка ввода данных, попробуйте снова.")
					validInput = false
					break
				}
				numbers = append(numbers, number)
			}
			if len(numbers) > 0 && validInput {
				return numbers
			}
		} else {
			fmt.Println("Пустой ввод, попробуйте снова.")
			continue
		}
	}
}

func calculateAvg(numbers []float64) float64 {
	sum := calculateSum(numbers)
	return sum / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += float64(num)
	}
	return sum
}

func calculateMed(numbers []float64) float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	n := len(sortedNumbers)
	if n%2 == 1 {
		return sortedNumbers[n/2]
	}
	return (sortedNumbers[n/2-1] + sortedNumbers[n/2]) / 2
}

func main() {
	fmt.Println("Калькулятор операций над числами")

	for {
		operation := readOperation()
		numbers := readNumbers()

		var result float64
		var operationName string

		switch operation {
		case "SUM":
			result = calculateSum(numbers)
			operationName = "Сумма"
		case "AVG":
			result = calculateAvg(numbers)
			operationName = "Среднее"
		case "MED":
			result = calculateMed(numbers)
			operationName = "Медиана"
		}

		fmt.Printf("%s: %.2f\n", operationName, result)
		fmt.Print("Хотите выполнить еще одну операцию? (y/n): ")
		answer, _ := readValue()
		if answer != "y" && answer != "Y" {
			break
		}
	}
}

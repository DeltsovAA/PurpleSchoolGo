package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("💱 Валютный калькулятор")
	fmt.Println("========================")
	fmt.Println("Поддерживаемые валюты: USD, EUR, RUB")

	fromCurrency := getCurrencyInput("Введите исходную валюту:")
	amount := getAmountInput("Введите сумму для конвертации:")
	toCurrency := getCurrencyInput("Введите целевую валюту:")

	if fromCurrency == toCurrency {
		fmt.Printf("\nРезультат: %.2f %s = %.2f %s (одинаковые валюты)\n", amount, fromCurrency, amount, toCurrency)
		return
	}

	rate, err := getExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result := amount * rate
	fmt.Printf("\nРезультат: %.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func getCurrencyInput(prompt string) string {
	for {
		fmt.Println(prompt)
		var input string
		fmt.Scanln(&input)
		input = strings.ToUpper(strings.TrimSpace(input))

		if input == "USD" || input == "EUR" || input == "RUB" {
			return input
		}

		fmt.Println("Неверная валюта. Попробуйте снова (USD, EUR, RUB).")
	}
}

func getAmountInput(prompt string) float64 {
	for {
		fmt.Println(prompt)
		var input string
		fmt.Scanln(&input)
		amount, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err == nil && amount > 0 {
			return amount
		}

		fmt.Println("Неверное число. Введите положительное значение.")
	}
}

func getExchangeRate(from, to string) (float64, error) {
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return 0.9, nil
		case "RUB":
			return 90, nil
		}
	case "EUR":
		switch to {
		case "USD":
			return 1.1, nil
		case "RUB":
			return 100, nil
		}
	case "RUB":
		switch to {
		case "USD":
			return 0.011, nil
		case "EUR":
			return 0.01, nil
		}
	}
	return 0, errors.New("нет доступного курса обмена для выбранных валют")
}

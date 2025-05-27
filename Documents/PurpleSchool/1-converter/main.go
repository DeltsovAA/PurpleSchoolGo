package main

import "fmt"

func main() {

	const usdToEur = 0.9261
	const usdToRub = 89.34

	result := usdToEur / usdToRub
	fmt.Println("Калькулятор валют:")
	fmt.Printf("1 USD = %.4f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.4f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.4f RUB (вычислено на основе первых двух)\n", result)
}

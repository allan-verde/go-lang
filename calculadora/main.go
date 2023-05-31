package main

import (
	"errors"
	"fmt"
)

func main() {
	resultadoSoma := Add(5, 3)
	fmt.Println("Soma:", resultadoSoma)

	resultadoSubtracao := Subtract(5, 3)
	fmt.Println("Subtração:", resultadoSubtracao)

	resultadoMultiplicacao := Multiply(5, 3)
	fmt.Println("Multiplicação:", resultadoMultiplicacao)

	resultadoDivisao, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Divisão:", resultadoDivisao)
	}
}

func Add(numbers ...float64) float64 {
	result := 0.0

	for _, number := range numbers {
		result += number
	}

	return result
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(numbers ...float64) float64 {
	result := 1.0

	for _, number := range numbers {
		result *= number
	}

	return result
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return a / b, nil
}

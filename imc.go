package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, massa float64

	altura = 1.80
	massa = 90 //kg

	imc := massa / math.Pow(altura, 2)
	fmt.Println("\nIMC para os dados inseridos: ", imc, "\n")
}

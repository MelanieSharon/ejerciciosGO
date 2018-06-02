package main

import (
	"fmt"
)

func main() {
	conditional(1)
	fmt.Println(conditionalIF(21))
	fmt.Println(conditionalDuo(0))
	fmt.Println(respuestaCondition(conditionalDuo(0)))
}

func conditional(number int) {
	switch number {
	case 0:
		fmt.Println("Hola")
	case 1:
		fmt.Println("vida")
	case 2:
		fmt.Println("Feo")
	default:
		fmt.Println("Melanie")
	}
}

func conditionalIF(number int) string {
	var name = ""
	if number <= 10 {
		name = "Ninio"
	} else if number > 10 && number <= 17 {
		name = "Adolecente"
	} else if number > 17 && number <= 30 {
		name = "Joven"
	} else {
		name = "Viejo"
	}
	return name
}

func conditionalDuo(number int) float64 {
	var rpt = 0.0

	switch number {
	case 0:
		rpt = 5 * 0.20
	case 1:
		rpt = 5.20 * 0.15
	case 2:
		rpt = 6.0 * 8.0
	}

	return rpt
}

func respuestaCondition(rpt float64) string {
	var result = ""
	if rpt <= 2 {
		result = "saldo completo"
	} else if rpt > 2 && rpt <= 4 {
		result = "saldo incompleto"
	} else {
		result = "Error sistema"
	}

	return result
}

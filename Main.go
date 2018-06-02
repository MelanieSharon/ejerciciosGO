package main

import (
	"fmt"
)

func main() {
	conditional(1)
	fmt.Println(conditionalIF(21))
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

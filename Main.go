package main

import (
	"fmt"
)

func main() {
	conditional(1)
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

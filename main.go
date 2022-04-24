package main

import (
	"fmt"
)

func main() {
	option := 0
	fmt.Println("Hola Mundo, ¿Qué deseas hacer?")
	fmt.Println("> 1. Imprmir Hola Mundo")
	fmt.Println("> 2. Imprmir Hola Mundo")
	fmt.Println("> 3. Imprmir Hola Mundo")
	fmt.Scanf("%d", &option)
	switch option {
	case 1:
		fmt.Println("Hola Mundo")
	case 2:
		fmt.Println("Hola Mundo")
	case 3:
		fmt.Println("Hola Mundo")
	default:
		fmt.Println("Opcion no valida")
	}
}

package main

import (
	"fmt"
)

var numero1 int
var numero2 int

func main() {
	fmt.Println("Ingrese el numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el numero 2: ")
	fmt.Scan(&numero2)

	fmt.Println("Los numeros a sumar son: ", numero1, "+", numero2)
	fmt.Printf("%d", numero1+numero2)
}

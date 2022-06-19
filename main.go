package main

import "fmt"

func main() {

	var a, b int

	//Entrada
	fmt.Print("Ingrese numero 01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese número 02: ")
	fmt.Scanln(&b)

	//Proceso
	suma := a + b
	resta := a - b
	multi := a * b
	divi := a / b
	mod := a % b

	//Salida
	fmt.Println("La suma de los dos numeros es: ", suma)
	fmt.Println("La resta de los dos numeros es: ", resta)
	fmt.Println("El producto de los dos numeros es: ", multi)
	fmt.Println("El cociente de los dos numeros es: ", divi)
	fmt.Println("El modulo es:", mod)

}

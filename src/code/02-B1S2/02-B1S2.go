package main

import "fmt"

func testing_variables() {

	var numero int = 42
	var decimal float64 = 3.1415
	var verdad bool = true
	var caracter rune = 'a'
	var texto byte = 'H'
	var complejo complex64 = 1 + 2i
	fmt.Println("Número:", numero)
	fmt.Println("Decimal:", decimal)
	fmt.Println("Booleano:", verdad)
	fmt.Println("Rune (Unicode):", caracter)
	fmt.Println("Byte (Caracter):", texto)
	fmt.Println("Número Complejo:", complejo)

}

func main() {
	testing_variables()
}

// go run src\code\02-B1S2\02-B1S2.go

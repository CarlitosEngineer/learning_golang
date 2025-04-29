package main

import "fmt"


// Existen varias formas de declarar una variable

// De la forma Explicita
var nombre string = "Carlitos"

// De la forma Implicita
var apodo = "Tangas"

// Incluso multiples variablers.
var a, b, c int = 1, 2, 3

// De forma corta y de poco alcance,solo dentro de funciones
// pais := "mexico"

func main() {
	// pais := "mexico" // Example
	fmt.Println("Hello, World!")
}

/*

- Esta es uan sintaxys basica de Go o Golang
	- Todo los programas empiezan por main
	- se importan librerias para nuevasfunciones como lo es el fmt
	- se escribe func para declarar funciones etc.

*/

// go run src\code\01-B1S1\01-B1S1.go
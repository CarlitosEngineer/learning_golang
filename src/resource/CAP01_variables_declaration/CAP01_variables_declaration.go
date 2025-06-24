package main

import (
	"fmt"
)

func learning_variable_declaration() {

	var nombre string = "Carlitos" // Forma explicita
	var apodo = "Tangas"           // De la forma Implicita
	var a, b, c int = 1, 2, 3      // Incluso multiples variablers.
	pais := "mexico"               // Forma corta y poco alcance, Solo dentro de funciones
	// IMPRIMIENDO RESULTADOS
	fmt.Println("FORMA EXPLICITA:", nombre)
	fmt.Println("FORMA IMPLICITA:", apodo)
	fmt.Println("MULTI VAR", a, b, c)
	fmt.Println("VAR CORTA", pais)

}

func main() {
	learning_variable_declaration()
}

// go run src\resource\01-variable_declaration\01-variable_declaration.go

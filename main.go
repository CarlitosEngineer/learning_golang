package main

import (
	"fmt"
	"learning_golang/src/bin" // 👈 tu módulo y ruta relativa
)

func main() {
	fmt.Println("Bienvenido a Go Learning Project 🚀")
	// Learning Data Types
	bin.Ejecutar_cap1Go()
	// Learning Strings
	bin.Testing_RawStringsLiterals1()
	bin.Testing_InterpretedStringLiterals1()
	bin.Testing_StringExtras1()

}

// go run main.go

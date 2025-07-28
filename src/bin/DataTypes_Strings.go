package bin

import (
	"fmt"
	"regexp"
)

// Raw String Literals (Cadenas sin procesar)
func Testing_RawStringsLiterals1() {
	variable1 := `Esto es una cadena sin procesar, tal cual como está1`
	variable2 := `Esto es una cadena sin procesar,
				tal cual como está2`
	fmt.Println("Constante Pi:", variable1, variable2)
}

// Interpreted String Literals (Cadenas interpretadas)
func Testing_InterpretedStringLiterals1() {
	variable3 := "Esta cadena tiene\nsaltos de línea y\t tabulaciones." //, Doble comillas,  {\n, \t, \", \\}
	fmt.Println("Constante Pi:", variable3)
}

func Testing_StringExtras1() {
	// Raw Strings para expresiones regulares:
	variable4 := regexp.MustCompile(`^\d{3}-\d{2}-\d{4}$`)
	// Raw Strings para html/template o text/template:
	variable5 := `Hola {{.Nombre}}, bienvenido.`
	// Interpreted Strings con fmt.Sprintf:
	variable6 := "Juan"
	variable7 := fmt.Sprintf("Hola, %s", variable6)
	fmt.Println("Variable1", variable4, variable5, variable6, variable7)
}

// go run src\bin\DataTypes_Strings.go

/*

APUNTES Y NOTAS:

	Interpretan secuencias de escape:
		\n → salto de línea
		\t → tabulación
		\" → comillas dentro de la cadena
		\\ → barra invertida

	- grave/backticks `

*/

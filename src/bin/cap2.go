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
	fmt.Println("Raw 1:\n", variable1)
	fmt.Println("Raw 2:\n", variable2)

}

// Interpreted String Literals (Cadenas interpretadas)
func Testing_InterpretedStringLiterals1() {
	variable3 := "Esta cadena tiene\nsaltos de línea y\t tabulaciones."
	fmt.Println("Testing_InterpretedStringLiterals1:", variable3)
}

// Consejos Extras
func Testing_StringExtras1() {
	// Raw Strings para expresiones regulares:
	variable4 := regexp.MustCompile(`^\d{3}-\d{2}-\d{4}$`)

	// Caso de uso : Este ejemplo usa expresiones regulares (regex) para validar un formato de texto.
	testSSN := "123-45-6789" // SSN (Social Security Number de EE.UU.)
	fmt.Println("¿Es válido?", variable4.MatchString(testSSN))

	// Raw Strings para html/template o text/template:
	variable5 := `Hola {{.Nombre}}, bienvenido.`
	// Interpreted Strings con fmt.Sprintf:
	variable6 := "Juan"
	variable7 := fmt.Sprintf("Hola, %s", variable6)
	fmt.Println("Testing_StringExtras1:", variable4, variable5, variable6, variable7)

	// Caso de uso 1: Go optimiza internamente los strings como inmutables. se pueden usar escapes Unicode
	unicodeStr := "Símbolo Unicode: \u263A"
	fmt.Println(unicodeStr)

}

/*

APUNTES Y NOTAS:

	Interpretan secuencias de escape:
		\n → salto de línea
		\t → tabulación
		\" → comillas dentro de la cadena
		\\ → barra invertida

	Tipos de comillas
		`` -> backticks
		"" -> Comillas dobles

	// Raw String Literals (Cadenas sin procesar) Enfocado a => Datos literales, código
		- Las cadenas sin procesar usan backticks.
		- No interpretan secuencias de escape.
		- Mantienen el formato original del texto.

	// Interpreted String Literals (Cadenas interpretadas) Enfocado a => Texto con formato/control
		- Texto con caracteres especiales.
		- Mensajes logs y consola.
		- Strings con un formato especifico.

*/

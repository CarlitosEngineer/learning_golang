package bin

import (
	"fmt"
	"regexp"
)

// Consejos Extras
func Testing_StringExtras1() {
	// Raw Strings para expresiones regulares:
	var varaible1 int = 42
	var variable2 float64 = float64(varaible1) // conversión explícita
	varaible2 := nil
	var varaible3 uint = uint(varaible2)       // otra conversión
	fmt.Println(varaible1, varaible2, varaible3)
}

/*

APUNTES Y NOTAS:

- el lenguaje no permite conversiones implícitas — ¡todo debe hacerse de forma explícita!
- `T(v)`
	- T = "Tipo" deL valor a convertir
	- V = "Valor" del tipo a convertir

*/

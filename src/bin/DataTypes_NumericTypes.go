package bin

import (
	"fmt"
	"math/cmplx"
)

func TestingVariables1() {
	// Constantes
	const variablePi = 3.14159
	const variableGreeting string = "Hola desde Go"

	// Enteros
	var variable1 int = 42                    // [32Bits: -2,147,483,648 a 2,147,483,647 - 64Bits: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807]
	var variable2 int8 = -128                 // [-128 a 127]
	var variable3 int16 = 30000               // [-32,768 a 32,767]
	var variable4 int32 = 1000000000          // [-2,147,483,648 a 2,147,483,647]
	var variable5 int64 = 9223372036854775807 // [-9,223,372,036,854,775,808 a 9,223,372,036,854,775,807]

	// Enteros sin signo
	var variable6 uint = 42                      // [32Bits: 0 a 4,294,967,295 - 64Bits: 0 a 18,446,744,073,709,551,615]
	var variable7 uint8 = 255                    // [0 a 255]
	var variable8 uint16 = 65535                 // [0 a 65,535]
	var variable9 uint32 = 4294967295            // [0 a 4,294,967,295]
	var variable10 uint64 = 18446744073709551615 // [0 a 18,446,744,073,709,551,615]

	// Flotantes
	var variable11 float32 = 3.1415            // [±1.18×10⁻³⁸ a ±3.4×10³⁸] con ~6-7 dígitos decimales de precisión
	var variable12 float64 = 2.718281828459045 // [±2.23×10⁻³⁰⁸ a ±1.8×10³⁰⁸] con ~15-17 dígitos decimales de precisión

	// Booleano
	var variable13 bool = true // [Valores posibles: true o false]

	// Byte y rune
	var variable14 byte = 'A' // [byte es alias de uint8: 0 a 255] Representa un caracter ASCII o un solo byte de datos binarios
	var variable15 rune = '😊' // [rune es alias de int32: -2,147,483,648 a 2,147,483,647] Representa un carácter Unicode (UTF-8 extendido)
	// Un rune es un alias de int32. Representa un carácter Unicode (UTF-8 extendido). Es útil para trabajar con cadenas que contienen emojis, letras acentuadas, o cualquier símbolo internacional.

	// Complejos
	var variable16 complex64 = complex(1.5, 2.5)     // [Basado en float32: ±1.18×10⁻³⁸ a ±3.4×10³⁸ por parte] Representa números complejos (real + imaginaria) con precisión de 32 bits
	var variable17 complex128 = cmplx.Sqrt(-5 + 12i) // [Basado en float64: ±2.23×10⁻³⁰⁸ a ±1.8×10³⁰⁸ por parte] Mayor precisión en operaciones complejas

	// uintptr (muy raro de usar)
	var variable18 uintptr = 0x12345678

	// Zero values - En Go, cuando declaras una variable sin asignarle un valor explícito, el compilador le asigna un valor por defecto, conocido como zero value.
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultBool bool     // false
	var defaultString string // "" (cadena vacía)
	// slice, map, chan, interface, func, pointer // nil

	// Declaración rápida
	shorthand := "Variable declarada con :="

	// Mostramos todos
	fmt.Println("Constante Pi:", variablePi)
	fmt.Println("Constante Greeting:", variableGreeting)
	fmt.Println("int:", variable1)
	fmt.Println("int8:", variable2)
	fmt.Println("int16:", variable3)
	fmt.Println("int32:", variable4)
	fmt.Println("int64:", variable5)
	fmt.Println("uint:", variable6)
	fmt.Println("uint8:", variable7)
	fmt.Println("uint16:", variable8)
	fmt.Println("uint32:", variable9)
	fmt.Println("uint64:", variable10)
	fmt.Println("float32:", variable11)
	fmt.Println("float64:", variable12)
	fmt.Println("bool:", variable13)
	fmt.Println("byte (uint8):", variable14)
	fmt.Println("rune (int32, unicode):", variable15)
	fmt.Println("complex64:", variable16)
	fmt.Println("complex128:", variable17)
	fmt.Println("uintptr:", variable18)
	fmt.Println("Zero int:", defaultInt)
	fmt.Println("Zero float64:", defaultFloat)
	fmt.Println("Zero bool:", defaultBool)
	fmt.Println("Zero string:", defaultString)
	fmt.Println("Shorthand:", shorthand)
	fmt.Printf("Tipo: %T, Valor: %v\n", variable14, variable14) // imprime tipo de dato mas su valor
	// go run src\bin\DataTypes_NumericTypes.go
}

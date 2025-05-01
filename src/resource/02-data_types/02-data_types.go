package main

import (
	"fmt"
	"math/cmplx"
)

func testing_variables() {
	// Enteros
	var a int = 42
	var b int8 = -120
	var c int16 = 30000
	var d int32 = 1000000000
	var e int64 = 9223372036854775807

	// Enteros sin signo
	var f uint = 42
	var g uint8 = 255
	var h uint16 = 65535
	var i uint32 = 4294967295
	var j uint64 = 18446744073709551615

	// Flotantes
	var k float32 = 3.1415
	var l float64 = 2.718281828459045

	// Booleano
	var m bool = true

	// Byte y rune
	var n byte = 'A' // alias de uint8
	var o rune = '😊' // alias de int32 (unicode)

	// Complejos
	var p complex64 = complex(1.5, 2.5)
	var q complex128 = cmplx.Sqrt(-5 + 12i)

	// uintptr (raro que lo uses, pero existe)
	var r uintptr = 0x12345678

	// Mostramos todos
	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)
	fmt.Println("uint:", f)
	fmt.Println("uint8:", g)
	fmt.Println("uint16:", h)
	fmt.Println("uint32:", i)
	fmt.Println("uint64:", j)
	fmt.Println("float32:", k)
	fmt.Println("float64:", l)
	fmt.Println("bool:", m)
	fmt.Println("byte (uint8):", n)
	fmt.Println("rune (int32, unicode):", o)
	fmt.Println("complex64:", p)
	fmt.Println("complex128:", q)
	fmt.Println("uintptr (dirección simulada):", r)
}

func main() {
	testing_variables()
}

// go run src\resource\02-data_types\02-data_types.go

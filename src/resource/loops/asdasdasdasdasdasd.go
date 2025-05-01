package main

import (
	"errors"
	"fmt"
)

// BUCLE --> for
func testing_loop_for() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración:", i)
	}
}

// BUCLE --> while
func testing_loop_while() {
	i := 0
	for i < 3 {
		fmt.Println("Valor de i:", i)
		i++
	}
}

// BUCLE --> Infinite loop
func Infinite_loop() {
	for {
		fmt.Println("Bucle infinito")
		break // ¡recuerda salir o será eterno!
	}
}

// BUCLE --> `range` (para recorrer slices, arrays, maps, etc.)
func testing_loop_range() {
	nombres := []string{"Luis", "Carla", "Diana"}
	for i, nombre := range nombres {
		fmt.Printf("Posición %d: %s\n", i, nombre)
	}
}

// BUCLE --> for range con slice
func testing_loop_for_slice() {
	nombres := []string{"Luis", "Carla", "Diana"}
	for i, nombre := range nombres {
		fmt.Printf("Posición %d: %s\n", i, nombre)
	}
}

// BUCLE --> for range con strings
func testing_loop_for_strings() {
	palabra := "¡Hola!"
	for i, char := range palabra {
		fmt.Printf("Índice %d - Rune %c\n", i, char)
	}
}

// BUCLE --> for range con map
func testing_loop_for_map() {
	capitales := map[string]string{
		"México":  "CDMX",
		"Francia": "París",
		"Japón":   "Tokio",
	}
	for pais, ciudad := range capitales {
		fmt.Printf("La capital de %s es %s\n", pais, ciudad)
	}
}

// BUCLE --> for con continue y break
func testing_loop_range_continue_break() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Salta el valor 3
		}
		if i == 5 {
			break // Sale del bucle al llegar a 5
		}
		fmt.Println(i)
	}

}

// `if` y `switch`
func testing_loop_if_switch() {
	edad := 20
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	dia := "lunes"
	switch dia {
	case "lunes":
		fmt.Println("Inicio de semana")
	case "viernes":
		fmt.Println("Casi finde!")
	default:
		fmt.Println("Otro día cualquiera")
	}
}

// `error`, `panic`, `recover`
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}

func testing_loop_error_panic_recover() {

	// Manejo de errores
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// Panic y recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se recuperó de un panic:", r)
		}
	}()

	fmt.Println("Antes del panic")
	panic("¡Algo salió muy mal!")
	fmt.Println("Este mensaje no se imprimirá") // nunca se ejecuta

}

func main() {
	fmt.Println("Change Topic HERE!")
	testing_loop_for()
	fmt.Println("Change Topic HERE!")
	testing_loop_while()
	fmt.Println("Change Topic HERE!")
	Infinite_loop()
	fmt.Println("Change Topic HERE!")
	testing_loop_range()
	fmt.Println("Change Topic HERE!")
	testing_loop_for_strings()
	fmt.Println("Change Topic HERE!")
	testing_loop_for_map()
	fmt.Println("Change Topic HERE!")
	testing_loop_range_continue_break()
	fmt.Println("Change Topic HERE!")
	testing_loop_if_switch()
	fmt.Println("Change Topic HERE!")
	testing_loop_error_panic_recover()
	fmt.Println("FINISH HERE!")
}

// go run src\code\03-B1S3\03-B1S3.go

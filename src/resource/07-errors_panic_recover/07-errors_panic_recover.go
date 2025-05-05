package main

import (
	"errors"
	"fmt"
)

// Example 1 - Aprendiendo a usar errores, FUNCIÓN DIVIDIR
func dividir(a, b int) (int, error) { // Recibe dos valores y devuelve dos valores: uno entero (resultado) y uno error (si llegase a ocurrir ocurre)
	if b == 0 { // Si el parametro b, llegase a recibir cero, debera ejecutar el siguiente bloque
		return 0, errors.New("Que te pasa crack, no sabes como dividir?") // devuelve 0 por default y un error con un mensaje propio
	}
	return a / b, nil // Y, si no es error. divirlos y devolverlos
}

// Example 1 - Funcion para probar el error
func learning_errors_panic_recover1() {
	resultado, err := dividir(10, 0)
	if err != nil { // Si sale mal, imprimir mensaje y el error que devuelva la funcion.
		fmt.Println("Ocurrió un error:", err)
		return
	}
	fmt.Println("Resultado:", resultado) // Si no sale mal, solo imprimir el resultado.
}

// Example 2 - Funcion para probar el PANIC!
func verificarNombre(nombre string) {
	// `defer` es una funcion que se ejecutara al final "PASE LO QUE PASE", con o sin panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Tranquilo, chico. el panic ha sido recuperado:", r)
		}
	}()

	if nombre == "" {
		panic("Recuerda que el nombre no puede estar vacio, man") // PANIC: Error grave
	}
	fmt.Println("Este nombre si es valido:", nombre)
}

func main() {
	learning_errors_panic_recover1() // LEARNING ERRORES
	verificarNombre("Carlos")        // LEARNING PANIC
	verificarNombre("")              // LEARNING PANIC - Esto causará un panic
	fmt.Println("El programa sigue vivo 🚀")
}

// go run src\resource\07-errors_panic_recover\07-errors_panic_recover.go

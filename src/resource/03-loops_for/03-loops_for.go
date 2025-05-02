package main

import (
	"fmt"
)

// BUCLE v1 --> CLASICO del `For`
func learning_loop_forv1() {
	fmt.Println("BUCLE v1 --> `for` CLASICO")
	for i := 0; i < 5; i++ { // Inicialización ( i := 0 ), Condicion ( i < 5  ), Incremento ( i++ )
		fmt.Println("Iteración:", i) // Post
	}
}

// BUCLE v2 --> SIMULACIÓN del `While`
func learning_loop_forv2() {
	fmt.Println("BUCLE v2 --> Simulación del `while`")
	i := 0
	for i < 3 {
		fmt.Println("Valor de i:", i)
		i++
	}
}

// BUCLE v3 --> SIMULACIÓN del `Do While`
func learning_loop_forv3() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}
}

// BUCLE v4 --> SIMULACIÓN del `Do While` - Infinite loop
func learning_loop_forv4() {
	for {
		fmt.Println("Bucle infinito")
		break // ¡recuerda salir o será eterno!
	}
}

// LLAMANDO A LOS CICLOS
func main() {
	learning_loop_forv1() // BUCLE v1 --> CLASICO del `For`
	learning_loop_forv2() // BUCLE v2 --> SIMULACIÓN del `While`
	learning_loop_forv3() // BUCLE v3 --> SIMULACIÓN del `Do While`
	learning_loop_forv4() // BUCLE v4 --> SIMULACIÓN del `Do While` - Infinite loop
}

// go run src\resource\03-loops_for\03-loops_for.go

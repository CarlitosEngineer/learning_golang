package main

import (
	"fmt"
)

// - Canales 📡

// BUCLE --> `for` & `range` with Arrays 🧮
func learning_loops_range_v1() {
	nombres := [3]string{"Carlitos", "Tangas", "TuPapa"}
	for i, nombre := range nombres {
		fmt.Printf("Posición %d: %s\n", i, nombre)
	}
}

// BUCLE --> `for` & `range` with Slices 📦
func learning_loops_range_v2() {
	nombres := []string{"Carlitos", "Tangas", "TuPapa"}
	for i, nombre := range nombres {
		fmt.Printf("Posición %d: %s\n", i, nombre)
	}
}

// BUCLE --> `for` & `range` with Maps 🗺️
func learning_loops_range_v3() {
	capitales := map[string]string{
		"México":  "CDMX",
		"Francia": "París",
		"Japón":   "Tokio",
	}
	for pais, ciudad := range capitales {
		fmt.Printf("La capital de %s es %s\n", pais, ciudad)
	}
}

// BUCLE --> `for` & `range` with Strings 🔤
func learning_loops_range_v4() {
	palabra := "¡Hola!"
	for i, char := range palabra {
		fmt.Printf("Índice %d - Rune %c\n", i, char)
	}
}

func main() {
	learning_loops_range_v1()
	learning_loops_range_v2()
	learning_loops_range_v3()
	learning_loops_range_v4()

}

// go run src\resource\04-collection_range\04-collection_range.go

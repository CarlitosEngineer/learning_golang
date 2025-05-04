package main

import (
	"fmt"
)

// Example 1
func learning_control_flow_switch1() {
	day := "martes"
	switch day {
	case "lunes":
		fmt.Println("Empezo es joda!")
	case "martes", "miércoles": // Se pueden agrupar multiples valores, para una sola situación.
		fmt.Println("Mitad de la joda.")
	case "viernes":
		fmt.Println("Finalizo la Joda!")
	default:
		fmt.Println("Soportando la joda!")
	}
}

// Example 2
func learning_control_flow_switch2() {
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Tremenda puntuación papu!")
	case score >= 70:
		fmt.Println("Bien echo crack, pero puedes mejorar")
	default:
		fmt.Println("GG!, Me repite el curso. papi")
	}
}

func main() {
	learning_control_flow_switch1()
	learning_control_flow_switch2()
}

// go run src\resource\06-control_flow_switch\06-control_flow_switch.go

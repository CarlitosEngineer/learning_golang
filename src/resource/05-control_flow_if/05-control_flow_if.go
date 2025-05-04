package main

import (
	"fmt"
)

// Example 1
func learning_control_flow_if1() {
	if score := 75; score >= 60 {
		fmt.Println("Aprobado")
	}
}

// Example 2
func learning_control_flow_if2() {
	age := 18

	if age >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if age > 12 {
		fmt.Println("Eres adolescente")
	} else {
		fmt.Println("usted es un culicagado!")
	}
}

// Example 3
func learning_control_flow_if3() {
	temperature := 32
	humidity := 40
	if temperature > 25 && humidity < 50 {
		fmt.Println("PRENDA ESA JODA! de ventilador carajo!")
	}
}

func main() {
	learning_control_flow_if1()
	learning_control_flow_if2()
	learning_control_flow_if3()
}

// go run src\resource\05-control_flow_if\05-control_flow_if.go

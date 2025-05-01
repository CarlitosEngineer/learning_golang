// **Conditionals**

package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("Eres mayor de edad.")
    } else {
        fmt.Println("Eres menor de edad.")
    }
}


// **Functions, Multiple/Named Returns**

// Función simple con múltiples retornos
func divide(a, b float64) (result float64, ok bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
    r, ok := divide(10, 2)
    if ok {
        fmt.Println("Resultado:", r)
    } else {
        fmt.Println("División inválida.")
    }
}


// **Packages, Imports y Exports**

// Archivo: mathutils/utils.go
package mathutils

func Sumar(a, b int) int {
    return a + b
}

// Archivo main.go
package main

import (
    "fmt"
    "tu_modulo/mathutils"
)

func main() {
    fmt.Println("Suma:", mathutils.Sumar(3, 4))
}


// **Type Casting**

func main() {
    var x int = 42
    var y float64 = float64(x) // conversión explícita
    fmt.Println(y)
}


// **Type Inference**

func main() {
    x := 10      // Go infiere que es un int
    y := 3.14    // float64
    z := "Hola"  // string
    fmt.Printf("%d, %f, %s\n", x, y, z)
}


// **Arrays**

func main() {
    var nums [3]int = [3]int{1, 2, 3}
    fmt.Println(nums[0]) // acceso al primer elemento
}


// **Slices**

func main() {
    nums := []int{1, 2, 3, 4}
    nums = append(nums, 5)
    fmt.Println(nums)
}


// **Maps**

func main() {
    edades := map[string]int{
        "Ana": 30,
        "Luis": 25,
    }
    fmt.Println(edades["Ana"])
}


// **make()**

func main() {
    m := make(map[string]int)
    m["clave"] = 42
    fmt.Println(m)
}


//  **Structs**

type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    p := Persona{Nombre: "Carlos", Edad: 28}
    fmt.Println(p.Nombre)
}


Con esto estás listo para construir estructuras de datos más complejas y organizar mejor tu código. ¿Quieres que armemos un mini-proyecto usando todo esto como práctica? 💡
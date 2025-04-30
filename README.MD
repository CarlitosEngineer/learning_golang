# **Go by Carlitos!**

Apuntes propios por mi sobre Go

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 1: First things first (Lo primero es lo primero)

- Basic Syntax

```go
package main
import "fmt"

func main() {
    fmt.Println("Hola mundo!") // Salida: Hola mundo!
}
```

- Variables and declaration

```go
var nombre string = "GoLang" // Declaración explícita con tipo
edad := 30 // Inferencia de tipo
var x, y int = 1, 2 // Declaración múltiple
var z int // z = 0 , // Variables sin inicializar -> valor cero
```

### Sub-Bloque 2: Data Types (Tipos de datos)

- int, int8/16/32/64
- uint, uint8/16/32/64
- float32, float64
- bool
- byte
- rune
- complex64/128
- uintptr

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

### Sub-Bloque 4: programming tools 2 (herramientas de programación 2)

- Conditionals
- Functions, multiple/named returns
- Packages, imports and exports
- Type Casting
- Type Inference
- Arrays
- Slices
- Maps
- make()
- Structs

## Bloque 2: Going Deeper (Profundizar en Golang)

- Proximo
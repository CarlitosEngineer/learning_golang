# **Go by Carlitos!**

## 🟦 **BLOQUE 1: Aprende lo Básico**

### ✅ **Sub-Bloque 1: Lo Primero es lo Primero**

#### 🔹 Sintaxis Básica

```go
package main
import "fmt"

func main() {
    fmt.Println("Hola mundo!") // Salida: Hola mundo!
}
```

#### 🔹 Declaración de Variables

```go
var nombre string = "GoLang" // Declaración explícita con tipo
edad := 30 // Inferencia de tipo
var x, y int = 1, 2 // Declaración múltiple
var z int // z = 0 // Variables sin inicializar -> valor cero
```

### ✅ **Sub-Bloque 2: Tipos de Datos**

#### 🔹 Números Enteros

| Tipo      | Rango                        |
|-----------|------------------------------|
| `int`     | depende de la arquitectura   |
| `int8`    | -128 a 127                   |
| `int16`   | -32,768 a 32,767             |
| `int32`   | -2,147,483,648 a 2,147,483,647 |
| `int64`   | ±9.2 cuatrillones            |

```go
var a int32 = 1000
```

#### 🔹 Números Enteros Sin Signo

| Tipo      | Rango                        |
|-----------|------------------------------|
| `uint`    | depende de la arquitectura   |
| `uint8`   | 0 a 255                      |
| `uint16`  | 0 a 65,535                   |
| `uint32`  | 0 a 4.29 mil millones        |
| `uint64`  | 0 a 18.4 cuatrillones        |

```go
var b uint16 = 65000
```

#### 🔹 Punto Flotante

```go
var pi float32 = 3.14
var e float64 = 2.71828
```

#### 🔹 Booleanos

```go
var activo bool = true
```

#### 🔹 Byte y Rune

```go
var letra byte = 'A'   // byte = uint8
var simbolo rune = '♥' // rune = int32 (unicode)
```

#### 🔹 Números Complejos

```go
var c complex64 = 2 + 3i
var d complex128 = complex(1.2, 3.4)
```

#### 🔹 uintptr (poco común, para manipular punteros)

```go
var p uintptr = uintptr(unsafe.Pointer(&x)) // requiere "unsafe"
```

---

¿Quieres que te lo prepare también en formato PDF o Markdown para imprimir o estudiar? 📄🧠
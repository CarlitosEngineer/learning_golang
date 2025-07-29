# Cheat Sheet de Go (Golang)

#### ➕ Sin Signo

| Tipo   | Bits    | Rango                            |
| ------ | ------- | -------------------------------- |
| uint8  | 8       | `0 a 255`                        |
| uint16 | 16      | `0 a 65,535`                     |
| uint32 | 32      | `0 a 4,294,967,295`              |
| uint64 | 64      | `0 a 18,446,744,073,709,551,615` |
| uint   | 32 o 64 | `Depende de la arquitectura`     |

#### ➖ Con Signo ➕

| Tipo    | Bits    | Rango                              |
| ------- | ------- | ---------------------------------- |
| `int8`  | 8       | `[-128 a 127]`                     |
| `int16` | 16      | `[32,768 a 32,767]`                |
| `int32` | 32      | `[-2,147,483,648 a 2,147,483,647]` |
| `int64` | 64      | `[± 9,223,372,036,854,775,807]`    |
| `int`   | 32 o 64 | `[Depende de la arquitectura]`     |


#### Variables & Constants

```go
var variable1 int = 30  // Explícita con tipo
variable2 := "Gopher"   // Implícita (solo dentro de funciones), osea local
const pi = 3.1416
```

##### Strings

```go
variable1 := `Un texto ejemplo1` // Raw String Literals => Uso ideal "Datos literales, código"
variable2 := "Un texto ejemplo2" // Interpreted String Literals  => Uso ideal "Texto con formato/control" 
```




```go
```



```go
```



```go
```


```go
```

# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 2: Data Types (Tipos de datos)

- int, int8/16/32/64
- uint, uint8/16/32/64
- float32, float64
- bool
- byte
- rune
- complex64/128
- uintptr

---

#### Let's get started (Empecemos)

- Go es **muy estricto** con los tipos. **int** , **int8/16/32/64** , **uint** , **uint8/16/32/64** , **float32** , **float64** , *bool* , **byte** , **rune** , **complex64/128** , **uintptr**

| Tipo       | Descripción |
|------------|-------------|
| `int`      | Entero con tamaño según la arquitectura (32 o 64 bits) |
| `int8`     | Entero de 8 bits (rango: -128 a 127) |
| `int16`    | Entero de 16 bits |
| `int32`    | Entero de 32 bits |
| `int64`    | Entero de 64 bits |
| `uint`     | Entero sin signo (positivo) |
| `uint8`    | Entero positivo de 8 bits (`byte` es un alias de `uint8`) |
| `uint16`   | Entero positivo de 16 bits |
| `uint32`   | Entero positivo de 32 bits |
| `uint64`   | Entero positivo de 64 bits |
| `float32`  | Número decimal de 32 bits |
| `float64`  | Número decimal de 64 bits (más preciso) |
| `bool`     | Booleano (`true` o `false`) |
| `byte`     | Alias de `uint8`, representa un byte |
| `rune`     | Alias de `int32`, representa un carácter Unicode |
| `complex64`| Número complejo (float32 parte real + float32 parte imaginaria) |
| `complex128`| Número complejo (float64 parte real + float64 parte imaginaria) |
| `uintptr`  | Número entero para almacenar direcciones de memoria |

#### **Recomendaciones**:

- nosotros debemos usar `int` y `float64` por defecto, los demas tipos, debemos razonar si es son necesarios usarlos.
- `rune` y `byte` para strings
- no te preocupes por `uintptr`, cosas de punteros muy especificas.
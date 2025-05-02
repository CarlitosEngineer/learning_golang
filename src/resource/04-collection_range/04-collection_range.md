# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

---

#### Let's Get Started (Empecemos)

#### Range

Es una funcion que nos va a permitir iterar sobre las collecciones de datos que tenemos en golang, existen varios tipos de colleciones de datos. que ya vimos antes.

```go
for index, value := range collection {
    // usar index y value
}
```

- 🧮 **Arrays**: Son **estructuras** de tamaño **fijo** que contienen **elementos del mismo tipo**.
- 📦 **Slices**: Son **estructuras** de tamaño **dinámico** basadas en arrays, muy usadas en Go.
- 🗺️ **Maps**: Son **estructuras** de clave-valor, como diccionarios en otros lenguajes.
- 📡 **Canales**: Son las **comunicaciones** que se usan para **goroutines**, base de la concurrencia en Go.
- **Strings**: Cada iteración entrega el **índice en bytes** y la **runa Unicode (tipo `rune`)**, **se comporta como colección de runas (caracteres Unicode)**

**Ejemplo**:

```go
var arr [3]int = [3]int{1, 2, 3} // - Arrays 🧮
slice := []int{1, 2, 3, 4} // - Slices 📦
m := map[string]int{"Ana": 25, "Luis": 30} // - Maps 🗺️
ch := make(chan int) // - Canales 📡
```

- **Strings**

```go
s := "hola"
for i, r := range s {
    fmt.Printf("Índice %d: Rune %c\n", i, r)
}
```

La sintaxis entre range de arrays y de slice es muy igual, pero recuerda que lo que los diferencia es cómo se declaran y usan:

| **Característica** | **Slice (`[]`)** | **Array (`[N]`)**         |
| -------------- | ------------ | --------------------- |
| Tamaño         | Dinámico     | Fijo (parte del tipo) |
| Tipo exacto    | `[]string`   | `[3]string`           |
| Uso            | Muy común    | Poco común            |

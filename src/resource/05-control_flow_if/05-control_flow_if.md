# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

---

#### Let's Get Started (Empecemos)

**Nota**: Nosotros podemos declarar variables dentro de un `if`.

#### Recomendaión

- Cuando las **condicionales son complejas**.

- Cuando **son simples y solo existen dos opciones (Si/No)**

- Cuando tenemos **condiciones que no estan relacionadas a una sola variables**.

```go
// Example 1
if user.IsAdmin && user.IsActive { // Se comparando dos o mas tipos de varaibles
    fmt.Println("Este usuario es valido")
}
// Example 2
temperature := 32
humidity := 40
if temperature > 25 && humidity < 50 {
    fmt.Println("PRENDA ESA JODA! de ventilador carajo!")
}
```

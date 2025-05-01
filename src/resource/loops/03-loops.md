# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

---

#### Let's get started (Empecemos)

- En **Go**, `for` es **el único bucle** disponible **que existe**. lo que significa que no existe cosas como *while* o **do while** deberas de aprender a adaptar el **for** paraque se comporte como ellos.
  
  - `for` la version clasica o tradicional.
  - Bucle estilo `while`
  - Bucle infinito

- `range` es una construcción que se **usa junto con `for`** para iterar sobre:
  
  - arrays
  - slices
  - maps
  - strings
  - canales (channels)

- Un **statement** son las instrucción completas que forma parte del flujo de control. No es un "comando", sino un bloque de ejecución completo como por ejemplo:
  
  - `if`, `else if`, `else` decisiones condicionales
  - `switch`: alternativa a múltiples `if`; más expresivo y seguro

- **Manejo de errores** en Go:
  
  - `error`: tipo de valor que representa un problema
  - `panic`: detiene la ejecución normal, se usa solo en situaciones graves
  - `recover`: función que permite capturar un `panic` y evitar que el programa termine abruptamente
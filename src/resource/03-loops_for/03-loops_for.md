# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

---

#### Let's get started (Empecemos)

En **Go**, **el único bucle** disponible **que existe** es `for`. lo que significa que no existe cosas como *while* o **do while** y otros. si los quieres, deberas de aprender a adaptar el **for** paraque se comporte como ellos y poder simularlos.

**Primero entendamos como se comporta cada uno de ellos, para entender como simularlos**

#### Loops - For Clasico

Sabemos que el `for` es una **estructura de control de flujo** que puede componerse de 4 partes: **inicialización, condición, incremento y cuerpo del bucle**.

- **Inicialización**: define y asigna la variable que controla el bucle.
- **Condición**: evalúa si se debe seguir ejecutando el bucle (una expresión que debe ser verdadera).
- **Incremento**: modifica la variable para acercarse al fin del ciclo.
- **Cuerpo del bucle**: es el bloque de código que se ejecuta en cada iteración mientras la condición se cumpla.

#### Loop - While Simulando
El `while` se conoce por ejecutarse múltiples veces mientras la condición sea **verdadera**. Es un tipo de bucle controlado por condición y se compone de 2 partes: condición y cuerpo del bucle.
**Si la condición es falsa desde el inicio, el bucle no se ejecuta en absoluto.**

- **Condición**: evalúa si se debe seguir ejecutando el bucle (una expresión que debe ser verdadera).
- **Cuerpo del bucle**: es el bloque de código que se ejecuta en cada iteración mientras la condición se mantenga verdadera.

#### Loop - Do While Simulando

El `do...while` se caracteriza por ejecutarse **al menos una vez**, sin importar si la condición es verdadera o falsa.
Esto ocurre porque el **cuerpo del bucle se ejecuta primero** y luego se **evalúa la condición**. Es, en esencia, el comportamiento inverso al de un bucle `while`.

- **Cuerpo del bucle**: es el bloque de código que se ejecuta en cada iteración mientras la condición se mantenga verdadera.
- **Condición**: evalúa si se debe seguir ejecutando el bucle (una expresión que debe ser verdadera).

#### Loop - Bucle infinito

Es solo otra forma **en codigo** de hacer el ``do while``

#### 📌 Clasificación Resumida

| **Estilo**           | **¿Se ejecuta al menos una vez?** | **¿Evalúa antes o después?** | **Condición visible** |
| -------------------- | ----------------------------- | ------------------------ | ----------------- |
| `while`              | No                            | Antes                    | Sí                |
| `do...while`         | Sí                            | Después                  | Sí                |
| `for` clásico        | Depende                       | Antes                    | Sí                |
| `for {}` con `break` | **Sí** (como este)            | **Después (implícita)**  | **No (manual)**   |
# **Go by Carlitos!**

## Bloque 1: Learn the Basics (Aprende lo basico)

### Sub-Bloque 3: programming tools 1 (herramientas de programación 1)

- For Loop
- range
- if, switch statements
- Errors, Panic, Recover

---

#### Let's Get Started (Empecemos)

##### Recommend

**Go** todos necesitamos saber manejar errores y prevenir situaciones hipoteticas. y go nos ayuda con **(Errors, Panic, Recover)**

- **Errors**: Siempre debemos recorda que **"los errores son valores"**, que las funciones devuelven **un `error` como segundo valor**, `if err != nil` se usa muchisimo en go. `errors` debe usarse para errores esperados o normales como por ejemplo (input inválido, archivo no existe).
- **Panic** :  se usa para situaciones **catastróficas**, el programa se detiene **inmediatamente, lo FRENA TODO!**, a menos que usemos `recover`, el panic debe de usarse cuando ocurrar **errores inesperados y criticos** 
- **Recover**: debe de usarse cuando queremos recuperarnos de un panic, para que el programa no se detenga abructamente y pueda **recobrar** el flujo de trabajo del programa.

Siempre que todo nos salga bien, debemos devolver un `nil` --> `err == nil` “Todo bien”.
Siempre que todo nos salga mal debemos deolver un `err` --> `err != nil` “Hubo un error”.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Este Panic ha sido recuperado: ", r)
        return
    }
}()
```

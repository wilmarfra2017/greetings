# Saludos en Go
Este paquete proporciona una forma simple de obtener saludos personales

## Instalación
Ejecuta el siguiente comando para instalar paquete:
```bash
go get -u github.com/wilmarfra2017/greetings
```
## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
    "fmt"
    "github.com/alexroel/greetings"
)

func main() {
    message, err := greetings.Hello("Wilmar")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}
```

Este ejemplo importa el paquete github.com/alexroel/greetings y llama a la función Hello para obtener un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
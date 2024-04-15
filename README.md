# go-poc

Api rest en Go

## Requisitos

- Go 1.22 o superior
- Docker (si se utiliza para la base de datos u otros servicios)


## Uso

### Ejecutar la aplicación

Para ejecutar la aplicación localmente, simplemente utiliza el comando `go run`:

```bash
go run main.go
```

### Ejecutar las pruebas

Para ejecutar las pruebas unitarias, utiliza el comando `go test`:

```bash
go test ./...
```

### Generar y utilizar mocks con mockgen

Para generar mocks utilizando `mockgen`, primero debes instalar la herramienta:

```bash
go install github.com/golang/mock/mockgen@latest
```

Luego, puedes generar un mock para una interfaz en tu código. Por ejemplo, supongamos que tienes una interfaz `UserService` en un paquete llamado `service`:

```bash
mockgen -destination=service/mock_service.go -package=service github.com/tu-usuario/tu-app/service UserService
```

Esto generará un archivo `mock_service.go` en el paquete `service` que contiene un mock para la interfaz `UserService`. Luego, puedes utilizar este mock en tus pruebas.

## Dependencias usadas

- [resty](https://pkg.go.dev/github.com/go-resty/resty/v2): cliente http utilizado para realizar peticiones a otros servicios.
- [testify](https://pkg.go.dev/github.com/getlantern/testify/assert): testify assersiones utilizadas en las pruebas unitarias.
- [gomock](https://pkg.go.dev/github.com/golang/mock/gomock) paquete utilizado para generar mocks en las pruebas unitarias.
- [httpmock](https://pkg.go.dev/github.com/jarcoal/httpmock): librería utilizada para mockear peticiones http en las pruebas unitarias.
- [httptest](https://pkg.go.dev/net/http/httptest): libreria utilizada para testear los handlers.

## Lecturas

- [making-and-using-http-middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- [prefer-table-driven-tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## Stack

- Go 1.22

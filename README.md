# go-poc

Api rest en Go

## Requisitos

- Go 1.22 o superior


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

Para generar mocks utilizando `mockgen`, primero hay que instalar la herramienta:

```bash
go install github.com/golang/mock/mockgen@latest
```

Luego, puedes generar un mock para una interfaz en tu código. Por ejemplo, supongamos que existe una interfaz `PokemonSrv` en un paquete llamado `domain`:

```bash
mockgen -source=internal/domain/pokemon.go -destination=mocks/pokemon.go
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
- [naming-test-cases](https://medium.com/getground/naming-tests-in-golang-c58c188bb9a1)

## Stack

- Go 1.22

## Arquitectura

La aqrquitectura de la app esta basada en parte en:

- https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc

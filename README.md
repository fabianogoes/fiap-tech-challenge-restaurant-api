# FIAP challenge food

> FIAP pós Software Architecture - Tech Challenge projeto de um Restaurante
> 
> [![CI/CD](https://github.com/fabianogoes/fiap-tech-challenge-restaurant-api/actions/workflows/deploy.yml/badge.svg)](https://github.com/fabianogoes/fiap-tech-challenge-restaurant-api/actions/workflows/deploy.yml)

Table of context
- [FIAP challenge food](#fiap-challenge-food)
  - [Project Architecture by Clean Architecture](#project-architecture-by-clean-architecture)
  - [Stack](#stack)
  - [Development](#development)
    - [Running](#running)
  - [Testing using Docker/Docker Compose](#testing-using-dockerdocker-compose)
    - [Pre-registered data](#pre-registered-data)
  - [Docker Commands](#docker-commands)
  - [Run Go test](#run-go-test)

---

## Project Architecture by Clean Architecture

- `app/web`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório ‘web’ contém o ponto de entrada principal a API REST.
- `domain/entities`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
- `domain/usecases`: diretório que contém Serviços de Domínio ou Use Cases.
- `domain/ports`: diretório que contém ‘interfaces’ ou contratos definidos que os adaptadores devem seguir.
- `adapters/payment`: adaptador para meio de pagamento externo.
- `adapters/delivery`: adaptador para meio de entrega externo.
- `frameworks/rest`: diretório que contém os controllers e manipulador de requisições REST.
- `frameworks/rest/dto`: diretório que contém objetos/modelo de request e response.
- `frameworks/repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
- `frameworks/repository/dbo`: diretório que contém objetos/entidades de banco de dados.
- `.infra`: diretório que contém arquivos de infrainstrutura
- `.infra/kubernetes`: diretório que contém os manifestos kubernetes
- `.infra/terraform`: diretório que contém os arquivos terraform para provisionar a infra do projeto

## Stack

- [x] [Go][0]
- [x] [Domain-Driven Design][6]
- [x] [Hexagonal Architecture][5]
- [x] [Gin Web Framework][1] - Routes, JSON validation, Error management, Middleware support
- [x] [PostgreSQL][3] - Database persistence
- [x] [GORM ORM library for Golang][2]
- [x] [Slog](https://pkg.go.dev/log/slog) - Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs. 
- [x] [GoDotEnv](https://github.com/joho/godotenv) - A Go (golang) port of dotenv project (which loads env vars from a .env file).
- [ ] [gin-swagger](https://github.com/swaggo/gin-swagger) - gin middleware to automatically generate RESTful API documentation with Swagger 2.0.
- [ ] [swag](https://github.com/swaggo/swag) - Swag converts Go annotations to Swagger Documentation 2.0
- [ ] [CORS gin's middleware](https://github.com/gin-contrib/cors) - Gin middleware/handler to enable CORS support.

## Development

Dependencies

- [Go Installation](https://go.dev/doc/install)

Check for go version 1.21.3

```shell
go version
```

Preparing app

```shell
git clone git@github.com:fabianogoes/fiap-tech-challenge-restaurant-api.git
cd fiap-tech-challenge-restaurant-api
go mod tidy
````

### Running

```shell
docker-compose up -d postgres && go run app/web/main.go
```

## Testing using Docker/Docker Compose

```shell
docker-compose up -d

curl --request GET --url http://localhost:8080/health

## response 
{"status":"UP"}
```

### Pre-registered data

Quando a ‘app’ subir será inserido dados necessários para testar a criação de pedidos 

Para verificar a **lista de produtos** pode ser usado a API:

```shell
http://localhost:8080/products
```

Para verificar a **lista de clientes** pode ser usado a API:

```shell
http://localhost:8080/customers
```

Para verificar a **lista de Atendentes** pode ser usado a API: 
```shell
http://localhost:8080/attendants
```

[Collection de Teste que pode ser importada no Postman](./__utils__/fiap-tech-challenge-Insomnia.json)

## Docker Commands

```shell
docker login -u=fabianogoes
docker build -t fabianogoes/restaurant-api:3.20240426.1 .
docker tag fabianogoes/restaurant-api:3.20240426.1 fabianogoes/restaurant-api:3.20240426.1
docker push fabianogoes/restaurant-api:3.20240426.1
```

## Run Go test

```shell
go test -v ./...
```

[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.postgresql.org/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP

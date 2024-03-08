# FIAP challenge food

> FIAP pós Software Architecture - Tech Challenge projeto de um Restaurante `gofood`

Table of context
- [FIAP challenge food](#fiap-challenge-food)
  - [Arquitetura do projeto](#project-architecture-by-clean-architecture)
  - [Stack](#stack)
  - [Para Desenvolver](#development)
  - [Para Testar a aplicação usando Docker/Docker Compose](#testing-using-dockerdocker-compose)
    - [Como testar usando o `curl`](#como-testar-usando-o-curl)
    - [Pode ser testado o fluxo completo usando a collection insomnia](#using-http-client-postman-or-insomnia)
  - [Docker Commands](#docker-commands)
  - [Referencias importantes](#referencias-importantes)

---

## Project Architecture by Clean Architecture

- `app/web`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório web contém o ponto de entrada principal a API REST.
- `domain/entities`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
- `domain/usecases`: diretório que contém Serviços de Domínio ou Use Cases.
- `domain/ports`: diretório que contém interfaces ou contratos definidos que os adaptadores devem seguir.
- `adapters/payment`: adaptador para meio de pagamento externo.
- `adapters/delivery`: adaptador para meio de entrega externo.
- `frameworks/rest`: diretório que contém os controllers e manipulador de requisições REST.
- `frameworks/rest/dto`: diretório que contém objetos/modelo de request e response.
- `frameworks/repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
- `frameworks/repository/dbo`: diretório que contém objetos/entidades de banco de dados.
- `k8s`: diretório com arquivos kubernetes

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

> Check for go version 1.21.3

```shell
go version
```

- `git clone https://github.com/fabianogoes/fiap-techchallenge-fase2.git`
- `cd fiap-techchallenge-fase2`
- `go mod tidy`

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

> Quando a app subir será inserido dados necessários para testar a criação de pedidos 

| Atentente ID  | Cliente CPF | Produto ID        |
|---------------|-------------|-------------------|
| 1             | 15204180001 | 1 (Big Lanche)    |
|               |             | 6 (Coca-Cola)     |
|               |             | 22 (Batata Frita) |

> - Para verificar a **lista de produtos** pode ser usado a API: `http://localhost:8080/products`
> - Para verificar a **lista de clientes** pode ser usado a API: `http://localhost:8080/customers`
> - Para verificar a **lista de Atendentes** pode ser usado a API: `http://localhost:8080/attendants`

### Como testar usando o `curl`

[Veja o documento](./__utils__/doc/entregavel-how-to-test-challenge.md)

### Using HTTP Client Postman or Insomnia

[Collection de Teste que pode ser importada no Postman](./__utils__/FIAP-GoFood.postman_collection.json)

## Docker Commands

```shell
docker login -u=fabianogoes
docker build -t fabianogoes/fiap-challenge:2.0 .
docker tag fabianogoes/fiap-challenge:2.0 fabianogoes/fiap-challenge:2.0
docker push fabianogoes/fiap-challenge:2.0
```

## Referencias importantes

- [Documento PDF entegável de como testar a API](./__utils__/doc/entregavel-how-to-test-challenge.pdf)
- [Documentação DDD Miro](https://miro.com/app/board/uXjVNpDpixg=/?share_link_id=459651604667)
- [Tech Challenge - Entregáveis fase 1](./__utils__/doc/EntragaFase1.md)
- [Como Testar usando `curl`](./__utils__/doc/ComoTestar.md)
- [Collection de Teste que pode ser importada no Insomnia ou Postman](./__utils__/Insomnia_collection_test.json)

[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.postgresql.org/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP

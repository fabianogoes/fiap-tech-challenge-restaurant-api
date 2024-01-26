# FIAP challenge food

> FIAP pós Software Architecture - Tech Challenge projeto de um Restaurante `gofood`


- [FIAP challenge food](#fiap-challenge-food)
  - [Arquitetura do projeto](#arquitetura-do-projeto)
  - [Stack](#stack)
  - [Para Desenvolver](#para-desenvolver)
  - [Para Testar a aplicação usando Docker/Docker Compose](#para-testar-a-aplicação-usando-dockerdocker-compose)
    - [Como testar usando o `curl`](#como-testar-usando-o-curl)
    - [Pode ser testado o fluxo completo usando a collection insomnia](#pode-ser-testado-o-fluxo-completo-usando-a-collection-insomnia)
  - [Referencias importantes](#referencias-importantes)

## Arquitetura do projeto

- `cmd`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório web contém o ponto de entrada principal a API REST.
- `internal`: diretório para conter o código do aplicativo que não deve ser exposto a pacotes externos.
  - `core`: diretório que contém a lógica de negócios central do aplicativo.
    - `domain`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
    - `port`: diretório que contém interfaces ou contratos definidos que os adaptadores devem seguir.
    - `service`: diretório que contém Serviços de Domínio ou Use Cases.
  - `adapters`: diretório para conter serviços externos que irão interagir com o core do aplicativo.
    - `handler`: diretório que contém os controllers e manipulador de requisições REST.
    - `handler\dto`: diretório que contém objetos/modelo de request e response.
    - `repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
    - `repository\dbo`: diretório que contém objetos/entidades de banco de dados.
    - `payment`: adaptador para meio de pagamento externo.
    - `delivery`: adaptador para meio de entrega externo.

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

## Para Desenvolver

Dependencias

- [Go Instalation](https://go.dev/doc/install)

> Certifique-se de ter Go 1.21 ou superior

```shell
go version
```

- Clonar o repostório
- Entrar na pasta e rodar o comando para baixar as dependências `go mod tidy`
- Fazer uma cópia do arquivo .env.example e renomear para .env `cp .env.example .env`

Para Rodar o projeto em development

```shell
docker-compose up -d postgres && go run cmd/web/main.go
```

## Para Testar a aplicação usando Docker/Docker Compose

```shell
docker-compose up -d

curl --request GET --url http://localhost:8080/health

## resposta esperada
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

### Pode ser testado o fluxo completo usando a collection insomnia

[Collection de Teste que pode ser importada no Insomnia ou Postman](./Insomnia_collection_test.json)

## Referencias importantes

- [Documento PDF entegável de como testar a API](./__utils__/doc/entregavel-how-to-test-challenge.pdf)
- [Documentação DDD Miro](https://miro.com/app/board/uXjVN8Gnn2s=/)
- [Tech Challenge - Entregáveis fase 1](./doc/EntragaFase1.md)
- [Como Testar usando `curl`](./__utils__/doc/ComoTestar.md)
- [Collection de Teste que pode ser importada no Insomnia ou Postman](./Insomnia_collection_test.json)

[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.postgresql.org/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP

# FIAP Tech Challenge - Restaurant

[![CI](https://github.com/fabianogoes/fiap-tech-challenge-restaurant-api/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/fabianogoes/fiap-tech-challenge-restaurant-api/actions/workflows/ci-cd.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fabianogoes_fiap-tech-challenge-restaurant-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=fabianogoes_fiap-tech-challenge-restaurant-api)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=fabianogoes_fiap-tech-challenge-restaurant-api&metric=coverage)](https://sonarcloud.io/summary/new_code?id=fabianogoes_fiap-tech-challenge-restaurant-api)
<img src="https://sonarcloud.io/images/project_badges/sonarcloud-white.svg" alt="Scanned on SonarCloud" height="20px" />


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
- [x] [Criptografia Simétrica AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

## Desafios

- Implementação da Arquitetura Hexagonal
- Implementação de Conteinerização usando Docker e Docker Compose
- Implementação de Orquestração de Containers usando Kubernetes e AWS EKS
- Implementação de CI/CD usando Github Actions
- Implementação da Arquitetura de Microservices com comunicação Assincrôna usando mensageria
- Implementação de API Gateway AWS Gateway com Autenticação e Autorização usando AWS Lambda
- Implementação de Autenticação usando JWT
- Implementação de qualidade de código usando SonarLint, SonarQube e SonarCloud
- Implementação de Transações distruibuidas usando o Padrão SAGA  
- Distribuição de processos usando AWS Lambda
- Uso de IaaS(Infrastructure as a Service) usando Terraform

## Motivações

Boas práticas e padrões usados para resolver os desafios

### Padrão SAGA

Na implementação do padrão SAGA optei por usar a estratégia de Coreografia para fazer uso dos seguintes benefícios:

1. **Desacoplamento e Autonomia dos Serviços**
   Para que os serviços fiquem mais independentes uns dos outros e não dependam de um orquestrador central o que adicionaria um outro ponto de falha. A ideia é que cada serviço conhece apenas sua própria lógica e como reagir a eventos específicos, permitindo maior autonomia no desenvolvimento e evolução dos serviços.

2. **Escalabilidade**
    Como não há um ponto central de controle, o sistema pode escalar melhor horizontalmente, já que o aumento na carga de trabalho passa ser distribuído entre os serviços, sem sobrecarregar um único orquestrador.

3. **Resiliência**
    A falha de um serviço não impede necessariamente que outros serviços continuem a operar. Cada serviço foi projetado para lidar com falhas de maneira mais isolada, aumentando a resiliência geral do sistema.

4. **Flexibilidade e Evolução do Sistema**
   Adicionar, modificar ou remover serviços é mais fácil e menos impactante, pois não há necessidade de alterar um orquestrador central. Isso tornou a arquitetura mais adaptável a mudanças de requisitos de negócios ou novas funcionalidades. Cada serviço pode ser desenvolvido e implantado de forma independente, o que facilita a adoção de novas tecnologias ou padrões sem impactar o sistema como um todo.

5. **Melhor Alinhamento com Arquiteturas Orientadas a Eventos**
    A coreografia alinhou bem com arquiteturas orientadas a eventos, onde os eventos dirigem o fluxo das operações, facilitando a implementação de arquiteturas reativas e altamente responsivas.


## Setup Development

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

[Collection](./.utils/fiap-tech-challenge-Insomnia.json)

## Docker Commands

```shell
docker login -u=fabianogoes
docker build -t fabianogoes/restaurant-api:latest .
docker tag fabianogoes/restaurant-api:3.20240426.1 fabianogoes/restaurant-api:latest
docker push fabianogoes/restaurant-api:latest
```

## Run Coverage

```shell
clear && go test -coverprofile=coverage.out ./... &&  go tool cover -func=coverage.out
```

[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.postgresql.org/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP

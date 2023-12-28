# FIAP challenge food

> FIAP pós Software Architecture - Tech Challenge projeto Food

## Arquitetura do projeto

![Hexagonal Structure](./assets/hexagonal-structure.png)

- `cmd`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório web contém o ponto de entrada principal a API REST.
- `internal`: diretório para conter o código do aplicativo que não deve ser exposto a pacotes externos.
- `core`: diretório que contém a lógica de negócios central do aplicativo.
  - `domain`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
  - `port`: diretório que contém interfaces ou contratos definidos que os adaptadores devem seguir.
  - `service`: diretório que contém Serviços de Domínio ou Use Cases.
- `adapters`: diretório para conter serviços externos que irão interagir com o core do aplicativo
  - `handler`: diretório que contém os controllers e manipulador de requisições REST.
  - `repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
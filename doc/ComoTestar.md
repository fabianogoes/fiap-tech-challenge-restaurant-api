# Como testar a aplicação

## Dados pré cadastrados para testes

### Atendentes

| ID    | Nome      |
|-------|-----------|
| 1     | Miguel    |
| 2     | Sophia0   |
| 3     | Alice     |
| 4     | Pedro     |
| 5     | Manuela   |

### Clientes

| ID    | Nome          | Email                     | CPF           |
|-------|---------------|---------------------------|---------------|
| 1     | Bernardo      | bernardo@gmail.com        | 29381510040   |
| 2     | Laura         | laura@hotmail.com         | 15204180001   |
| 3     | Lucas         | lucas@gmail.com           | 43300921074   |
| 4     | Maria Eduarda | meduarda@uol.com.br       | 85752055016   |
| 5     | Guilherme     | guilherme@microsoft.com	| 17148604001   |

### Categorias de Produtos

| ID    | Nome              |
|-------|-------------------|
| 1     | Sanduíches        |
| 2     | Bebidas Frias     |
| 3     | Bebidas Quentes   |
| 4     | Combos            |
| 5     | Sobremesas        |
| 6     | Acompanhamentos   |
| 7     | Café da Manhã     |

 | ID   | Nome                              | Preço     | Category_id   |
 |------|-----------------------------------|-----------|---------------|
 | 1    | Big Lanche                        | 29.9      | 1             |
 | 2    | X-Burguer                         | 19.9      | 1             |
 | 3    | X-Salada                          | 21.9      | 1             |
 | 4    | X-Bacon                           | 23.9      | 1             |
 | 5    | X-Tudo                            | 27.9      | 1             |
 | 6    | Coca-Cola                         | 5.9       | 2             |
 | 7    | Guaraná                           | 5.9       | 2             |
 | 8    | Fanta                             | 5.9       | 2             |
 | 9    | Suco de Laranja                   | 5.9       | 2             |
 | 10   | Suco de Uva                       | 5.9       | 2             |
 | 11   | Café                              | 3.9       | 3             |
 | 12   | Cappuccino                        | 4.9       | 3             |
 | 13   | Chocolate Quente                  | 4.9       | 3             |
 | 14   | Misto Quente + Fritas             | 9.9       | 4             |
 | 15   | X-Burguer + Fritas + Coca-Cola    | 29.9      | 4             |
 | 16   | X-Salada + Fritas + Coca-Cola     | 31.9      | 4             |
 | 17   | X-Bacon + Fritas + Coca-Cola      | 33.9      | 4             |
 | 18   | X-Tudo + Fritas + Coca-Cola       | 37.9      | 4             |
 | 19   | Sorvete                           | 7.9       | 5             |

## Teste API de Atendentes usando o curl

Criação de um novo Atendente

```shell
curl --request POST \
  --url http://localhost:8080/attendants \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Attendant 1" }'
```

Busca de Atendente por ID

```shell
curl --request GET \
  --url http://localhost:8080/attendants/1 \
```

Lista de Atendentes

```shell
curl --request GET \
  --url http://localhost:8080/attendants
```

Atualização de Atendente

```shell
curl --request PUT \
  --url http://localhost:8080/attendants/1 \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Attendant Updated" }'
```

Exclusão de Atendente

```shell
curl --request DELETE \
  --url http://localhost:8080/attendants/1
```

## Teste API de Clientes usando o curl

Criação de um novo Cliente

```shell
curl --request POST \
  --url http://localhost:8080/customers \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Customer 2", "email": "customer2@test.com", "cpf": "12345678902" }'
```

Busca de Cliente por ID

```shell
curl --request GET \
  --url http://localhost:8080/customers/1
```

Busca de Cliente por CPF

```shell
curl --request GET \
  --url http://localhost:8080/customers/cpf/15204180001
```

Lista de Clientes

```shell
curl --request GET \
  --url http://localhost:8080/customers
```

Atualização de Cliente

```shell
curl --request PUT \
  --url http://localhost:8080/customers/1 \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Customer Updated", "email": "customer.updated@test.com" }'
```

Exclusão de Cliente

```shell
curl --request DELETE \
  --url http://localhost:8080/customers/1
```

## Teste API de Produtos usando o curl

Criação de novo Produto

```shell
curl --request POST \
  --url http://localhost:8080/products \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Product 2", "price": 55.55, "categoryID": 2 }'
```

Busca de Produto por ID

```shell
curl --request GET \
  --url http://localhost:8080/products/1
```

Lista de Produtos

```shell
curl --request GET \
  --url http://localhost:8080/products
```

Atualização de Produto

```shell
curl --request PUT \
  --url http://localhost:8080/products/1 \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Product updated", "price": 10.55, "categoryID": 1}'
```

Exclusão de Produto

```shell
curl --request DELETE \
  --url http://localhost:8080/products/1
```

## Teste API de Orders usando o curl

Start Order

```shell
curl --request POST \
  --url http://localhost:8080/orders/start \
  --header 'Content-Type: application/json' \
  --data '{ "customerCPF": "15204180001", "attendantID": 1 }'
```

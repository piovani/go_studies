go_api

# Objetivos
 - [ ] CRUD de Cliente
    - Nome
    - CPF
    - EMAIL
    - Telefone
    - Endereço
 - [ ] CRUD de Produto
    - Nome
    - Preço
    - Peso
    - Imagem
    - Descrição
    - Estoque
    - Status
 - [ ] CRUD of Pedido
    - Codigo do Pedido
    - Cliente
    - Status
    - Subtotal dos Produtos
    - Valor do Frete
    - Total
    - Data de Criação
    - Forma de Pagamento
    - Forma de Entrega
 - [ ] CRUD do Item do Pedido
    - Codigo do Pedido
    - COdigo do Produto
    - Nome do Produto
    - Quantidade
    - Oreço Unitário
    - Total
 - [ ] Autenticação JWT
 - [ ] Apenas Produtos com Status ativo podem ser add ao Pedido
 - [ ] Apenas Produtos com Estoque disponivel podem ser add ao Pedido
 - [ ] Ao add um novo Produto ao Pedido o Total do Pedido deve ser atualizado
 - [ ] Apenas uma forma de entrega, o valor do frete é calculado da seguinte forma: Cada KG é add R$ 5,00, o valor é arredodado para cima.
 - [ ] Pagamento a vista tem 5% de desconto
 - [ ] Pagamento a prazo deve ser informado o numero de parcelas
 - [ ] Rota para acessar todos os pedidos realizados pelo cliente.
 - [ ] Teste Unitario para os Services do Sistema
 - [ ] Teste Automatizado para todas as Rotas do Sistema
 - [ ] Utilizar GraphQL na aplicação


 Obs: Deixar o Go executando: ``` CompileDaemon ```

 # Set up the project

 ## First copy the environment variables file

 ``` 
 cp .env.example .env 
 ```

 ## Upload the docker containers on your machine

 ```
 docker-compose up -d
 ```

 <b>Note:</b> 
 1) the docker and docker-compose are required to be installed on your machine
 2) versions used:
      * Docker: 19.03.8
      * Docker-compose: 1.25.0
 3) try to use the same version or higher

 # How to teste

## In the project repository we have the collection of Postman requisitions, just import it into your Postman that will have all the requisition routes ready

<b>file name:</b> `API.postman_collection.json`


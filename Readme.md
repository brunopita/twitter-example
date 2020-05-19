
# Twitter Example

## Objetivo

## Estrutura do projeto

Projeto foi desenvolvido utizando [Golang](https://golang.org/), utilizando as ferramentas da Elastic para coleta de logs e monitoramento da aplicação.

O projeto é composto por 3 aplicações:

- twitter-migrate: Responsável por criar as tabelas do banco de dados, indices e realizar alterações estruturais caso necessário; 
- twitter-consumer: Responsável por coletar os dados da API do Twitter e armazena-las no Postgres; 
- twitter-api: API Rest para consumo das informações;

## Ambiente 

Pré requisitos:

- [Docker](https://docs.docker.com/desktop/#download-and-install)
- [docker-compose](https://docs.docker.com/compose/install/#install-compose)

Execução: 

 - docker-compose up --build

Coleção do [Postman](www.postman.com) para teste - [Link](https://github.com/brunopita/twitter-example/blob/master/docs/postman/Twitter-api.postman_collection.json)
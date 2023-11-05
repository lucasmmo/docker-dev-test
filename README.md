# go-microservice-boilerplate

Este é um breve guia sobre como rodar o projeto usando os comandos definidos no Makefile.

## Pré-requisitos

Certifique-se de ter o seguinte instalado em sua máquina:

- [Golang](https://go.dev/) - Certifique-se de ter o Golang instalado para conseguir utilizar suas ferramentas.
- [Docker](https://www.docker.com/) - Para rodar os ambientes locais com o Docker Compose.
- [Make](https://www.gnu.org/software/make/) - O utilitário `make` é usado para automatizar os comandos do projeto.

## Configuração Inicial

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto no terminal.

## Como rodar o projeto

O projeto possui os seguintes comandos Makefile:

### 1. `up`

Este instalará as dependencias e inicia os ambientes locais usando o Docker Compose.

```bash
make up
```

Isso fará o seguinte:

- Instalará as dependências locais usando o go mod tidy.
- Inicializará os ambientes locais usando o Docker Compose.

### 2. `down`

Este comando interrompe os ambientes locais usando o Docker Compose e também apagará a pasta tmp.

```bash
make down
```
Isso fará o seguinte:

- Apagará a pasta tmp.
- Parará os ambientes locais usando o Docker Compose.

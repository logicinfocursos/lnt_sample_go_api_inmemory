# API de Filmes em Go (CRUD em memória)

## Requisitos
- Go instalado (https://go.dev/dl/)

## Passo a passo

### 1. Instalar o Go
- Baixe e instale o Go pelo site oficial: https://go.dev/dl/
- Verifique a instalação:
  ```bash
  go version
  ```

### 2. Instalar as dependências
- No diretório do projeto, execute:
  ```bash
  go mod init api_inmemory
  go get github.com/gin-gonic/gin
  go get github.com/joho/godotenv
  ```

O comando go mod init api_inmemory serve para inicializar um novo módulo Go no diretório do projeto. Ele cria o arquivo go.mod, que gerencia as dependências do projeto, permitindo que você use pacotes externos e controle versões de forma organizada.

Iniciar um novo módulo em Go significa criar um projeto independente, com seu próprio gerenciamento de dependências. O módulo permite que você controle quais bibliotecas externas o projeto usa, facilita o compartilhamento do código e garante que ele funcione corretamente em qualquer ambiente, pois todas as dependências ficam registradas no arquivo go.mod. Isso torna o projeto organizado, reprodutível e pronto para evoluir ou ser distribuído.

### 3. Configurar variáveis de ambiente
- Edite o arquivo `.env` para definir a porta da API:
  ```env
  API_PORT=8080
  ```

### 4. Executar o projeto
- No diretório do projeto, rode:
  ```bash
  go run main.go
  ```
- Acesse a API em: http://localhost:8080

## Endpoints
- `GET /movies` - Lista todos os filmes
- `GET /movies/:id` - Busca filme por ID
- `POST /movies` - Cria novo filme
- `PUT /movies/:id` - Atualiza filme
- `DELETE /movies/:id` - Remove filme

## Observações
- Os dados são salvos em `movies.json`.
- O projeto não utiliza banco de dados, apenas arquivo local.

# API de Filmes em Go (CRUD em memória)

repositório do projeto: https://github.com/logicinfocursos/lnt_sample_go_api_inmemory.git

Esse projeto demonstra como criar uma API de filmes em Go usando um CRUD em memória, gerar um container Docker e um cluster Kubernetes usando Kind. 

## Requisitos
- Go instalado ([download](https://go.dev/dl/))
- Docker e Docker Compose ([download](https://docs.docker.com/get-docker/))
- Kind ([guia de instalação](https://kind.sigs.k8s.io/docs/user/quick-start/#installation))
- kubectl ([guia de instalação](https://kubernetes.io/docs/tasks/tools/))

### 1. Executar o projeto em GO (Golang)
Verifique a instalação:
```
go version
```

#### 1.1. Instalar as dependências
No diretório do projeto, execute:
```
go mod init api_inmemory
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
```

O comando go mod init api_inmemory serve para inicializar um novo módulo Go no diretório do projeto. Ele cria o arquivo go.mod, que gerencia as dependências do projeto, permitindo que você use pacotes externos e controle versões de forma organizada.

Iniciar um novo módulo em Go significa criar um projeto independente, com seu próprio gerenciamento de dependências. O módulo permite que você controle quais bibliotecas externas o projeto usa, facilita o compartilhamento do código e garante que ele funcione corretamente em qualquer ambiente, pois todas as dependências ficam registradas no arquivo go.mod. Isso torna o projeto organizado, reprodutível e pronto para evoluir ou ser distribuído.

#### 1.2. Configurar variáveis de ambiente
Edite o arquivo .env para definir a porta da API:
```
API_PORT=8080
```

#### 1.3. Executar o projeto
No diretório do projeto, rode:
```
go run main.go
```
Acesse a API em: http://localhost:8080 ou na porta definida no .env (exemplo: API_PORT=8091)

### 1.4. Criar o arquivo executável
Para gerar o arquivo .exe (executável para Windows) do seu projeto Go, execute no terminal:
```
go build -o api_inmemory.exe main.go
```
Isso criará o arquivo api_inmemory.exe no diretório atual. Você pode rodar o executável diretamente no Windows com:
```
./api_inmemory.exe
```
Se quiser gerar o .exe em outro sistema (Linux/Mac), use:
```
GOOS=windows GOARCH=amd64 go build -o api_inmemory.exe main.go
```

#### 1.4. Endpoints da api:
- GET /movies - Lista todos os filmes
- GET /movies/:id - Busca filme por ID
- POST /movies - Cria novo filme
- PUT /movies/:id - Atualiza filme
- DELETE /movies/:id - Remove filme

**Observações**
- Os dados são salvos em movies.json.
- O projeto não utiliza banco de dados, apenas arquivo local.
- tem um outro repositório de exemplo usando o GORM e mysql: https://github.com/logicinfocursos/lnt_sample_go_api.git
- tem um outro repoistório com um app web para exibir os dados da api:
https://github.com/logicinfocursos/lnt_sample_go_appweb.git
- se você quiser aprender várias linguagens de programação através de analogias com os conhecimentos que você já tem: https://github.com/logicinfocursos/learning_new_techs.git e veja o meu blog com um post completo sobre o tema: https://automaticlab.com.br/posts/aprenda-por-analogia


### 2. Usar Docker
Seguem as instruções para build e uso do container com Docker e Docker Compose: 
Build da imagem:
```bash
docker build -t logicinfocursos/api_inmemory:latest .
```
Executar container:
```bash
docker run -p 8091:8091 logicinfocursos/api_inmemory:latest
```
Acesse: http://localhost:8091

#### 2.1. Usar Docker Compose
```bash
docker-compose up --build
```
Para parar e remover:
```bash
docker-compose down --rmi all
```

#### 2.2. Subir imagem para Docker Hub
```bash
docker login
docker push logicinfocursos/api_inmemory:latest
```

### 3. Usar no Kubernetes
#### 3.1. Criar cluster local com Kind
```bash
kind create cluster --name meu-cluster
kubectl cluster-info --context kind-meu-cluster
kubectl get nodes
``` 
Criar cluster local com Kind
```bash
kind create cluster --name meu-cluster
kubectl cluster-info --context kind-meu-cluster
kubectl get nodes
```

### 3.2. Deploy no Kubernetes
Edite o arquivo `deployment.yaml` para usar sua imagem do Docker Hub:
```yaml
image: logicinfocursos/api_inmemory:latest
```
Aplicar os manifests:
```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
Acesse a API via NodePort: http://localhost:30091

#### 3.3. Remover cluster Kind
```bash
kind delete cluster --name meu-cluster
```

### 4. Instruções para build e uso do container com Docker e Docker Compose:

Certifique-se de ter o Docker e o Docker Compose instalados.
Para buildar e subir o container usando o docker-compose:
```
docker-compose up --build
```

O serviço será iniciado com o nome e porta definidos no .env (exemplo: API_PORT=8091).

Para acessar a API, use:
```
http://localhost:8091
```

Para parar e remover o container:
```
docker-compose down
```

Se quiser buildar apenas a imagem (sem subir o container):
```
docker-compose build
```

Apagar a imagem gerada:
```
docker rmi nome_da_imagem
```

Se quiser remover tudo de uma vez (container, rede e imagem):
```
docker-compose down --rmi all
```

### 5. Instruções para criar um cluster local com Kind e gerenciar com kubect:
Instale o Kind:
```
https://kind.sigs.k8s.io/docs/user/quick-start/#installation
```

Instale o kubectl:
```
https://kubernetes.io/docs/tasks/tools/
```

Crie o cluster com Kind:
```
kind create cluster --name meu-cluster
```

Verifique se o cluster está ativo:
```
kubectl cluster-info --context kind-meu-cluster
```

Para listar os nós:
```
kubectl cluster-info --context kind-meu-cluster
```

Para deletar o cluster:
```
kind delete cluster --name meu-cluster
```
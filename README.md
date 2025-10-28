# API de Filmes em Go (CRUD em memória)

repositório do projeto: https://github.com/logicinfocursos/lnt_sample_go_api_inmemory.git

Esse projeto demonstra como criar uma API de filmes em Go usando um CRUD em memória, gerar um container Docker e um cluster Kubernetes usando Kind. 

## Requisitos
- Go instalado ([download](https://go.dev/dl/))
- Docker e Docker Compose ([download](https://docs.docker.com/get-docker/))
- Kind ([guia de instalação](https://kind.sigs.k8s.io/docs/user/quick-start/#installation))
- kubectl ([guia de instalação](https://kubernetes.io/docs/tasks/tools/))

## Passo a passo

### 1. Executar localmente com Go
```bash
go mod init api_inmemory
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go run main.go
```
Acesse: http://localhost:8091


## Endpoints da API
- `GET /movies` - Lista todos os filmes
- `GET /movies/:id` - Busca filme por ID
- `POST /movies` - Cria novo filme
- `PUT /movies/:id` - Atualiza filme
- `DELETE /movies/:id` - Remove filme

## Observações
- Os dados são salvos em `movies.json`.
- O projeto não utiliza banco de dados, apenas arquivo local.
- O container e o cluster podem ser acessados livremente pela porta definida no `.env`.


### 2. Usar Docker
Build da imagem:
```bash
docker build -t logicinfocursos/api_inmemory:latest .
```
Executar container:
```bash
docker run -p 8091:8091 logicinfocursos/api_inmemory:latest
```
Acesse: http://localhost:8091

### 3. Usar Docker Compose
```bash
docker-compose up --build
```
Para parar e remover:
```bash
docker-compose down --rmi all
```

### 4. Subir imagem para Docker Hub
```bash
docker login
docker push logicinfocursos/api_inmemory:latest
```

### 5. Criar cluster local com Kind
```bash
kind create cluster --name meu-cluster
kubectl cluster-info --context kind-meu-cluster
kubectl get nodes
```

### 6. Deploy no Kubernetes
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

### 7. Remover cluster Kind
```bash
kind delete cluster --name meu-cluster
```

### 8. Instruções para build e uso do container com Docker e Docker Compose:

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

### 9. Instruções para criar um cluster local com Kind e gerenciar com kubect:
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
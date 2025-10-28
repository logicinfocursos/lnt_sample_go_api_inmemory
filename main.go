// 0. Importar as dependências
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Estrutura para representar um filme
// Movie representa um filme no sistema
// As tags json permitem o mapeamento correto dos campos
// ao serializar/deserializar
// Exemplo: {"id":1, "title":"Filme", "year":2020}
// Estrutura Movie compatível com movies.json
type Movie struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Overview  string `json:"overview"`
	PosterURL string `json:"posterurl"`
	Genres    string `json:"genres"`
}

var movies []Movie // Slice para armazenar os filmes em memória

// 1. Carregar as variáveis do .env
func loadEnv() {
	_ = godotenv.Load()
}

// 2. Definir a porta da API (obter a porta do .env)
func getPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Porta padrão caso não definida
	}
	return port
}

// Função para carregar os filmes do arquivo JSON
func loadMovies() {
	f, err := os.Open("movies.json")
	if err != nil {
		movies = []Movie{}
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		movies = []Movie{}
		return
	}
	_ = json.Unmarshal(data, &movies)
}

// Função para salvar os filmes no arquivo JSON
func saveMovies() {
	data, _ := json.MarshalIndent(movies, "", "  ")
	os.WriteFile("movies.json", data, 0644)
}

func main() {
	loadEnv()         // Carrega variáveis do .env
	port := getPort() // Obtém a porta da API
	loadMovies()      // Carrega os filmes do arquivo

	// 3. Inicializar o framework e estrutura de dados
	// gin.Default() já inclui logger e recovery middleware
	r := gin.Default()

	// Rota de instrução e status da API
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensagem":   "API de exemplo em Go está funcionando!",
			"instrucoes": "Use /movies (GET) para listar, /movies/{id} (GET) para buscar, /movies (POST) para criar, /movies/{id} (PUT) para atualizar, /movies/{id} (DELETE) para remover.",
		})
	})

	// Descobre o IP local da máquina
	ip := "localhost"
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	fmt.Println("Api de exemplo em Go sendo executado em:", ip+":"+port)

	// 4. Configurar os Middlewares (CORS, etc)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Permite acesso de qualquer origem
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 5. Definir as rotas da API para CRUD
	// GET /movies - lista todos os filmes
	r.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, movies)
	})

	// GET /movies/:id - busca um filme por id
	r.GET("/movies/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, m := range movies {
			if m.ID == id {
				c.JSON(http.StatusOK, m)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
	})

	// POST /movies - cria um novo filme
	r.POST("/movies", func(c *gin.Context) {
		var newMovie Movie
		if err := c.ShouldBindJSON(&newMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}
		// Gera um novo ID
		maxID := 0
		for _, m := range movies {
			if m.ID > maxID {
				maxID = m.ID
			}
		}
		newMovie.ID = maxID + 1
		movies = append(movies, newMovie)
		saveMovies() // Salva no arquivo
		c.JSON(http.StatusCreated, newMovie)
	})

	// PUT /movies/:id - atualiza um filme existente
	r.PUT("/movies/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updateMovie Movie
		if err := c.ShouldBindJSON(&updateMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}
		for i, m := range movies {
			if m.ID == id {
				updateMovie.ID = id // Garante que o ID não seja alterado
				movies[i] = updateMovie
				saveMovies()
				c.JSON(http.StatusOK, updateMovie)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
	})

	// DELETE /movies/:id - remove um filme
	r.DELETE("/movies/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, m := range movies {
			if m.ID == id {
				movies = append(movies[:i], movies[i+1:]...)
				saveMovies()
				c.JSON(http.StatusOK, gin.H{"message": "Filme removido"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
	})

	// 6. Inicializar o servidor na porta definida (API_PORT)
	r.Run(":" + port)
}

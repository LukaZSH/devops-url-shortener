package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

var (
	redisClient *redis.Client
	amqpChannel *amqp.Channel
	ctx         = context.Background()
)

// Tipo um try-catch do java
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type ShortenRequest struct {
	URL string `json:"url"`
}
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234S9"
	const codeLength = 6
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Handler para a rota de encurtamento (/shorten)
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()
	err := redisClient.Set(ctx, shortCode, req.URL, 24*time.Hour).Err()
	if err != nil {
		http.Error(w, "Não foi possível salvar a URL", http.StatusInternalServerError)
		return
	}

	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "http://localhost:3000"
	}
	
	resp := ShortenResponse{
		ShortURL: fmt.Sprintf("%s/%s", host, shortCode),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Erro ao codificar resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}
}

// Handler para a rota de redirecionamento (/)
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	if shortCode == "" || shortCode == "shorten" { // Ignora a rota da API
		http.NotFound(w, r)
		return
	}

	originalURL, err := redisClient.Get(ctx, shortCode).Result()
	if err == redis.Nil {
		http.Error(w, "URL não encontrada", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Erro no banco de dados", http.StatusInternalServerError)
		return
	}

	// Publica a mensagem de forma assíncrona. Se falhar, não impede o redirecionamento.
	go publishClickEvent(shortCode)

	http.Redirect(w, r, originalURL, http.StatusFound)
}

// NOVA FUNÇÃO PARA PUBLICAR NO RABBITMQ
func publishClickEvent(shortCode string) {
	err := amqpChannel.Publish(
		"",       // exchange (sem exchange)
		"clicks", // routing key (o nome da fila)
		false,    // mandatory (não obrigatório)
		false,    // immediate (não imediato)
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(shortCode),
		})
	if err != nil {
		log.Printf("Falha ao publicar mensagem para o código %s: %s", shortCode, err)
	} else {
		log.Printf("Mensagem para o código %s enviada para a fila de clicks.", shortCode)
	}
}

func main() {
	// Conecta ao Redis
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisClient = redis.NewClient(&redis.Options{Addr: redisAddr})

	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}
	amqpConn, err := amqp.Dial(amqpURL)
	failOnError(err, "Falha ao conectar ao RabbitMQ")
	// defer amqpConn.Close() // Não fecha a conexão aqui para mantê-la viva

	amqpChannel, err = amqpConn.Channel()
	failOnError(err, "Falha ao abrir um canal")
	// defer amqpChannel.Close() // Não fecha o canal

	// Declara a fila para garantir que ela exista
	_, err = amqpChannel.QueueDeclare(
		"clicks", true, false, false, false, nil,
	)
	failOnError(err, "Falha ao declarar a fila")

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("API Gateway rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp" 
)

// Tipo um try-catch do java
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Conexão com o RabbitMQ
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(amqpURL)
	failOnError(err, "Falha ao conectar ao RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Falha ao abrir um canal")
	defer ch.Close()

	// Declara a fila 'clicks'. Se não existir, será criada.
	q, err := ch.QueueDeclare(
		"clicks", // nome da fila
		true,     // durable (persiste após reinicialização)
		false,    // delete when unused (deleta quando não utilizado)
		false,    // exclusive (outras conexões podem usar)
		false,    // no-wait (espera pela confirmação)
		nil,      // arguments (sem argumentos)
	)
	failOnError(err, "Falha ao declarar a fila")

	// Conexão com o PostgreSQL
	pgConnStr := os.Getenv("POSTGRES_CONN_STR")
	if pgConnStr == "" {
		pgConnStr = "postgres://postgres:postgres@localhost:5432/analytics?sslmode=disable"
	}
	db, err := sql.Open("postgres", pgConnStr)
	failOnError(err, "Falha ao conectar ao PostgreSQL")
	defer db.Close()

	// Cria a tabela de clicks se ela não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clicks (
		id SERIAL PRIMARY KEY,
		short_code VARCHAR(10) NOT NULL,
		clicked_at TIMESTAMP NOT NULL
	);`)
	failOnError(err, "Falha ao criar a tabela 'clicks'")
	
	// Prepara para consumir mensagens da fila
	msgs, err := ch.Consume(
		q.Name, // queue (nome da fila)
		"",     // consumer (sem nome específico)
		true,   // auto-ack (confirma o recebimento automaticamente)
		false,  // exclusive (outras conexões podem usar)
		false,  // no-local (não recebe mensagens de si mesmo)
		false,  // no-wait (não espera por confirmação)
		nil,    // args (sem argumentos)
	)
	failOnError(err, "Falha ao registrar um consumidor")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			shortCode := string(d.Body)
			log.Printf("Recebida mensagem: %s", shortCode)

			// Insere o dado no banco de dados PostgreSQL
			_, err := db.ExecContext(context.Background(),
				"INSERT INTO clicks (short_code, clicked_at) VALUES ($1, $2)",
				shortCode, time.Now())

			if err != nil {
				log.Printf("Falha ao inserir click no DB: %s", err)
			} else {
				log.Printf("Click do código '%s' registrado com sucesso!", shortCode)
			}
		}
	}()

	log.Printf(" [*] Aguardando por mensagens de cliques. Para sair, pressione CTRL+C")
	<-forever
}
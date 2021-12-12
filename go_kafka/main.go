package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql" // Driver
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type Env struct {
	ApiPort      string
	ConnectionDB string
	KafkaBroker  string
	KafkaGroup   string
	KafkaTopic   string
}

var env Env

func main() {
	initConfig()
	getMessages()

	// e := echo.New()

	// e.GET("/", list)
	// e.GET("/receive", receive)
	// e.POST("/send", send)

	// e.Logger.Fatal(e.Start(env.ApiPort))
}

func logError(err error) {
	fmt.Println("DEU RUIM, erro: %v", err)
}

func initConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		logError(err)
	}

	env.ApiPort = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	env.ConnectionDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	env.KafkaBroker = os.Getenv("KAFKA_BROKER")
	env.KafkaGroup = os.Getenv("KAFKA_GROUP")
	env.KafkaTopic = os.Getenv("KAFKA_TOPIC")
}

func list(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func receive(c echo.Context) error {
	if !getMessages() {
		return c.String(http.StatusInternalServerError, "DEU RUIM")
	}

	return c.String(http.StatusOK, "DEU BOM")
}

func send(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!2")
}

//======================================================================================================================

type MessageKafka struct {
	Message string `json:"message"`
}

func getMessages() bool {
	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.ChannelBufferSize = 2000

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sarama.NewConsumerGroup(strings.Split(env.KafkaBroker, ","), env.KafkaGroup, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			err := client.Consume(ctx, strings.Split(env.KafkaTopic, ","), &consumer)
			if err != nil {
				fmt.Println("DEU RUIM, ERRO: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")

	case <-sigterm:
		log.Println("terminating: via signal")
	}

	cancel()
	wg.Wait()

	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}

	return true
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

		session.MarkMessage(message, "")
	}

	return nil
}

//======================================================================================================================

func saveMessage(message string) {
	db := getConnectionDB()
	statement, err := db.Prepare(
		"INSERT INTO messages (message) VALUES (?)",
	)

	if err != nil {
		log.Fatal("Erro: ", err)
	}
	defer statement.Close()

	res, err := statement.Exec(message)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	fmt.Println(lastID)
}

func getConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", env.ConnectionDB)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal("Erro: ", err)
	}

	return db
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql" // Driver
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", list)
	e.GET("/receive", receive)
	e.POST("/send", send)

	e.Logger.Fatal(e.Start(":7000"))
}

func list(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func receive(c echo.Context) error {
	// message := "sdjasida"
	getMessages()

	return c.String(http.StatusOK, "Hello, World!")
}

func send(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!2")
}

//======================================================================================================================

type MessageKafka struct {
	Message string `json:"message"`
}

// func getMessages() {
// 	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
// 	if err != nil {
// 		log.Fatal("Erro2: ", err)
// 	}

// 	pc, err := consumer.ConsumePartition("test", int32(0), sarama.OffsetNewest)
// 	if err != nil {
// 		fmt.Printf("failed to start consumer for partition %d,err:%v\n", 0, err)
// 		return
// 	}
// 	defer pc.AsyncClose()

// 	go func(sarama.PartitionConsumer) {
// 		for msg := range pc.Messages() {
// 			var message MessageKafka
// 			json.Unmarshal(msg.Value, &message)

// 			fmt.Println(message)

// 			// fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
// 		}
// 	}(pc)

// 	select {}
// }

// func getMessages() {
// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true
// 	brokers := []string{"localhost:9092"}

// 	cluster, err := sarama.NewConsumer(brokers, config)
// 	if err != nil {
// 		log.Fatal("Erro: ", err)
// 	}

// 	defer func() {
// 		if err := cluster.Close(); err != nil {
// 			log.Fatal("Erro: ", err)
// 		}
// 	}()

// 	topic, _ := cluster.Topics()
// 	for index := range topic {
// 		fmt.Println(topic[index])
// 	}
// }

func getMessages() {
	// config := sarama.NewConfig()
	// config.Consumer.Return.Errors = true

	// brokers := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions("test")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}

		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)

		select {}
	}
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
	user := "user"
	password := "secret"
	host := "go_kakfa-db"
	port := "3306"
	database := "database"

	stringConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal("Erro: ", err)
	}

	return db
}

package banco

import (
	"api/src/config"
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql" // Driver
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Banco interface {
	Conectar() (sql.Conn, error)
}

type MySQL struct {
	Banco
}

type Mongo struct {
	Banco
}

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}

func (mysql MySQL) Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}

func (m Mongo) Conectar() (*mongo.Client, error) {
	conf := options.Client().ApplyURI(config.MongoHost)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, erro := mongo.Connect(ctx, conf)
	if erro != nil {
		return nil, erro
	}

	if erro = client.Ping(ctx, readpref.Primary()); erro != nil {
		return nil, errors.New("Falha ao conectar ao mongo")
	}

	return client, nil
}

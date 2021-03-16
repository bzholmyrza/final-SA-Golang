package main

import (
	"context"
	"final-SA-Golang/user_microservice/pkg/models/postgresql"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	dsn := flag.String("dsn", "postgres://postgres:123@localhost:5432/musicApp", "PostgreSql data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := &postgresql.UserModel{DB: db}

	id, err := models.UpdateUser("beibarys", "beibarys", "beibarys", 2, 1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

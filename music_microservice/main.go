package main

import (
	"context"
	"flag"
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
	/*
		models := &postgresql.SongModel{DB: db}
		id, err := models.GetSongByID(2)

		fmt.Println(id)
		songs, err := models.GetAllSongs()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(songs)*/
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

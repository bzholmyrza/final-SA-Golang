package main

import (
	"context"
	"final-SA-Golang/music_microservice/pkg/models/postgresql"
	protopb "final-SA-Golang/songpb"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	dsn := flag.String("dsn", "postgres://postgres:123@localhost:5432/musicApp", "PostgreSql data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s := grpc.NewServer()
	protopb.RegisterSongServiceServer(s, &Server{
		ps: &postgresql.SongModel{DB: db},
	})
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

type Server struct {
	protopb.UnimplementedSongServiceServer
	ps *postgresql.SongModel
}

func (s *Server) CreateSong(ctx context.Context, req *protopb.CreateSongRequest) (*protopb.CreateSongResponse, error) {
	id, err := s.ps.CreateSong(req.Song.Title, req.Song.Author, req.Song.ReleaseDate)
	if err != nil {
		return nil, err
	}
	res := &protopb.CreateSongResponse{Id: int32(id)}
	return res, nil
}

func (s *Server) DeleteSong(ctx context.Context, req *protopb.DeleteSongRequest) {
	s.ps.DeleteSong(int(req.Id))
}

func (s *Server) GetSong(req *protopb.GetSongRequest) (*protopb.GetSongResponse, error) {
	song, err := s.ps.GetSongByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	res := &protopb.GetSongResponse{Song: &protopb.Song{
		Id:          int64(song.ID),
		Title:       song.Title,
		Author:      song.Author,
		ReleaseDate: song.ReleaseDate,
	}}
	return res, nil
}

func (s *Server) GetAllSongs(req *protopb.GetAllSongsRequest, stream protopb.SongService_GetAllSongsServer) error {
	songs, err := s.ps.GetAllSongs()
	if err != nil {
		return err
	}
	for i := 0; i < len(songs); i++ {
		res := &protopb.GetAllSongsResponse{Song: &protopb.Song{
			Id:          int64(songs[i].ID),
			Title:       songs[i].Title,
			Author:      songs[i].Author,
			ReleaseDate: songs[i].ReleaseDate,
		}}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending greet many times responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

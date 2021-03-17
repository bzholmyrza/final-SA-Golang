package main

import (
	"context"
	"final-SA-Golang/favoritepb"
	"final-SA-Golang/favorites_microservice/pkg/models/postgresql"
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
	favoritepb.RegisterFavoriteServiceServer(s, &Server{
		ps: &postgresql.FavoritesModel{DB: db},
	})
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Server struct {
	favoritepb.UnimplementedFavoriteServiceServer
	ps *postgresql.FavoritesModel
}

func (s *Server) CreateFavorite(ctx context.Context, req *favoritepb.CreateFavoriteRequest) (*favoritepb.CreateFavoriteResponse, error) {
	id, err := s.ps.CreateFavorites(int(req.Favorite.UserId), int(req.Favorite.SongId))
	if err != nil {
		return nil, err
	}
	res := &favoritepb.CreateFavoriteResponse{Id: int32(id)}
	return res, nil
}

func (s *Server) DeleteFavorite(ctx context.Context, req *favoritepb.DeleteFavoriteRequest) {
	s.ps.DeleteFavorites(int(req.UId), int(req.SId))
}

func (s *Server) GetFavorite(req *favoritepb.GetFavoriteRequest, stream favoritepb.FavoriteService_GetFavoriteServer) error {
	songs, err := s.ps.GetFavoritesByUserID(int(req.UserId))
	if err != nil {
		return err
	}
	for i := 0; i < len(songs); i++ {
		res := &favoritepb.GetFavoriteResponse{Favorite: &favoritepb.Favorite{
			Id:     int64(songs[i].ID),
			UserId: int64(songs[i].UserID),
			SongId: int64(songs[i].SongID),
		}}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending greet many times responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

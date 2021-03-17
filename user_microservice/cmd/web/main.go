package main

import (
	"context"
	"final-SA-Golang/user_microservice/pkg/models/postgresql"
	"final-SA-Golang/userpb"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
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
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &Server{
		ps: &postgresql.UserModel{DB: db},
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
	userpb.UnimplementedUserServiceServer
	ps *postgresql.UserModel
}

func (s *Server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	id, err := s.ps.CreateUser(req.User.Username, req.User.Email, req.User.Password, int(req.User.Role))
	if err != nil {
		return nil, err
	}
	res := &userpb.CreateUserResponse{Id: int32(id)}
	return res, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id, err := s.ps.UpdateUser(req.User.Username, req.User.Email, req.User.Password, int(req.User.Id), int(req.User.Role))
	if err != nil {
		return nil, err
	}
	res := &userpb.UpdateUserResponse{Id: int32(id)}
	return res, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) {
	s.ps.DeleteUser(int(req.Id))
}

func (s *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.ps.GetUser(int(req.Id))
	if err != nil {
		return nil, err
	}
	res := &userpb.GetUserResponse{User: &userpb.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     int64(user.Role),
	}}
	return res, nil
}

func (s *Server) LoginUser(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	user, err := s.ps.GetUserByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	loginRsponce := &userpb.LoginResponse{User: &userpb.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     int64(user.Role),
	}}
	return loginRsponce, nil
}

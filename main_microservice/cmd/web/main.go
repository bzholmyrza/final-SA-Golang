package main

import (
	"final-SA-Golang/favoritepb"
	"final-SA-Golang/songpb"
	"final-SA-Golang/userpb"
	"flag"
	"github.com/golangcollege/sessions"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	userClient     userpb.UserServiceClient
	songClient     protopb.SongServiceClient
	favoriteClient favoritepb.FavoriteServiceClient
	errorLog       *log.Logger
	infoLog        *log.Logger
	session        *sessions.Session
	templateCache  map[string]*template.Template
}

func main() {
	userConn, userErr := grpc.Dial("localhost:50051", grpc.WithInsecure())
	songConn, songErr := grpc.Dial("localhost:50052", grpc.WithInsecure())
	favoriteConn, favoriteErr := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if userErr != nil {
		log.Fatalf("could not connect: %v", userErr)
	}
	defer userConn.Close()
	if songErr != nil {
		log.Fatalf("could not connect: %v", songErr)
	}
	defer songConn.Close()
	if favoriteErr != nil {
		log.Fatalf("could not connect: %v", favoriteErr)
	}
	defer favoriteConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)
	songClient := protopb.NewSongServiceClient(songConn)
	favoriteClient := favoritepb.NewFavoriteServiceClient(favoriteConn)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		userClient:     userClient,
		songClient:     songClient,
		favoriteClient: favoriteClient,
		errorLog:       errorLog,
		infoLog:        infoLog,
		session:        session,
		templateCache:  templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server running on port %v", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

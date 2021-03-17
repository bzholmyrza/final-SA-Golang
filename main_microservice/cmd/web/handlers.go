package main

import (
	"context"
	"final-SA-Golang/favoritepb"
	_ "final-SA-Golang/favorites_microservice/pkg/models"
	"final-SA-Golang/music_microservice/pkg/models"
	protopb "final-SA-Golang/songpb"
	"final-SA-Golang/userpb"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	getAllSongs(app.songClient)
}

func (app *application) createUserWrapper(w http.ResponseWriter, r *http.Request) {
	createUser(app.userClient)
}

func createUser(c userpb.UserServiceClient) {
	ctx := context.Background()

	request := &userpb.CreateUserRequest{User: &userpb.User{
		Id:       -1,
		Username: "yolo62442",
		Password: "25335",
		Email:    "yolo62442@gmail.com",
		Role:     0,
	}}

	response, err := c.CreateUser(ctx, request)
	if err != nil {
		log.Fatalf("error while calling User server RPC %v", err)
	}
	log.Printf("response from User server:%v", response.Id)
}

func (app *application) getUserWrapper(w http.ResponseWriter, r *http.Request) {
	getUser(app.userClient)
}

func getUser(c userpb.UserServiceClient) {
	ctx := context.Background()

	request := &userpb.GetUserRequest{Id: 3}

	response, err := c.GetUser(ctx, request)
	if err != nil {
		log.Fatalf("error while calling User server RPC %v", err)
	}
	log.Printf("response from User server:%v", response.User)
}

func (app *application) deleteUserWrapper(w http.ResponseWriter, r *http.Request) {
	deleteUser(app.userClient)
}

func deleteUser(c userpb.UserServiceClient) {
	ctx := context.Background()

	request := &userpb.DeleteUserRequest{Id: 3}

	response, err := c.DeleteUser(ctx, request)
	if err != nil {
		log.Fatalf("error while calling User server RPC %v", err)
	}
	log.Printf("response from User server:%v", response)
}

func (app *application) updateUserWrapper(w http.ResponseWriter, r *http.Request) {
	updateUser(app.userClient)
}

func updateUser(c userpb.UserServiceClient) {
	ctx := context.Background()

	request := &userpb.UpdateUserRequest{User: &userpb.User{
		Id:       2,
		Username: "yolo62442",
		Password: "123123",
		Email:    "yolo62442@gmail.com",
		Role:     0}}

	response, err := c.UpdateUser(ctx, request)
	if err != nil {
		log.Fatalf("error while calling User server RPC %v", err)
	}
	log.Printf("response from User server:%v", response.Id)
}

func (app *application) loginUserWrapper(w http.ResponseWriter, r *http.Request) {
	loginUser(app.userClient)
}

func loginUser(c userpb.UserServiceClient) {
	ctx := context.Background()

	request := &userpb.LoginRequest{Email: "yolo62442@gmail.com", Password: "123123"}

	response, err := c.LoginUser(ctx, request)
	if err != nil && err != pgx.ErrNoRows {
		log.Fatalf("error while calling User LOGIN RPC %v", err)
	}
	log.Printf("response from User server:%v", response.User)
}

func (app *application) createSongWrapper(w http.ResponseWriter, r *http.Request) {
	createSong(app.songClient)
}

func createSong(c protopb.SongServiceClient) {
	ctx := context.Background()

	request := &protopb.CreateSongRequest{Song: &protopb.Song{
		Id:          -1,
		Title:       "Hello",
		Author:      "Adel",
		ReleaseDate: "12.10.2013",
	}}

	response, err := c.CreateSong(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Song server RPC %v", err)
	}
	log.Printf("response from Song server:%v", response.Id)
}

func (app *application) getSongWrapper(w http.ResponseWriter, r *http.Request) {
	getSong(app.songClient, 1)
}

func getSong(c protopb.SongServiceClient, id int32) {
	ctx := context.Background()

	request := &protopb.GetSongRequest{Id: id}

	res, err := c.GetSong(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Song server RPC %v", err)
	}
	s := &models.Song{
		ID:          -1,
		Title:       "Hello",
		Author:      "Adel",
		ReleaseDate: "12.10.2013",
	}
	log.Printf("response from Song server:%v", res.Song.Title)
	fmt.Println(s)
	/*app.render(w, r, "show.page.tmpl", &templateData{
		Song: s,
	})*/
}

func (app *application) deleteSongWrapper(w http.ResponseWriter, r *http.Request) {
	deleteSong(app.songClient)
}

func deleteSong(c protopb.SongServiceClient) {
	ctx := context.Background()

	request := &protopb.DeleteSongRequest{Id: 1}

	response, err := c.DeleteSong(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Song server RPC %v", err)
	}
	log.Printf("response from Song server:%v", response)
}

func (app *application) updateSongWrapper(w http.ResponseWriter, r *http.Request) {
	updateSong(app.songClient)
}

func updateSong(c protopb.SongServiceClient) {
	ctx := context.Background()

	request := &protopb.UpdateSongRequest{Song: &protopb.Song{
		Id:          2,
		Title:       "Watermelon Sugar High",
		Author:      "Harry Stiles",
		ReleaseDate: "09.07.2019",
	}}

	response, err := c.UpdateSong(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Song server RPC %v", err)
	}
	log.Printf("response from Song server:%v", response.Id)
}

func (app *application) getAllSongsWrapper(w http.ResponseWriter, r *http.Request) {
	getAllSongs(app.songClient)
}

func getAllSongs(c protopb.SongServiceClient) {
	ctx := context.Background()
	req := &protopb.GetAllSongsRequest{}

	stream, err := c.GetAllSongs(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GET ALL POSTS RPC %v", err)
	}
	defer stream.CloseSend()

	var songs []*models.Song

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error while reciving from get all posts RPC %v", err)
		}
		s := res.Song
		songs = append(songs, &models.Song{
			ID:          int(s.Id),
			Title:       s.Title,
			Author:      s.Author,
			ReleaseDate: s.ReleaseDate,
		})
		/*app.render(w, r, "home.page.tmpl", &templateData{
			Songs: songs,
		})*/
		log.Printf("response from get all posts:%v \n", res.GetSong().Title)
	}
}

func (app *application) createFavoritesWrapper(w http.ResponseWriter, r *http.Request) {
	createFavorites(app.favoriteClient)
}

func createFavorites(c favoritepb.FavoriteServiceClient) {
	ctx := context.Background()

	request := &favoritepb.CreateFavoriteRequest{Favorite: &favoritepb.Favorite{
		Id:     -1,
		UserId: 1,
		SongId: 2,
	}}

	response, err := c.CreateFavorite(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Song server RPC %v", err)
	}
	log.Printf("response from Song server:%v", response.Id)
}

func (app *application) showFavoritesWrapper(w http.ResponseWriter, r *http.Request) {
	showFavorites(app.favoriteClient)
}

func showFavorites(c favoritepb.FavoriteServiceClient) {
	ctx := context.Background()
	req := &favoritepb.GetFavoriteRequest{UserId: 2}

	stream, err := c.GetFavorite(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from GreetManyTimes RPC %v", err)
		}
		log.Printf("response from GreetManyTimes:%v \n", res.GetFavorite())
	}

}

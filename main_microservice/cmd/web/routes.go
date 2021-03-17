package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.loginUserWrapper))
	mux.Post("/users/registration", dynamicMiddleware.ThenFunc(app.createUserWrapper))
	mux.Get("/users/:id", dynamicMiddleware.ThenFunc(app.getUserWrapper))
	mux.Get("/users/delete/:id", dynamicMiddleware.ThenFunc(app.deleteUserWrapper))
	mux.Get("/users/update/:id", dynamicMiddleware.ThenFunc(app.updateUserWrapper))

	mux.Post("/songs", dynamicMiddleware.ThenFunc(app.getAllSongsWrapper))
	mux.Post("/songs/add", dynamicMiddleware.ThenFunc(app.createSongWrapper))
	mux.Get("/songs/:id", dynamicMiddleware.ThenFunc(app.getSongWrapper))
	mux.Get("/songs/delete/:id", dynamicMiddleware.ThenFunc(app.deleteSongWrapper))
	mux.Get("/songs/update/:id", dynamicMiddleware.ThenFunc(app.updateSongWrapper))

	mux.Post("/favorite/add", dynamicMiddleware.ThenFunc(app.createFavoritesWrapper))
	mux.Get("/favorite/:id", dynamicMiddleware.ThenFunc(app.showFavoritesWrapper))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}

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
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.login))
	mux.Get("/users/registration", dynamicMiddleware.ThenFunc(app.registrationForm))
	mux.Post("/users/registration", dynamicMiddleware.ThenFunc(app.registration))
	mux.Get("/users/:id", dynamicMiddleware.ThenFunc(app.showUser))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}

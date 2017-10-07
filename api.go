package main

import "net/http"
import "github.com/julienschmidt/httprouter"

// Bind adds all routes to the application route
func (app *App) Bind() {
	app.Router.GET("/api/", app.home)
}

func (app *App) home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("API home"))
}

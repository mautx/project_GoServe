package main

import "net/http"

// Función que conecta los handlers con el regex
func (app *application) routes() *http.ServeMux {

	mix := http.NewServeMux()
	mix.HandleFunc("/", app.home)
	mix.HandleFunc("/snippet", app.show)
	mix.HandleFunc("/snippet/create", app.create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mix.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mix
}

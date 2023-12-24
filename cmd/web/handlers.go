package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Esta Función se encarga de decir qué hacer
// w es el encargado de responder
// r es el posible request que se pide al servidor
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//Por default, root es tratado como un comodín o /*, esta sección elimina eso
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	w.Header().Add("X-Custom-Header", "Cachete")
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// We then use the Execute() method on the template set to write the templa
	// content as the response body. The last parameter to Execute() represents
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) show(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Mostrando snippet %d", id)

}

func (app *application) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not alowe", 405)
		return
	}
}

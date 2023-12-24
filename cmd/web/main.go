package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	//En go es común realizar "banderas" para mejorar las variables que vamos a ocupar
	// se puede usar -help en cmd
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Aqui instanciamos el ServerMux; su función es mapear el patrón
	// URL con la función.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", show)
	mux.HandleFunc("/snippet/create", create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Activamos el server en el puerto 4001
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

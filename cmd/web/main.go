package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	//En go es común realizar "banderas" para mejorar las variables que vamos a ocupar
	// se puede usar -help en cmd
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//Infolog es un mensaje que aparece en terminal con informacion relevante
	infoLog := log.New(os.Stdout, "INFO---\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR---\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Aqui instanciamos el ServerMux; su función es mapear el patrón
	// URL con la función.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", show)
	mux.HandleFunc("/snippet/create", create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Activamos el server en el puerto 4001
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}

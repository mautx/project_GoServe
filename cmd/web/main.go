package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//En go es común realizar "banderas" para mejorar las variables que vamos a ocupar
	// se puede usar -help en cmd
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//Infolog es un mensaje que aparece en terminal con informacion relevante
	infoLog := log.New(os.Stdout, "INFO---\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR---\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//Aqui instanciamos el ServerMux; su función es mapear el patrón
	// URL con la función.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.show)
	mux.HandleFunc("/snippet/create", app.create)

	//Manejo y response de archivos estáticos creando un directorio
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Creamos el objeto servidor con las direcciones, mensaje de error y handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//Activamos el server en el puerto 4001
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

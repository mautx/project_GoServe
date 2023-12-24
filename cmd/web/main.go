package main

import (
	"Snipperclips/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	//En go es común realizar "banderas" para mejorar las variables que vamos a ocupar
	// se puede usar -help en cmd
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=True", "MariaDB database User")

	flag.Parse()

	//Info log es un mensaje que aparece en terminal con información relevante
	infoLog := log.New(os.Stdout, "INFO---\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR---\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	//Aqua instanciamos el ServerMux; su función es mapear el patrón
	// URL con la función.
	mux := app.routes()

	//Creamos el objeto servidor con las direcciones, mensaje de error y handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//Activamos el server en el puerto 4001
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

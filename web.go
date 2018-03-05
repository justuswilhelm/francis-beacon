package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
)

const (
	DefaultPort = "8080"
)

var (
	db      *sql.DB
	connStr string
)

func index(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC3339)

	_, err := db.Exec(
		"INSERT INTO beacon_hit (date, referer, path, host, query) VALUES ($1, $2, $3, $4, $5)",
		now,
		r.Referer(),
		r.URL.Path,
		r.Host,
		r.URL.RawQuery,
	)
	if err != nil {
		log.Printf("Error when inserting beacon hit: %+v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}

func main() {
	var err error
	connStr = os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	n.Run(":" + port)
}

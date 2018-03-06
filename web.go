package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
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

	if r.Referer() == "" {
		http.Error(w, "Empty Referer", http.StatusBadRequest)
		return
	}

	url, err := url.Parse(r.Referer())

	if err != nil {
		log.Printf("Error when inserting beacon hit: %+v", err)
		http.Error(w, "Invalid Referer", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO beacon_hit (date, scheme, host, path, query, fragment) VALUES ($1, $2, $3, $4, $5, $6)",
		now,
		url.Scheme,
		url.Host,
		url.Path,
		url.RawQuery,
		url.Fragment,
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

package main

import (
	"encoding/json"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

const (
	DefaultPort = "8080"
)

func index(w http.ResponseWriter, r *http.Request) {
	response := "Hello World"
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("Error when encoding response: %+v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func main() {
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

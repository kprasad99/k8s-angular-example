package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("/static/")))
	http.Handle("/app2/", http.StripPrefix("/app2/", http.FileServer(http.Dir("/static/"))))
	// http.HandleFunc("/app1/", serveFiles)

	port := os.Getenv("WEB_PORT")
	host := os.Getenv("WEB_HOST")

	address := ""

	if port == "" {
		port = "8080"
	}

	if host == "" {
		address = ":" + port
	} else {
		address = host + ":" + port
	}

	log.Printf("Web server start at %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/index.html")
}

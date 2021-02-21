package router

import (
	"log"
	"net/http"
	"os"
)

// RootRouter is router that binds passed handler functions to the root path
func RootRouter(handler http.HandlerFunc) {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

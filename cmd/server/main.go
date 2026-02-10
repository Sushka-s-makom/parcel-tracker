package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Sushka-s-makom/parcel-tracker/internal/httpapi"
)

func main() {
	h := httpapi.Handlers{}
	router := httpapi.NewRouter(h)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}

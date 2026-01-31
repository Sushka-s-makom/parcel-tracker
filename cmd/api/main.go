package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Track struct {
	ID           string `json:"id"`
	TrackNumber  string `json:"track_number"`
	Carrier      string `json:"carrier"`
	Status       string `json:"status"`
	LastUpdateAt string `json:"last_update_at"`
	CreatedAt    string `json:"created_at"`
}

type CreateTrackRequest struct {
	TrackNumber string `json:"track_number"`
	Carrier     string `json:"carrier"`
}

var (
	mu     sync.Mutex
	tracks = make(map[string]Track)
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		_, _ = fmt.Fprintln(w, "ok")
	})

	mux.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte(`{"error":"not implemented yet"}`))
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("listening on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())

}

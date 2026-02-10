package httpapi

import "net/http"

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/db-health", h.DBHealth)
	mux.HandleFunc("/tracks", h.Tracks)
	return mux
}

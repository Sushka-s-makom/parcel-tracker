package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Sushka-s-makom/parcel-tracker/internal/config"
	"github.com/Sushka-s-makom/parcel-tracker/internal/httpapi"
	"github.com/Sushka-s-makom/parcel-tracker/internal/repository/pg"
)

func main() {
	cfg := config.Load()
	if cfg.Database == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx := context.Background()

	db, err := pg.NewPool(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

package controller

import (
	"cheki-back/config"
	"cheki-back/db"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	_, cleanup, err := db.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	return mux, cleanup, nil
}

package main

import (
	"context"
	"net/http"

	"cheki-back/config"
	"cheki-back/controller"
	"cheki-back/db"
	"cheki-back/middleware"
	"github.com/go-chi/chi/v5"
	mid "github.com/go-chi/chi/v5/middleware"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.Use(mid.Logger)
	mux.Get("/health", middleware.AddHeader(controller.GetHealth))

	xdb, cleanup, err := db.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	at := &controller.SqlxUser{DB: xdb}
	mux.Get("/users", at.GetUser)

	return mux, cleanup, nil
}

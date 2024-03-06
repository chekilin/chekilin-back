package middleware

import (
	"log"
	"net/http"
)

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Add header")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
		log.Println("Completed in Adding header")
	})
}

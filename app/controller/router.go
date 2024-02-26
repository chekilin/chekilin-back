package controller

import (
	"net/http"
)

var mux http.ServeMux

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Test)
}

func NewMux() *http.ServeMux {
	return &mux
}

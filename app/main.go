package main

import (
	"net/http"

	"cheki-back/controller"
)

func main() {
	mux := controller.NewMux()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

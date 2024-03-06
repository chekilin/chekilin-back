package controller

import (
	"net/http"
)

func GetHealth(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}

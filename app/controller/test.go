package controller

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, _ *http.Request) {
	hello := "Hello World"
	result, err := json.Marshal(hello)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(result)
}

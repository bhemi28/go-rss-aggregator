package main

import (
	"net/http"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

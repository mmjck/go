package main

import (
	"net/http"
)

func handlerReadliness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

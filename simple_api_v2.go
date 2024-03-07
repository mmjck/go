package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content=Type", "application/json")

	json.NewEncoder(w).Encode([]User{{
		Id: 1, Name: "Joao",
	}})
}

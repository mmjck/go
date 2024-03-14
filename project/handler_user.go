package main

import (
	"encoding/json"
	"example/hello/project/internal/database"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: ", err))
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	log.Println(user)

	if err != nil {
		log.Println(err)
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: ", err))
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *ApiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/xiayulin123/GO_BackendWebServer/interal/auth"
	"github.com/xiayulin123/GO_BackendWebServer/interal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create User: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKEY(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GETAPIKEY(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't find user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

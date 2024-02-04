package main

import (
	"fmt"
	"net/http"

	"github.com/xiayulin123/GO_BackendWebServer/interal/auth"
	"github.com/xiayulin123/GO_BackendWebServer/interal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)

	}
}

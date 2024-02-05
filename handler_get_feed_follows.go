package main

import (
	"fmt"
	"net/http"

	"github.com/xiayulin123/GO_BackendWebServer/interal/database"
)

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_follows, err := apiCfg.DB.GetFeedFollowes(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feed_follows))
}

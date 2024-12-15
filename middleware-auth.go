package main

import (
	"fmt"
	"net/http"

	"github.com/rashmod/rssagg/internal/auth"
	"github.com/rashmod/rssagg/internal/database"
)

type authhandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

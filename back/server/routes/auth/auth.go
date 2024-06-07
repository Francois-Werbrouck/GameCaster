package auth

import (
	"GameCaster/main/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRouteWithAuth(router *mux.Router, path string) *mux.Router {

	authenticatedRoute := router.PathPrefix(path).Subrouter()
	addRouteTest(authenticatedRoute)

	return authenticatedRoute
}

func addRouteTest(router *mux.Router) {

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		user, ok := r.Context().Value("authToken").(db.User)
		if !ok {
			res := map[string]string{
				"status": "failed",
				"error":  "not session token found",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			res := map[string]string{
				"status":  "succeded",
				"message": fmt.Sprint(user.Email),
			}

			json.NewEncoder(w).Encode(res)
		}
	}).Methods("GET")
}

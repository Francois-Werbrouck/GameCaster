package main

import (
	"GameCaster/main/server"
	"GameCaster/main/server/routes/auth"
	"GameCaster/main/sqlobjects"
	"GameCaster/main/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	port, succes := os.LookupEnv("PORT")
	if !succes {
		port = "8000"
	}

	database := OpenDB("db/dev.db")

	Serve(mux.NewRouter(), database, port)

}

func Serve(route *mux.Router, database *gorm.DB, port string) {

	// setup les middlewares les plus frequents
	route.Use(mux.CORSMethodMiddleware(route))
	route.Use(server.Cors)
	route.Use(server.Logger)

	route.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
	}).Methods("GET")

	route.HandleFunc("/api/status", server.StatusHandler).Methods("GET")

	route.Handle("/", http.FileServer(http.Dir("../front/dist/"))).Methods("GET")

	authRoute := route.PathPrefix("/api/auth").Subrouter()

	authenticatedRoute := auth.HandleRouteWithAuth(route, "/api/authenticated")
	authenticatedRoute.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("authToken")

			var ctx context.Context
			if err != nil {

				next.ServeHTTP(w, r)
				return
			}
			user, err := sqlobjects.GetUserByToken(cookie.Value, database)

			if err != nil {
				fmt.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			ctx = context.WithValue(r.Context(), "authToken", user)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		})
	})

	authRoute.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This will soon be a nice page")
	})

	authRoute.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var userJson UserJson

		err = json.Unmarshal(body, &userJson)

		_, err = sqlobjects.GetUserByEmail(userJson.Email, database)

		if err != nil {
			if err == gorm.ErrRecordNotFound {

				user, _ := sqlobjects.CreateUser(userJson.Email, userJson.Password, database)
				fmt.Println("after create user and before CreateSessionToken")

				token, _ := sqlobjects.CreateSessionToken(user, database)
				coockie := http.Cookie{
					Name:     "authToken",
					Value:    token,
					Path:     "/",
					MaxAge:   3600 * 24 * 5, // 3600 (seconds) 24 (hours) 5 (days)
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				}
				http.SetCookie(w, &coockie)
				res := map[string]string{
					"status": "succeded",
				}
				json.NewEncoder(w).Encode(res)

			} else {

				http.Error(w, "Error interfacing with database", http.StatusBadRequest)
			}

		} else {
			res := map[string]string{
				"status": "failed",
				"error":  "The user already exist. Please login or click on `Forgot my password`",
			}
			fmt.Println(res)
			json.NewEncoder(w).Encode(res)
		}

	}).Methods("POST")

	authRoute.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		var userRequest UserJson
		err := utils.ReadToJson(&userRequest, r)
		if err != nil {
			res := map[string]string{
				"status": "failed",
				"error":  "Unable to parse json",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		user, err := sqlobjects.GetUserByEmail(userRequest.Email, database)

		if err != nil {

			res := map[string]string{
				"status": "failed",
				"error":  "Unable to parse find user",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		if sqlobjects.IsPasswordValid(userRequest.Password, user) {

			token, _ := sqlobjects.CreateSessionToken(user, database)

			coockie := http.Cookie{
				Name:     "authToken",
				Value:    token,
				Path:     "/",
				MaxAge:   3600 * 24 * 5, // 3600 (seconds) 24 (hours) 5 (days)
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, &coockie)
			res := map[string]string{
				"status": "succeded",
				"error":  "Authenticated succesfully",
			}
			json.NewEncoder(w).Encode(res)

		} else {

			res := map[string]string{
				"status": "failed",
				"error":  "Credentials did not match",
			}
			json.NewEncoder(w).Encode(res)

		}

	}).Methods("POST")

	fmt.Printf("Listening at http://localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), route)
}

func OpenDB(path string) *gorm.DB {

	database, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("could not open databse")
	}

	database.AutoMigrate(&sqlobjects.UserToken{})
	database.AutoMigrate(&sqlobjects.User{})

	return database

}

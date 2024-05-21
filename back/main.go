package main

import (
	"GameCaster/main/server"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func main() {
	port, succes := os.LookupEnv("PORT")
	if !succes {
		port = "8000"
	}

	openDB("db/dev.db")

	http.Handle("GET /", http.FileServer(http.Dir("../front/dist/")))

	http.HandleFunc("GET /status", server.Middling(server.StatusHandler, server.Cors, server.Logger))
	http.HandleFunc("GET /message", server.Middling(server.MessageHandler, server.Cors, server.Logger))

	fmt.Printf("Listening at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func openDB(path string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("could not open databse")
	}

	db.AutoMigrate(&User{})

	return db

}

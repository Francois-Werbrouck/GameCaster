package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(res)

}

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

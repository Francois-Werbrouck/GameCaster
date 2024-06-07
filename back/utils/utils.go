package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadToJson(variable any, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &variable)

	if err != nil {
		return err
	}
	return nil

}

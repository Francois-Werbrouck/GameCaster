package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

const SERVER_PORT = "12244"

type ServerSuite struct {
	suite.Suite
}

func (server *ServerSuite) SetupSuite() {

	go Serve(mux.NewRouter(), OpenDB("db/dev.db"), SERVER_PORT)
	time.Sleep(15 * time.Millisecond)
}

func TestMain(t *testing.T) {
	suite.Run(t, new(ServerSuite))

}

func (server *ServerSuite) TestStatus() {

	res, err := server.Get("/api/status")

	server.Assertions.Nil(err)
	defer res.Body.Close()
	server.Assertions.Equal(res.StatusCode, 200)

	expected := map[string]string{
		"status": "ok",
	}

	var got map[string]string
	json.NewDecoder(res.Body).Decode(&got)

	server.Assertions.Equal(got, expected)

}

func (server *ServerSuite) Get(path string) (resp *http.Response, err error) {
	res, err := http.Get("http://localhost:" + SERVER_PORT + path)
	return res, err
}

func (server *ServerSuite) Post(path string, contentType string, body io.Reader) (resp *http.Response, err error) {
	res, err := http.Post("http://localhost:"+SERVER_PORT+path, contentType, body)
	return res, err
}

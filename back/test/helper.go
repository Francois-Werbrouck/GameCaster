package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const DEFAULT_TESTING_PORT = "13455"

func SetupServer() {

	httptest.NewServer()
}

func Get(url string) (resp *http.Response, err error) {
	return http.Get("http://localhost:" + DEFAULT_TESTING_PORT + url)
}

func AssertEqual(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func ExpectError(t *testing.T, err error, expectedErrMsg string) {
	if err == nil {
		t.Errorf("expected error but got nil")
	} else if err.Error() != expectedErrMsg {
		t.Errorf("expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
	}
}

func LogInfo(t *testing.T, msg string, args ...interface{}) {
	t.Logf("[INFO] "+msg, args...)
}

func LogError(t *testing.T, msg string, args ...interface{}) {
	t.Logf("[ERROR] "+msg, args...)
}

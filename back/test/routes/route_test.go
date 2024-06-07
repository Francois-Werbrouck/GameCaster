package routes

import (
	`GameCaster/main/test`
	"testing"
)

func status_route_test(t *testing.T) {

	test.SetupServer()
	req, err := test.Get("/status")
	test.AssertNoError(t, err)
	test.AssertEqual(t, req.Body, "nice")
}

package integration_setup

import (
	"net/http/httptest"

	"github.com/gabriel-jm/clone-tabnews/src/server"
)

func RunTestServer() *httptest.Server {
	return httptest.NewServer(server.SetupServer())
}

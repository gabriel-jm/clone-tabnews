package server

import (
	"net/http"

	"github.com/gabriel-jm/clone-tabnews/src/status"
)

func SetupServer() *http.ServeMux {
	mainMux := http.NewServeMux()

	statusMux := http.NewServeMux()
	statusMux.HandleFunc("/status", status.GetStatus)

	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", statusMux))

	return mainMux
}

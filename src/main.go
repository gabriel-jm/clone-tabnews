package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabriel-jm/clone-tabnews/src/status"
)

func main() {
	mainMux := http.NewServeMux()

	statusMux := http.NewServeMux()
	statusMux.HandleFunc("/status", status.GetStatus)

	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", statusMux))

	fmt.Println("Server Running")
	var serverErr = http.ListenAndServe(":8000", mainMux)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

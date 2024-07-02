package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabriel-jm/clone-tabnews/src/status"
)

func main() {
	http.HandleFunc("/api/v1/status", status.GetStatus)

	fmt.Println("Server Running")
	var serverErr = http.ListenAndServe(":8000", nil)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabriel-jm/clone-tabnews/src/db"
	"github.com/gabriel-jm/clone-tabnews/src/status"
	"github.com/joho/godotenv"
)

const query string = `
	CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text
	);
`

func main() {
	loadEnvErr := godotenv.Load(".env")

	if loadEnvErr != nil {
		log.Fatalln("Error loading env file:", loadEnvErr)
	}

	db.Init()
	defer db.Conn.Close()

	db.Conn.MustExec(query)

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

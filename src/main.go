package main

import (
	"log"
	"net/http"

	"github.com/gabriel-jm/clone-tabnews/src/db"
	"github.com/gabriel-jm/clone-tabnews/src/server"
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

	serverMux := server.SetupServer()
	println("Server Running")
	var serverErr = http.ListenAndServe(":8000", serverMux)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

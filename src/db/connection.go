package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func Init() {
	connectionString := os.Getenv("DB_URL")

	conn, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	err = conn.Ping()

	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	Conn = conn

	println("DB Connection successful")
}

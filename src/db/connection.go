package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func Init() {
	connectionString := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

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

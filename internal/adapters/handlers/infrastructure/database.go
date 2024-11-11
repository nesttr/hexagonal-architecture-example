package infrastructure

import (
	"database/sql"
	"log"
)

type Sql struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init(s Sql) *sql.DB {

	// will be dynamically set by the environment variables for the drivers
	connectionString := "postgres://" + s.User + ":" + s.Password + "@" + s.Host + "/" + s.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to SQL:", err)
	}
	return db
}

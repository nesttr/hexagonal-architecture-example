package http

import (
	"database/sql"
	"log"
	"net/http"
	UserRepository "odev-1/internal/adapters/repositories/postgres/user"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer db.Close()

	userRepository := UserRepository.NewUserRepository(db)

	services := NewServices(userRepository)
	router := NewRouter(services)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

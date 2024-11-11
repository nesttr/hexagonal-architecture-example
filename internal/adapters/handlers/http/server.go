package http

import (
	_ "github.com/lib/pq"
	"hexagonal-architecture-example/internal/adapters/handlers/infrastructure"
	UserRepository "hexagonal-architecture-example/internal/adapters/repositories/postgres/user"
	"log"
	"net/http"
)

type Serve struct {
	Port     string
	Postgres infrastructure.Sql
}

func ListenServe(Serve Serve) {
	db := infrastructure.Init(Serve.Postgres)
	defer db.Close()
	userRepository := UserRepository.NewUserRepository(db)

	services := NewServices(userRepository)
	router := NewRouter(services)

	log.Println("Server is running on port "+Serve.Port, "...")

	err := http.ListenAndServe(":"+Serve.Port, router)
	if err != nil {
		return
	}

}

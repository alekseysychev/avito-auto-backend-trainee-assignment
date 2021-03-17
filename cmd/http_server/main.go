package main

import (
	"log"
	"os"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/api"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/database"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/usecases"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	}
	addr := os.Getenv("HTTP_SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	linkStorage, err := database.NewPgEventStorage(dsn)
	if err != nil {
		log.Println(err)
		return
	}
	linkCases := &usecases.LinkService{
		LinkStorage: linkStorage,
	}
	server := &api.HttpServer{
		LinkUseCases: linkCases,
	}
	server.Serve(addr)
}

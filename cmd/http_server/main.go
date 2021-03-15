package main

import (
	"log"
	"os"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/database"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/http/api"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/usecases"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	addr := os.Getenv("HTTP_SERVER_ADDR")

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

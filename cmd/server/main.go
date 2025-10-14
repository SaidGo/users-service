package main

import (
	"log"
	"os"

	"github.com/SaidGo/users-service/internal/database"
	"github.com/SaidGo/users-service/internal/transport/grpc"
	"github.com/SaidGo/users-service/internal/user"
)

func main() {
	addr := os.Getenv("GRPC_ADDR")
	if addr == "" {
		addr = ":50051"
	}

	database.Init()
	repo := user.NewRepository(database.DB)
	if err := repo.Migrate(); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	svc := user.NewService(repo)

	if err := grpc.RunGRPC(svc, addr); err != nil {
		log.Fatalf("grpc serve: %v", err)
	}
}

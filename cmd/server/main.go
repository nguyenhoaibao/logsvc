package main

import (
	"log"
	"net"

	"github.com/go-pg/pg"
	"github.com/nguyenhoaibao/logsvc"
	"github.com/nguyenhoaibao/logsvc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	defer logger.Sync()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "123456",
		Database: "pg_migrations_example",
	})
	defer db.Close()

	// init repository
	repo := logsvc.NewRepository(db)
	// init service handler
	svc := logsvc.NewService(repo)
	svc = logsvc.NewLoggingService(logger)(svc)

	server := grpc.NewServer()
	pb.RegisterLogServiceServer(server, svc)

	if err := server.Serve(listener); err != nil {
		logger.Sugar().Fatalf("failed to serve: %v", err)
	}
}

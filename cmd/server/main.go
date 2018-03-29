package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"

	"github.com/go-pg/pg"
	"github.com/nguyenhoaibao/logsvc"
	"github.com/nguyenhoaibao/logsvc/config"
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

	confFile, err := ioutil.ReadFile("config/conf.yml")
	if err != nil {
		// if the config file doesn't exist, it's likely because we just ship
		// the binary alone to the server without the config file.
		// The config package handle this by also read from the enviroment variables
		// to make sure all needed variables exist.
		logger.Sugar().Warn("could not read the config file: %v", err)
	}
	conf, err := config.Load(bytes.NewReader(confFile))
	if err != nil {
		logger.Sugar().Fatalf("could not load config: %v", err)
	}

	listener, err := net.Listen("tcp", conf.Server.Addr)
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     conf.Database.Addr,
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Database: conf.Database.Name,
	})
	defer db.Close()

	// init repository
	repo := logsvc.NewRepository(db)
	// init service handler
	svc := logsvc.NewService(repo)
	svc = logsvc.NewLoggingMiddleware(logger, svc)

	server := grpc.NewServer()
	pb.RegisterLogServiceServer(server, svc)

	logger.With(
		zap.String("transport", "grpc"),
		zap.String("address", conf.Server.Addr),
	).Info("listening")

	if err := server.Serve(listener); err != nil {
		logger.Sugar().Fatalf("failed to serve: %v", err)
	}
}

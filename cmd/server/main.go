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
		// if file doesn't exist, use empty byte slice to
		// pass to the config.Load(). It'll read from the env vars.
		confFile = []byte{}
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

	if err := server.Serve(listener); err != nil {
		logger.Sugar().Fatalf("failed to serve: %v", err)
	}
}

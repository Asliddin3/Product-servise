package main

import (
	"net"

	"github.com/Asliddin3/Product-servise/config"
	pb "github.com/Asliddin3/Product-servise/genproto/product"
	"github.com/Asliddin3/Product-servise/pkg/db"
	"github.com/Asliddin3/Product-servise/pkg/logger"
	"github.com/Asliddin3/Product-servise/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpcclient "github.com/Asliddin3/Product-servise/service/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "")
	defer logger.Cleanup(log)

	log.Info("main:sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("datbase", cfg.PostgresDatabase))
	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("error while connect to clients", logger.Error(err))
	}

	productService := service.NewProductService(grpcClient,connDb, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterProductServiceServer(s, productService)
	log.Info("main: server runing",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

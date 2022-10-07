package main

import (
	"log"
	"net"
	"template-service/pkg/logger"
	"template/Product-servise/pkg/logger"
	"template/Template/pkg/logger"
	"template/Template/service"

	"github.com/Asiddin3/Product-servise/config"
	product "github.com/Asiddin3/Product-servise/genproto"
	"github.com/Asiddin3/Product-servise/pkg/db"
	"github.com/Asliddin3/Product-service/pkg/logger"
	"github.com/Asliddin3/Product-service/service"
	pb "github.com/Asliddin3/Product-servise/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	cfg:=config.Load()

	log:=logger.New(cfg.LogLevel,"")
	defer logger.Cleanup(log)

	log.Info("main:sqlxConfig",
	logger.String("host",cfg.PostgresHost),
	logger.Int("port",cfg.PostgresPort),
	logger.String("datbase",cfg.PostgresDatabase))
	connDb,err:=db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error",logger.Error(err))
	}
	productService:=service.NewUserService
}
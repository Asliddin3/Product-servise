package service

import (
	"context"

	pb "github.com/Asliddin3/Product-service/genproto"
	l "github.com/Asliddin3/Product-service/pkg/logger"
	"github.com/Asliddin3/Product-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct{
	storage storage.IStorage
	logger l.logger
}

func NewProductService(db *sqlx.DB,log l.Logger) *ProductService{
	return &ProductService{
		storage: storage.NewStoragePg(db),
		logger: log,
	}
}

func (s *ProductService) Create(ctx context.Context,req *pb.ProductRequest)(*pb.Product,error){
	productReq,err:=s.storage.Product().Create(req)
	if err != nil {
		s.logger.Error("error creating product",l.Any("error creating product",err))
		return &pb.Product{},status.Error(codes.Internal,"something went internal input")
	}
	return productReq,nil
}
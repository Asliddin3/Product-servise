package service

import (
	"context"

	pb "github.com/Asliddin3/Product-servise/genproto"
	l "github.com/Asliddin3/Product-servise/pkg/logger"
	"github.com/Asliddin3/Product-servise/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct{
	storage storage.IStorage
	logger l.Logger
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
func (s *ProductService) CreateType(ctx context.Context,req *pb.TypeRequest)(*pb.Type,error){
	typeReq,err:=s.storage.Product().CreateType(req)
	if err != nil {
		s.logger.Error("error creating product",l.Any("error creating product",err))
		return &pb.Type{},status.Error(codes.Internal,"something went internal input")
	}
	return typeReq,nil
}
func (s *ProductService) CreateCategory(ctx context.Context,req *pb.CategoryRequest)(*pb.Category,error){
	categoryReq,err:=s.storage.Product().CreateCategory(req)
	if err != nil {
		s.logger.Error("error creating product",l.Any("error creating product",err))
		return &pb.Category{},status.Error(codes.Internal,"something went internal input")
	}
	return categoryReq,nil
}
package service

import (
	"context"
	"fmt"

	pb "github.com/Asliddin3/Product-servise/genproto/product/product"
	l "github.com/Asliddin3/Product-servise/pkg/logger"
	"github.com/Asliddin3/Product-servise/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewProductService(db *sqlx.DB, log l.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
func (s *ProductService) Update(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product, err := s.storage.Product().Update(req)
	if err != nil {
		s.logger.Error("error while updating product", l.Any("error updating product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong")
	}
	return product, nil
}

func (s *ProductService) GetProducts(ctx context.Context, req *pb.Empty) (*pb.Products, error) {
	products, err := s.storage.Product().GetProducts(req)
	fmt.Println(products, err)
	if err != nil {
		s.logger.Error("error while geting products", l.Any("error getting products", err))
		return &pb.Products{}, status.Error(codes.Internal, "something went wrong")
	}
	return products,nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductId) (*pb.ProductResponse, error) {
	product, err := s.storage.Product().GetProduct(req)
	if err != nil {
		s.logger.Error("error while geting product", l.Any("error deleting product", err))
		return &pb.ProductResponse{}, status.Error(codes.Internal, "somethig went wrong")
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.GetProductId) (*pb.Empty, error) {
	_, err := s.storage.Product().DeleteProduct(req)
	if err != nil {
		s.logger.Error("error deleting product", l.Any("error deleting product", err))
		return &pb.Empty{}, status.Error(codes.Internal, "something went wrong invalid argument")
	}
	return &pb.Empty{}, nil
}

func (s *ProductService) Create(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	productReq, err := s.storage.Product().Create(req)
	if err != nil {
		s.logger.Error("error creating product", l.Any("error creating product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went internal input")
	}
	return productReq, nil
}
func (s *ProductService) CreateType(ctx context.Context, req *pb.TypeRequest) (*pb.Type, error) {
	typeReq, err := s.storage.Product().CreateType(req)
	if err != nil {
		s.logger.Error("error creating type", l.Any("error creating type", err))
		return &pb.Type{}, status.Error(codes.Internal, "something went internal input")
	}
	return typeReq, nil
}
func (s *ProductService) CreateCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.Category, error) {
	categoryReq, err := s.storage.Product().CreateCategory(req)
	if err != nil {
		s.logger.Error("error creating category", l.Any("error creating category", err))
		return &pb.Category{}, status.Error(codes.Internal, "something went internal input")
	}
	return categoryReq, nil
}

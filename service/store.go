package service

import (
	"context"

	pb "githab.com/Asliddin3/store-servis/genproto/store"
	"github.com/Asliddin3/store-servis/pkg/logger"
	"github.com/Asliddin3/store-servis/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	grpcclient "github.com/Asliddin3/Product-servise/service/grpc_client"
)
type StoreService struct {
	Client  *grpcclient.ServiceManager
	Storage storage.IsStorage
	Logger  logger.Logger
}


func (s *StoreService) Create(ctx context.Context, req *pb.Store) (*pb.StoreInfo, error) {
	res, err := s.Storage.Store().CreateStore(req)
	if err != nil {
		s.Logger.Error("Error while creating store", logger.Any("insert", err))
		return &pb.StoreInfo{}, status.Error(codes.Internal, "Please recheck store data")
	}

	return res, nil
}

func (s *StoreService) GetStore(c context.Context, req *pb.Id) (*pb.StoreInfo, error) {
	res, err := s.Storage.Store().GetStoreById(req)
	if err != nil {
		s.Logger.Error("Error while getting storeinfo", logger.Any("get", err))
		return res, status.Error(codes.Internal, "Not found")
	}
	return res, nil
}


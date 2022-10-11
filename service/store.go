package service

import (
	"context"

	grpcclient "github.com/Asliddin3/Product-servise/service/grpc_client"
	store "github.com/Asliddin3/store-servis/genproto"
	"github.com/Asliddin3/store-servis/pkg/logger"
	"github.com/Asliddin3/store-servis/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoreService struct {
	Client  *grpcclient.ServiceManager
	Storage storage.IStorage
	Logger  logger.Logger
}

func (s *StoreService) Create(ctx context.Context, req *store.StoreRequest) (*store.StoreResponse, error) {
	res, err := s.Storage.Store().Create(&store.StoreRequest{})
	if err != nil {
		s.Logger.Error("Error while creating store", logger.Any("insert", err))
		return &store.StoreResponse{}, status.Error(codes.Internal, "Please recheck store data")
	}
	return res, nil
}

func (s *StoreService) GetStore(c context.Context, req *store.GetStoreInfoById) (*store.StoreResponse, error) {
	res, err := s.Storage.Store().GetStore(req)
	if err != nil {
		s.Logger.Error("Error while getting storeinfo", logger.Any("get", err))
		return res, status.Error(codes.Internal, "Not found")
	}
	return res, nil
}

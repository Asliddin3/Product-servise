package service

import (
	"context"

	grpcclient "github.com/Asliddin3/Product-servise/service/grpc_client"
	store "github.com/Asliddin3/Product-servise/genproto/store"
	"github.com/Asliddin3/Product-servise/pkg/logger"
	"github.com/Asliddin3/Product-servise/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoreService struct {
	Client  *grpcclient.ServiceManager
	Storage storage.IStorage
	Logger  logger.Logger
}

func (s *StoreService) Create(ctx context.Context, req *store.StoreRequest) (*store.StoreResponse, error) {
	storeReq:=store.StoreRequest{
		Name: req.Name,
	}
	for _,addressReq:=range req.Addresses{
		storeReq.Addresses=append(storeReq.Addresses, &store.Address{
			District: addressReq.District,
			Street: addressReq.Street,
		})
	}
	res, err := s.Client.StoreService().Create(context.Background(),&storeReq)
	if err != nil {
		s.Logger.Error("Error while creating store", logger.Any("insert", err))
		return &store.StoreResponse{}, status.Error(codes.Internal, "Please recheck store data")
	}
	return res, nil
}

func (s *StoreService) GetStore(c context.Context, req *store.GetStoreInfoById) (*store.StoreResponse, error) {
	id:=store.GetStoreInfoById{
		Id: req.Id,
	}
	res, err := s.Client.StoreService().GetStore(context.Background(),&id)
	if err != nil {
		s.Logger.Error("Error while getting storeinfo", logger.Any("get", err))
		return res, status.Error(codes.Internal, "Not found")
	}
	return res, nil
}

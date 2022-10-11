package grpcclient

import (
	"fmt"

	"github.com/Asliddin3/Product-servise/config"
	// productPb "github.com/Asliddin3/Product-servise/genproto/store"
	storePB "github.com/Asliddin3/Product-servise/genproto/store"

	"google.golang.org/grpc"
)

type ServiceManager struct {
	conf           config.Config
	storeServisce  storePB.StoreserviceClient
}

func New(cnfg config.Config) (*ServiceManager, error) {
	connStore, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cnfg.ReviewServiceHost, cnfg.ReviewServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("error while dial product service: host: %s and port: %d",
			cnfg.ReviewServiceHost, cnfg.ReviewServicePort)
	}

	serviceManager := &ServiceManager{
		conf:           cnfg,
		storeServisce: storePB.NewStoreserviceClient(connStore),
	}

	return serviceManager, nil
}

func (s *ServiceManager) ProductService() storePB.StoreserviceClient {
	return s.storeServisce
}

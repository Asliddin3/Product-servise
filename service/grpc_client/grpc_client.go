package grpcclient

import (
	"fmt"

	"github.com/Asliddin3/Product-servise/config"
	productPb "github.com/Asliddin3/Product-servise/genproto"
	"google.golang.org/grpc"
)

type ServiceManager struct {
	conf           config.Config
	productService productPb.ProductServiceClient
}

func NewStore(cnfg config.Config) (*ServiceManager, error) {
	connProduct, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cnfg.ReviewServiceHost, cnfg.ReviewServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("error while dial product service: host: %s and port: %d",
			cnfg.ReviewServiceHost, cnfg.ReviewServicePort)
	}

	serviceManager := &ServiceManager{
		conf:           cnfg,
		productService: productPb.NewProductServiceClient(connProduct),
	}

	return serviceManager, nil
}

func (s *ServiceManager) ProductService() productPb.ProductServiceClient {
	return s.productService
}

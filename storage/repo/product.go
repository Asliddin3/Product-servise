package repo

import (
	pb "github.com/Asiddin3/Product-servise/genproto"
)

type ProductStorageI interface{
	Create(*pb.ProductRequest) (*pb.Product,error)
}
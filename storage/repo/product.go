package repo

import (
	pb "github.com/Asiddin3/Product-servise/genproto"
)

type ProductStorageI interface{
	Create(*pb.ProductRequest) (*pb.Product,error)
	CreateType(*pb.TypeRequest)(*pb.Type,error)
	CreateCategory(*pb.CategoryRequest)(*pb.Category,error)
}
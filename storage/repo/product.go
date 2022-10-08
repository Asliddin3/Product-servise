package repo

import (
	pb "github.com/Asliddin3/Product-servise/genproto"
)

type ProductStorageI interface{
	Create(*pb.ProductRequest) (*pb.Product,error)
	CreateType(*pb.TypeRequest)(*pb.Type,error)
	CreateCategory(*pb.CategoryRequest)(*pb.Category,error)
	DeleteProduct(*pb.GetProductId)(*pb.Empty,error)
	Update(*pb.Product)(*pb.Product,error)
	GetProduct(*pb.GetProductId)(*pb.ProductResponse,error)
	GetProducts(*pb.Empty)(*pb.Products,error)
}
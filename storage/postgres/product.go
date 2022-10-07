package postgres

import (
	pb "github.com/Asiddin3/Product-servise/genproto"
	"github.com/jmoiron/sqlx"
)

type productRepo struct{
	db *sqlx.DB
}


func NewProduct(db *sqlx.DB) *productRepo{
	return &productRepo{db: db}
}

func (r *productRepo) Create(req *pb.ProductRequest) (*pb.Product,error){
	productRepo:=pb.Product{}
	err:=r.db.QueryRow(`
	insert into products(name,categoryid,typeid)
	values($1,$2,$3) returning id,name,categoryid,typeid`,
	req.Name,req.Categoryid,req.Typeid).Scan(
		&productRepo.Id,
		&productRepo.Name,
		&productRepo.Categoryid,
		&productRepo.Typeid,
	)
	if err != nil {
		return &pb.Product{}, err
	}
	return &productRepo,nil

}
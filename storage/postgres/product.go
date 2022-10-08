package postgres

import (
	"fmt"

	pb "github.com/Asliddin3/Product-servise/genproto"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{db: db}
}
func (r *productRepo) CreateType(req *pb.TypeRequest) (*pb.Type, error) {
	typeRepo := pb.Type{}
	err := r.db.QueryRow(
		`insert into types(name)
		values($1) returning id,name`, req.Name,
	).Scan(
		&typeRepo.Id,
		&typeRepo.Name,
	)
	if err != nil {
		return &pb.Type{}, err
	}
	return &typeRepo, nil
}

func (r *productRepo) CreateCategory(req *pb.CategoryRequest) (*pb.Category, error) {
	CategoryRepo := pb.Category{}
	err := r.db.QueryRow(
		`insert into categories(name)
		values($1) returning id,name`, req.Name,
	).Scan(
		&CategoryRepo.Id,
		&CategoryRepo.Name,
	)
	if err != nil {
		fmt.Println(err)
		return &pb.Category{}, err
	}
	return &CategoryRepo, nil
}

func (r *productRepo) Create(req *pb.ProductRequest) (*pb.Product, error) {
	productRepo := pb.Product{}
	err := r.db.QueryRow(`
	insert into products(name,price,categoryid,typeid)
	values($1,$2,$3,$4) returning id,name,price,categoryid,typeid`,
		req.Name, req.Price, req.Categoryid, req.Typeid).Scan(
		&productRepo.Id,
		&productRepo.Name,
		&productRepo.Price,
		&productRepo.Categoryid,
		&productRepo.Typeid,
	)
	if err != nil {
		return &pb.Product{}, err
	}
	return &productRepo, nil
}

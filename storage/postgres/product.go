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
func (r *productRepo) GetProducts(req *pb.Empty) (*pb.Products, error) {
	rows, err := r.db.Query(`
	select
	p.id,
	p.name,
	p.price,
	c.name,
	t.name
	from products p
	inner join categories c
	on c.id=p.categoryid
	inner join types t
	on t.id=p.typeid
	`)
	if err != nil {
		return &pb.Products{}, err
	}
	products := pb.Products{}
	for rows.Next() {
		product := pb.ProductResponse{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Category,
			&product.Type)
		if err != nil {
			return &pb.Products{}, err
		}
		products.Products = append(products.Products, &product)
	}
	return &products, nil
}

func (r *productRepo) GetProduct(req *pb.GetProductId) (*pb.ProductResponse, error) {
	product := pb.ProductResponse{}
	err := r.db.QueryRow(`
	select
	p.name,
	p.price,
	c.name,
	t.name
	from products p
	inner join categories c
	on c.id=p.categoryid
	inner join types t
	on t.id=p.typeid
	where p.id=$1
	`, req.Id).Scan(&product.Name, &product.Price,
		&product.Category,
		&product.Type,
	)
	fmt.Println(err)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return &product, nil
}

func (r *productRepo) Update(req *pb.Product) (*pb.Product, error) {
	_, err := r.db.Exec(`
	update products
	set name=$1,
	price=$2,
	categoryid=$3,
	typeid=$4
	where id=$5`, req.Name,
		req.Price,
		req.Categoryid,
		req.Typeid,
		req.Id)
	if err != nil {
		return &pb.Product{}, err
	}
	return req, nil
}

func (r *productRepo) DeleteProduct(req *pb.GetProductId) (*pb.Empty, error) {
	_, err := r.db.Exec(`
	delete from products
	where id=$1`, req.Id)
	fmt.Println(err)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
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

package postgres

import (
	pb "github.com/Asliddin3/store-servis/genproto"
	"github.com/jmoiron/sqlx"
)

type storeRepo struct {
	db *sqlx.DB
}

func NewStoreRepo(db *sqlx.DB) *storeRepo {
	return &storeRepo{db: db}
}
func (r *storeRepo) Create(req *pb.StoreRequest) (*pb.StoreResponse, error) {
	storeResp := pb.StoreResponse{}
	err := r.db.QueryRow(`
	insert into stores(name) values($1) returning id ,name`, req.Name).Scan(
		&storeResp.Id, &storeResp.Name)
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	addresses := []*pb.AddressResp{}
	for _, address := range req.Addresses {
		addressResp := pb.AddressResp{}
		err = r.db.QueryRow(`
		insert into addresses(district,street) values($1,$2)
		returning id,district,street`, address.District, address.Street).Scan(
			&addressResp.Id, &addressResp.District, &addressResp.Street,
		)
		_, err = r.db.Exec(`
		insert into store_addresses (store_id,address_id)
		values($1,$2)`, storeResp.Id, addressResp.Id)
		if err != nil {
			return &pb.StoreResponse{}, err
		}
		addresses = append(addresses, &addressResp)
	}
	storeResp.Addresses = addresses
	return &storeResp, nil
}
func (r *storeRepo) GetStore(req *pb.GetStoreInfoById) (*pb.StoreResponse, error) {
	store := pb.StoreResponse{}
	err := r.db.QueryRow(`
	select id,name from stores where id=$1`, req.Id).Scan(&store.Id,
		&store.Name)
	rows, err := r.db.Query(`
	select a.id,a.district,a.street
	from addresses a inner join store_addresses sa
	on sa.address_id=a.id and sa.store_id=$1
	`, store.Id)
	addresses := []*pb.AddressResp{}
	for rows.Next() {
		address := pb.AddressResp{}
		err = rows.Scan(&address.Id, &address.District, &address.Street)
		if err != nil {
			return &pb.StoreResponse{}, err
		}
		addresses = append(addresses, &address)
	}
	store.Addresses = addresses
	return &store, nil
}

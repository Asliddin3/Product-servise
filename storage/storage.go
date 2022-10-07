package storage

import (
	"github.com/Asliddin3/Product-servise/storage/postgres"
	"github.com/Asliddin3/Product-servise/storage/repo"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db          *sqlx.DB
	productRepo repo.ProductStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		productRepo: postgres.NewProductRepo(db),
	}
}
func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}

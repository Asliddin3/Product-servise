package db

import (
	"fmt"
	"template/Product-servise/config"

	"github.com/Asiddin3/Product-servise/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDb(cfg config.Config) (*sqlx.DB,error){

}
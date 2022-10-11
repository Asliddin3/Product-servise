package postgres

import (
	"testing"

	"github.com/Asliddin3/Product-servise/config"
	pb "github.com/Asliddin3/Product-servise/genproto/product"
	"github.com/Asliddin3/Product-servise/pkg/db"
	"github.com/Asliddin3/Product-servise/storage/repo"
	"github.com/stretchr/testify/suite"
)

type ProductSuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.ProductStorageI
}

func (suite *ProductSuiteTest) SetupSuite() {
	pgPool, cleanupfunc := db.ConnectToDbForSuite(config.Load())

	suite.Repository = NewProductRepo(pgPool)
	suite.CleanUpFunc = cleanupfunc
}
func (s *ProductSuiteTest) TestProductCrud() {
	productCreate := pb.ProductRequest{
		Name:       "suite product",
		Price:      12300,
		Categoryid: 2,
		Typeid:     1,
	}
	product, err := s.Repository.Create(&productCreate)
	s.Nil(err)
	s.NotNil(product)

	updateProduct := pb.Product{
		Id:         product.Id,
		Name:       "updated user name",
		Price:      30000,
		Categoryid: 2,
		Typeid:     1,
	}
	product, err = s.Repository.Update(&updateProduct)
	s.Nil(err)
	s.NotNil(product)

	getProduct, err := s.Repository.GetProduct(&pb.GetProductId{Id: updateProduct.Id})
	s.Nil(err)
	s.NotNil(getProduct)
	s.Equal(updateProduct.Name, getProduct.Name)
	_, err = s.Repository.DeleteProduct(&pb.GetProductId{Id: updateProduct.Id})
	s.Nil(err)

}

func (suite *ProductSuiteTest) TearDownSuite() {
	suite.CleanUpFunc()
}
func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductSuiteTest))
}

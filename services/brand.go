package services

import (
	"context"
	"database/sql"
	"errors"
	"shop-api/helper"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"

	"github.com/jinzhu/copier"
)

// Brand represents all possible actions available for brand services
type Brand interface {
	Add(context.Context, *types.InputAddBrand) (int64, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, map[string]string, []string, int64, int64) ([]types.OutputBrand, error)
	GetByID(context.Context, int64) (types.OutputBrand, error)
	UpdateByID(context.Context, int64, *types.InputUpdateBrand) error
}

// BrandService defines properties
type BrandService struct {
	Storage storage.Storage
	Orm     helper.OrmInterface
}

// NewBrandService map properties storage and return BrandService
func NewBrandService() (s BrandService) {
	s.Storage = storage.NewStorage()
	s.Orm = helper.NewOrm()
	return
}

// Add method for add a new brand by InputAddBrand
// Validate status and call to storage
func (s BrandService) Add(ctx context.Context, input *types.InputAddBrand) (id int64, err error) {

	var brand = models.Brand{}  // Init variable brand
	copier.Copy(&brand, &input) // Map data input to model

	ormer := s.Orm.NewOrms()                     // Declare ormer
	ormer.BeginTx(ctx, &sql.TxOptions{})         // Begin transaction
	id, err = s.Storage.Brand.Add(ormer, &brand) // Execute method Add
	if id < 1 || err != nil {
		ormer.Rollback()
		return
	}
	err = ormer.Commit()
	return
}

// Delete method delete brand by ID
func (s BrandService) Delete(ctx context.Context, id int64) (err error) {
	errorMessage := "Not found"
	brand := models.Brand{
		ID: id,
	}
	var num int64
	ormer := s.Orm.NewOrms()
	num, err = s.Storage.Brand.Delete(ormer, &brand)
	if num < 1 {
		err = errors.New(errorMessage)
	}
	return
}

// GetByID service for retrieve brand BY ID
func (s BrandService) GetByID(ctx context.Context, id int64) (result types.OutputBrand, err error) {
	ormer := s.Orm.NewOrms()
	brand, err := s.Storage.Brand.GetByID(ormer, id)
	if err != nil {
		return
	}
	copier.Copy(&result, &brand)
	return
}

// GetAll service for retrieves all Brand matches certain condition
func (s BrandService) GetAll(
	ctx context.Context,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (results []types.OutputBrand, err error) {
	ormer := s.Orm.NewOrms()
	categories, err := s.Storage.Brand.GetAll(ormer, query, order, offset, limit)
	copier.Copy(&results, &categories)
	return
}

// UpdateByID service for update brand by ID and InputUpdateBrand
func (s BrandService) UpdateByID(ctx context.Context, id int64, input *types.InputUpdateBrand) (err error) {
	errorMessage := "Not found" // Init error message

	ormer := s.Orm.NewOrms()
	brand, err := s.Storage.Brand.GetByID(ormer, id)

	copier.Copy(&brand, &input) // Map input object to model object
	num, err := s.Storage.Brand.UpdateByID(ormer, &brand)
	if num < 1 || err != nil {
		err = errors.New(errorMessage)
	}
	return
}

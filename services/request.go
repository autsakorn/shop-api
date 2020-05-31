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

// Request represents all possible actions available for request services
type Request interface {
	Add(context.Context, *types.InputAddRequest) (int64, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, map[string]string, []string, int64, int64) ([]types.OutputRequest, error)
	GetByID(context.Context, int64) (types.OutputRequest, error)
	UpdateByID(context.Context, int64, *types.InputUpdateRequest) error
}

// RequestService defines properties
type RequestService struct {
	Storage storage.Storage
	Orm     helper.OrmInterface
}

// NewRequestService map properties storage and return RequestService
func NewRequestService() (s RequestService) {
	s.Storage = storage.NewStorage()
	s.Orm = helper.NewOrm()
	return
}

// Add method for add a new request by InputAddRequest
// Validate status and call to storage
func (s RequestService) Add(ctx context.Context, input *types.InputAddRequest) (id int64, err error) {

	var request = models.Request{} // Init variable request
	copier.Copy(&request, &input)  // Map data input to model

	ormer := s.Orm.NewOrms()             // Declare ormer
	ormer.BeginTx(ctx, &sql.TxOptions{}) // Begin transaction

	brandTitle := input.Brand
	brand, err := s.Storage.Brand.GetByTitle(ormer, brandTitle)
	if brand.ID < 1 || err != nil {
		brandID, errBrandAdd := s.Storage.Brand.Add(ormer, &models.Brand{Title: brandTitle})
		if errBrandAdd != nil {
			return
		}
		request.Brand = &models.Brand{ID: brandID}
	} else {
		request.Brand = &brand
	}

	categoryName := input.Category
	category, err := s.Storage.Category.GetByName(ormer, categoryName)
	if category.ID < 1 || err != nil {
		categoryID, errCategoryID := s.Storage.Category.Add(ormer, &models.Category{Name: categoryName})
		if errCategoryID != nil {
			return
		}
		request.Category = &models.Category{ID: categoryID}
	} else {
		request.Category = &category
	}

	id, err = s.Storage.Request.Add(ormer, &request) // Execute method Add
	if id < 1 || err != nil {
		ormer.Rollback()
		return
	}
	err = ormer.Commit()
	return
}

// GetByID service for retrieve request BY ID
func (s RequestService) GetByID(ctx context.Context, id int64) (result types.OutputRequest, err error) {
	ormer := s.Orm.NewOrms()
	request, err := s.Storage.Request.GetByID(ormer, id)
	if err != nil {
		return
	}
	copier.Copy(&result, &request)
	return
}

// GetAll service for retrieves all Request matches certain condition
func (s RequestService) GetAll(
	ctx context.Context,
	query map[string]string,
	request []string,
	offset int64,
	limit int64,
) (results []types.OutputRequest, err error) {
	ormer := s.Orm.NewOrms()
	categories, err := s.Storage.Request.GetAll(ormer, query, request, offset, limit)
	copier.Copy(&results, &categories)
	return
}

// UpdateByID service for update request by ID and InputUpdateRequest
func (s RequestService) UpdateByID(ctx context.Context, id int64, input *types.InputUpdateRequest) (err error) {
	errorMessage := "Not found" // Init error message

	ormer := s.Orm.NewOrms()
	request, err := s.Storage.Request.GetByID(ormer, id)

	num, err := s.Storage.Request.UpdateByID(ormer, &request)
	if num < 1 || err != nil {
		err = errors.New(errorMessage)
	}
	return
}

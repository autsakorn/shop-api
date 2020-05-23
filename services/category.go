package services

import (
	"context"
	"database/sql"
	"errors"
	"shop-api/helper"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"
	"shop-api/utils"

	"github.com/jinzhu/copier"
)

// Category represents all possible actions available for category services
type Category interface {
	Add(context.Context, *types.InputAddCategory) (int64, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, map[string]string, []string, int64, int64) ([]types.OutputCategory, error)
	GetByID(context.Context, int64) (types.OutputCategory, error)
	UpdateByID(context.Context, int64, *types.InputUpdateCategory) error
}

// CategoryService defines properties
type CategoryService struct {
	Storage storage.Storage
	Orm     helper.OrmInterface
}

// NewCategoryService map properties storage and return CategoryService
func NewCategoryService() (s CategoryService) {
	s.Storage = storage.NewStorage()
	s.Orm = helper.NewOrm()
	return
}

// Add method for add a new category by InputAddCategory
// Validate status and call to storage
func (s CategoryService) Add(ctx context.Context, input *types.InputAddCategory) (id int64, err error) {
	errorMessage := "Please enter valid status, Must be either [Active|Inactive]" // Init error message

	var category = models.Category{} // Init variable category
	copier.Copy(&category, &input)   // Map data input to model

	indexStatus := int32(utils.IndexOf(input.Status, models.CategoryStatus)) // Get index from status(string)
	if indexStatus < 0 {                                                     // Not found index from status(string)
		err = errors.New(errorMessage)
		return
	}
	category.Status = indexStatus                      // Set category status
	ormer := s.Orm.NewOrms()                           // Declare ormer
	ormer.BeginTx(ctx, &sql.TxOptions{})               // Begin transaction
	id, err = s.Storage.Category.Add(ormer, &category) // Execute method Add
	if id < 1 || err != nil {
		ormer.Rollback()
		return
	}
	err = ormer.Commit()
	return
}

// Delete method delete category by ID
func (s CategoryService) Delete(ctx context.Context, id int64) (err error) {
	errorMessage := "Not found"
	category := models.Category{
		ID: id,
	}
	var num int64
	ormer := s.Orm.NewOrms()
	num, err = s.Storage.Category.Delete(ormer, &category)
	if num < 1 {
		err = errors.New(errorMessage)
	}
	return
}

// GetByID service for retrieve category BY ID
func (s CategoryService) GetByID(ctx context.Context, id int64) (result types.OutputCategory, err error) {
	ormer := s.Orm.NewOrms()
	category, err := s.Storage.Category.GetByID(ormer, id)
	if err != nil {
		return
	}
	copier.Copy(&result, &category)
	return
}

// GetAll service for retrieves all Category matches certain condition
func (s CategoryService) GetAll(
	ctx context.Context,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (results []types.OutputCategory, err error) {
	ormer := s.Orm.NewOrms()
	categories, err := s.Storage.Category.GetAll(ormer, query, order, offset, limit)
	copier.Copy(&results, &categories)
	return
}

// UpdateByID service for update category by ID and InputUpdateCategory
func (s CategoryService) UpdateByID(ctx context.Context, id int64, input *types.InputUpdateCategory) (err error) {
	errorMessage := "Not found" // Init error message
	ormer := s.Orm.NewOrms()
	category, err := s.Storage.Category.GetByID(ormer, id)
	m := models.Category{
		ID:        id,
		Name:      input.Name,
		Detail:    input.Detail,
		CreatedAt: category.CreatedAt,
	}
	num, err := s.Storage.Category.UpdateByID(ormer, &m)
	if num < 1 || err != nil {
		err = errors.New(errorMessage)
	}
	return
}

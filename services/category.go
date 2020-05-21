package services

import (
	"errors"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"
	"shop-api/utils"

	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/copier"
)

// Category represents all possible actions available for category services
type Category interface {
	Add(ormer orm.Ormer, input *types.InputAddCategory) (responseCode int, id int64, err error)
	Delete(ormer orm.Ormer, id int64) (responseCode int, err error)
	GetAll(ormer orm.Ormer,
		query map[string]string,
		order []string,
		offset int64,
		limit int64,
	) (responseCode int, results []types.OutputCategory, err error)
	GetByID(ormer orm.Ormer, id int64) (responseCode int, result types.OutputCategory, err error)
	UpdateByID(ormer orm.Ormer, id int64, category *types.InputUpdateCategory) (responseCode int, err error)
}

// CategoryService defines properties
type CategoryService struct {
	Storage storage.Storage
	Orm     orm.Ormer
}

// NewCategoryService map properties storage and return CategoryService
func NewCategoryService() (s CategoryService) {
	s.Storage = storage.NewStorage()
	return
}

// Add method for add a new category by InputAddCategory
// Validate status and call to storage
func (s CategoryService) Add(ormer orm.Ormer, input *types.InputAddCategory) (responseCode int, id int64, err error) {
	errorMessage := "Please enter valid status, Must be either [Active|Inactive]"
	responseCode = types.ResponseCode["BadRequest"]
	var category = models.Category{}
	// Map data input to model
	copier.Copy(&category, &input)
	indexStatus := int32(utils.IndexOf(input.Status, models.CategoryStatus))
	if indexStatus < 0 {
		err = errors.New(errorMessage)
		ormer.Rollback()
		return
	}
	category.Status = indexStatus
	// Execute method Add
	id, err = s.Storage.Category.Add(ormer, &category)
	if err != nil {
		ormer.Rollback()
		return
	}
	err = ormer.Commit()
	if id > 0 && err == nil {
		responseCode = types.ResponseCode["CreatedSuccess"]
	}
	return
}

// Delete method delete category by ID
func (s CategoryService) Delete(ormer orm.Ormer, id int64) (responseCode int, err error) {
	errorMessage := "Not found"
	responseCode = types.ResponseCode["Success"]
	category := models.Category{
		ID: id,
	}
	var num int64
	num, err = s.Storage.Category.Delete(ormer, &category)
	if num < 1 {
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetByID service for retrieve category BY ID
func (s CategoryService) GetByID(ormer orm.Ormer, id int64) (responseCode int, result types.OutputCategory, err error) {
	category, err := s.Storage.Category.GetByID(ormer, id)
	copier.Copy(&result, &category)
	responseCode = types.ResponseCode["Success"]
	return
}

// GetAll service for retrieves all Category matches certain condition
func (s CategoryService) GetAll(
	ormer orm.Ormer,
	query map[string]string,
	order []string,
	offset int64,
	limit int64,
) (responseCode int, results []types.OutputCategory, err error) {
	categories, err := s.Storage.Category.GetAll(ormer, query, order, offset, limit)
	copier.Copy(&results, &categories)
	responseCode = types.ResponseCode["Success"]
	return
}

// UpdateByID service for update category by ID and InputUpdateCategory
func (s CategoryService) UpdateByID(ormer orm.Ormer, id int64, category *types.InputUpdateCategory) (responseCode int, err error) {
	errorMessage := "Not found"
	responseCode = types.ResponseCode["Success"]
	m := models.Category{
		ID:     id,
		Name:   category.Name,
		Detail: category.Detail,
	}
	num, err := s.Storage.Category.UpdateByID(ormer, &m)
	if num < 1 {
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

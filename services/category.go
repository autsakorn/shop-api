package services

import (
	"errors"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/types"
	"shop-api/utils"

	"github.com/jinzhu/copier"
)

// Category defines interface
type Category interface {
	Add(input *types.InputAddCategory) (responseCode int, id int64, err error)
	Delete(id int64) (responseCode int, err error)
	GetAll(query map[string]string, order []string,
		offset int64, limit int64) (responseCode int, results []types.OutputCategory, err error)
	GetByID(id int64) (responseCode int, result types.OutputCategory, err error)
	UpdateByID(id int64, category *types.InputUpdateCategory) (responseCode int, err error)
}

// CategoryService ...
type CategoryService struct {
	Storage storage.Storage
}

// NewCategoryService ...
func NewCategoryService() (s CategoryService) {
	s.Storage = storage.NewStorage()
	return
}

// Add ...
func (s CategoryService) Add(input *types.InputAddCategory) (responseCode int, id int64, err error) {
	errorMessage := "Please enter valid status, Must be either [Active|Inactive]"
	responseCode = types.ResponseCode["BadRequest"]
	var category = models.Category{}
	// Map data input to model
	copier.Copy(&category, &input)
	indexStatus := int32(utils.IndexOf(input.Status, models.CategoryStatus))
	if indexStatus < 0 {
		err = errors.New(errorMessage)
		return
	}
	category.Status = indexStatus
	// Execute method Add
	id, err = s.Storage.Category.Add(&category)
	if id > 0 {
		responseCode = types.ResponseCode["CreatedSuccess"]
		return
	}
	return
}

// Delete ...
func (s CategoryService) Delete(id int64) (responseCode int, err error) {
	errorMessage := "Not found"
	responseCode = types.ResponseCode["Success"]
	category := models.Category{
		ID: id,
	}
	var num int64
	num, err = s.Storage.Category.Delete(&category)
	if num < 1 {
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

// GetByID ...
func (s CategoryService) GetByID(id int64) (responseCode int, result types.OutputCategory, err error) {
	category, err := s.Storage.Category.GetByID(id)
	copier.Copy(&result, &category)
	responseCode = types.ResponseCode["Success"]
	return
}

// GetAll ...
func (s CategoryService) GetAll(query map[string]string, order []string,
	offset int64, limit int64) (responseCode int, results []types.OutputCategory, err error) {
	categories, err := s.Storage.Category.GetAll(query, order, offset, limit)
	copier.Copy(&results, &categories)
	responseCode = types.ResponseCode["Success"]
	return
}

// UpdateByID ...
func (s CategoryService) UpdateByID(id int64, category *types.InputUpdateCategory) (responseCode int, err error) {
	errorMessage := "Not found"
	responseCode = types.ResponseCode["Success"]
	m := models.Category{
		ID:     id,
		Name:   category.Name,
		Detail: category.Detail,
	}
	num, err := s.Storage.Category.UpdateByID(&m)
	if num < 1 {
		err = errors.New(errorMessage)
		responseCode = types.ResponseCode["BadRequest"]
	}
	return
}

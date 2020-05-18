package storage

// Storage ...
type Storage struct {
	Category Category
	Product  Product
}

// NewStorage ...
func NewStorage() (storage Storage) {
	storage.Category = NewCategoryStorage()
	storage.Product = NewProductStorage()
	return
}

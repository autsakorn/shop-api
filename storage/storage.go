package storage

// Storage defines properties storage
// We need to add when we have new model
type Storage struct {
	Brand    Brand
	Category Category
	Product  Product
	Client   Client
	Request  Request
}

// NewStorage return out storage
func NewStorage() (storage Storage) {
	storage.Brand = NewBrandStorage()
	storage.Category = NewCategoryStorage()
	storage.Product = NewProductStorage()
	storage.Client = NewClientStorage()
	storage.Request = NewRequestStorage()
	return
}

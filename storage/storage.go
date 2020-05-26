package storage

// Storage defines properties storage
// We need to add when we have new model
type Storage struct {
	Category Category
	Product  Product
	Client   Client
}

// NewStorage return out storage
func NewStorage() (storage Storage) {
	storage.Category = NewCategoryStorage()
	storage.Product = NewProductStorage()
	storage.Client = NewClientStorage()
	return
}

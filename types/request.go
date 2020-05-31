package types

import "time"

// InputAddRequest defines properties input add order
type InputAddRequest struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Category        string    `json:"category"`
	Brand           string    `json:"brand"`
	DueDate         time.Time `json:"dueDate"` // Example 2020-05-25T14:14:26Z
	RequirePackage  string    `json:"requirePackage"`
	Remark          string    `json:"remark"`
	ShippingAddress string    `json:"shippingAddress"`
	ShippingMethod  string    `json:"shippingMethod"`
	Country         string    `json:"country"`
}

// InputUpdateRequest defines properties input update order
type InputUpdateRequest struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Category        string    `json:"category"`
	Brand           string    `json:"brand"`
	DueDate         time.Time `json:"dueDate"` // Example 2020-05-25T14:14:26Z
	RequirePackage  string    `json:"requirePackage"`
	Remark          string    `json:"remark"`
	ShippingAddress string    `json:"shippingAddress"`
	ShippingMethod  string    `json:"shippingMethod"`
	Country         string    `json:"country"`
}

// OutputRequest defines properties output order
type OutputRequest struct {
	ID              int64          `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Price           float64        `json:"price"`
	Category        OutputCategory `json:"Category"`
	Brand           OutputBrand    `json:"Brand"`
	DueDate         time.Time      `json:"dueDate"` // Example 2020-05-25T14:14:26Z
	RequirePackage  string         `json:"requirePackage"`
	Remark          string         `json:"remark"`
	ShippingAddress string         `json:"shippingAddress"`
	ShippingMethod  string         `json:"shippingMethod"`
	Country         string         `json:"country"`
}

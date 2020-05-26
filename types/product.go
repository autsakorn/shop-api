package types

// productCategory defines properties category input
type productCategory struct {
	ID int64 `json:"id"`
}

// InputAddProduct defines properties of input new product
type InputAddProduct struct {
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Detail   string          `json:"detail"`
	Brand    string          `json:"brand"`
	Model    string          `json:"model"`
	Stock    int             `json:"stock"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// InputUpdateProduct defines properties of input update and upsert product
type InputUpdateProduct struct {
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Detail   string          `json:"detail"`
	Brand    string          `json:"brand"`
	Model    string          `json:"model"`
	Stock    int             `json:"stock"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// OutputProduct defines properties output product
type OutputProduct struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Price    float64        `json:"price"`
	Detail   string         `json:"detail"`
	Brand    string         `json:"brand"`
	Model    string         `json:"model"`
	Stock    int            `json:"stock"`
	Category OutputCategory `json:"Category"`
}

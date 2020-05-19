package types

// Response ...
type Response struct {
	Message string `json:"message"`
}

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
	Quantity int             `json:"quantity"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// InputDeleteProduct defines properties of input delete product
type InputDeleteProduct struct {
	ID int64 `json:"id"`
}

// InputUpdateProduct defines properties of input update and upsert product
type InputUpdateProduct struct {
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Detail   string          `json:"detail"`
	Brand    string          `json:"brand"`
	Model    string          `json:"model"`
	Quantity int             `json:"quantity"`
	Cost     float64         `json:"cost"`
	Category productCategory `json:"Category"`
}

// OutputProduct ...
type OutputProduct struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Price    float64        `json:"price"`
	Detail   string         `json:"detail"`
	Brand    string         `json:"brand"`
	Model    string         `json:"model"`
	Quantity int            `json:"quantity"`
	Cost     float64        `json:"cost"`
	Category OutputCategory `json:"Category"`
}

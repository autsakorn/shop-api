package types

// InputAddCategory defines properties input add category
type InputAddCategory struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

// InputUpdateCategory defines properties input update category
type InputUpdateCategory struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

// OutputCategory defines properties output category
type OutputCategory struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	StatusRes string `json:"status"`
}

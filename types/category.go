package types

// InputAddCategory ...
type InputAddCategory struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

// InputUpdateCategory ...
type InputUpdateCategory struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

// OutputCategory ...
type OutputCategory struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	StatusRes string `json:"status"`
}

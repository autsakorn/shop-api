package types

// InputAddBrand defines properties input add category
type InputAddBrand struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

// InputUpdateBrand defines properties input update category
type InputUpdateBrand struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

// OutputBrand defines properties output category
type OutputBrand struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

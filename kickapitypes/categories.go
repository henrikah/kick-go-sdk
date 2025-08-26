package kickapitypes

type Category struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type GetCategoriesResponse struct {
	Data    []Category `json:"data,omitempty"`
	Message string     `json:"message,omitempty"`
}

type GetCategoryResponse struct {
	Data    Category `json:"data,omitempty"`
	Message string   `json:"message,omitempty"`
}

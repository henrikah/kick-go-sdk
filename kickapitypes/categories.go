package kickapitypes

type Categories struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Category struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Tags        []string `json:"tags"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	ViewerCount int      `json:"viewer_count"`
}

type GetCategoriesResponse struct {
	Data    []Categories `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
}

type GetCategoryResponse struct {
	Data    Category `json:"data"`
	Message string   `json:"message,omitempty"`
}

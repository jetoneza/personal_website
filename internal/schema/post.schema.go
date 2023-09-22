package schema

type CreatePostSchema struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Category  string `json:"category,omitempty"`
	Published bool   `json:"published,omitempty"`
}

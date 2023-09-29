package schema

type CreatePostSchema struct {
	Title           string `json:"title" validate:"required"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Content         string `json:"content" validate:"required"`
	Category        string `json:"category"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	Published       bool   `json:"published"`
}

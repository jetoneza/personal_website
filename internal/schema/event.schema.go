package schema

type CreateEventSchema struct {
	Title  string `json:"title" validate:"required"`
	Type   string `json:"type"`
	AllDay string `json:"all_day"`
	Notes  string `json:"notes"`
	Start  string `json:"start" validate:"required"`
	End    string `json:"end" validate:"required"`
}

type UpdateEventSchema struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	AllDay string `json:"all_day"`
	Notes  string `json:"notes"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

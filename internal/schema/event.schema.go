package schema

type CreateEventSchema struct {
	Title  string `json:"title" validate:"required"`
	Type   string `json:"type"`
	AllDay string `json:"all_day"`
	Notes  string `json:"notes" validate:"required"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

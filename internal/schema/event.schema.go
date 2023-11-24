package schema

type CreateEventSchema struct {
	Title  string `json:"title" validate:"required"`
	AllDay string `json:"all_day"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

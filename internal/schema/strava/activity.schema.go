package strava

type CreateActivity struct {
	Name           string `json:"name" validate:"required"`
	Type           string `json:"type" validate:"required"`
	SportType      string `json:"sport_type" validate:"required"`
	Trainer        string `json:"trainer" validate:"required"`
	Distance       string `json:"distance" validate:"required"`
	ElapsedTime    string `json:"elapsed_time" validate:"required"`
	StartDateLocal string `json:"start_date_local" validate:"required"`
}

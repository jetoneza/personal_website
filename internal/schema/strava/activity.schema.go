package strava

type Activity struct {
	Name           string `json:"name" validate:"required"`
	Type           string `json:"type"`
	SportType      string `json:"sport_type"`
	Trainer        string `json:"trainer"`
	Distance       int16  `json:"distance"`
	ElapsedTime    int16  `json:"elapsed_time"`
	StartDateLocal string `json:"start_date_local"`
}

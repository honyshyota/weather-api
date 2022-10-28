package models

type SwaggerUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SwaggerLoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SwaggerCity struct {
	Name string `json:"fav_city,omitempty"`
}

type SwaggerFullForecast struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type SwaggerShortForecastCity struct {
	Name string `json:"name"`
}

package models

import "time"

type City struct {
	Name    string  `json:"name,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
	Lon     float64 `json:"lon,omitempty"`
	Country string  `json:"country,omitempty"`
}

type CityArray []struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

type FullForecastResponse struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type FullForecast struct {
	Cod     string `json:"cod,omitempty"`
	Message int    `json:"message,omitempty"`
	Cnt     int    `json:"cnt,omitempty"`
	List    []struct {
		Dt   int `json:"dt,omitempty"`
		Main struct {
			Temp      float64 `json:"temp,omitempty"`
			FeelsLike float64 `json:"feels_like,omitempty"`
			TempMin   float64 `json:"temp_min,omitempty"`
			TempMax   float64 `json:"temp_max,omitempty"`
			Pressure  int     `json:"pressure,omitempty"`
			SeaLevel  int     `json:"sea_level,omitempty"`
			GrndLevel int     `json:"grnd_level,omitempty"`
			Humidity  int     `json:"humidity,omitempty"`
			TempKf    float64 `json:"temp_kf,omitempty"`
		} `json:"main,omitempty"`
		Weather []struct {
			ID          int    `json:"id,omitempty"`
			Main        string `json:"main,omitempty"`
			Description string `json:"description,omitempty"`
			Icon        string `json:"icon,omitempty"`
		} `json:"weather,omitempty"`
		Clouds struct {
			All int `json:"all,omitempty"`
		} `json:"clouds,omitempty"`
		Wind struct {
			Speed float64 `json:"speed,omitempty"`
			Deg   int     `json:"deg,omitempty"`
			Gust  float64 `json:"gust,omitempty"`
		} `json:"wind,omitempty"`
		Visibility int     `json:"visibility,omitempty"`
		Pop        float64 `json:"pop,omitempty"`
		Sys        struct {
			Pod string `json:"pod,omitempty"`
		} `json:"sys,omitempty"`
		DtTxt string `json:"dt_txt,omitempty"`
		Rain  struct {
			ThreeH float64 `json:"3h,omitempty"`
		} `json:"rain,omitempty"`
	} `json:"list,omitempty"`
	City struct {
		Name  string `json:"name,omitempty"`
		Coord struct {
			Lat float64 `json:"lat,omitempty"`
			Lon float64 `json:"lon,omitempty"`
		} `json:"coord,omitempty"`
		Country string `json:"country,omitempty"`
	} `json:"city,omitempty"`
}

type CompleteWeather struct {
	Weather FullForecast `json:"weather"`
	Temp    float64      `json:"temp" db:"temp"`
	Date    time.Time    `json:"date"`
	Data    []byte       `json:"data" db:"data"`
}

type CityNamesResponse []struct {
	Name string `json:"name"`
}

type ShortForecastResponse struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	AvgTemp  float64  `json:"avg_temp"`
	DateList []string `json:"date_list"`
}

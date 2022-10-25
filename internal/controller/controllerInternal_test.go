package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/honyshyota/weather-api/internal/models"
	mock "github.com/honyshyota/weather-api/internal/usecase/mock"
	"github.com/stretchr/testify/assert"
)

func TestController_getCityList(t *testing.T) {
	type mockBehavior func(s *mock.MockUsecase)

	testCases := []struct {
		name                string
		cities              []string
		mockBehavior        mockBehavior
		expectedCode        int
		expectedRequestBody string
	}{
		{
			name:   "valid",
			cities: []string{"London", "Moscow", "Berlin"},
			mockBehavior: func(s *mock.MockUsecase) {
				s.EXPECT().GetAllCities().Return([]string{
					"London",
					"Moscow",
					"Berlin",
				}, nil)
			},
			expectedCode:        http.StatusOK,
			expectedRequestBody: `[{"name":"London"},{"name":"Moscow"},{"name":"Berlin"}]` + "\n",
		},
		{
			name:   "invalid",
			cities: []string{},
			mockBehavior: func(s *mock.MockUsecase) {
				s.EXPECT().GetAllCities().Return(nil, errors.New("error"))
			},
			expectedCode:        http.StatusInternalServerError,
			expectedRequestBody: `{"error":"internal error"}` + "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock.NewMockUsecase(c)
			tc.mockBehavior(usecase)

			controller := NewController(usecase, sessions.NewCookieStore([]byte("some_key")))

			r := mux.NewRouter()
			r.HandleFunc("/", controller.getCityList)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Equal(t, tc.expectedRequestBody, w.Body.String())
		})
	}
}

func TestController_getShortForecast(t *testing.T) {
	type mockBehavior func(s *mock.MockUsecase, city models.City)

	testCases := []struct {
		name                string
		input               string
		model               models.City
		mockBehavior        mockBehavior
		expectedCode        int
		expectedRequestBody string
	}{
		{
			name:  "valid",
			input: `{"name":"London"}`,
			model: models.City{
				Name: "London",
			},
			mockBehavior: func(s *mock.MockUsecase, city models.City) {
				s.EXPECT().GetShortForecast(city.Name).Return(&models.ShortForecastResponse{
					Name:     "London",
					Country:  "GB",
					AvgTemp:  3.4,
					DateList: []string{"2022-10-29 00:00:00", "2022-10-29 03:00:00", "2022-10-29 06:00:00", "2022-10-29 09:00:00"},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedRequestBody: `{"name":"London","country":"GB","avg_temp":3.4,` +
				`"date_list":["2022-10-29 00:00:00","2022-10-29 03:00:00","2022-10-29 06:00:00","2022-10-29 09:00:00"]}` + "\n",
		},
		{
			name:  "invalid",
			input: `{"name":"Londo"}`,
			model: models.City{
				Name: "Londo",
			},
			mockBehavior: func(s *mock.MockUsecase, city models.City) {
				s.EXPECT().GetShortForecast(city.Name).Return(nil, errors.New("error"))
			},
			expectedCode:        http.StatusBadRequest,
			expectedRequestBody: `{"error":"incorrect input"}` + "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock.NewMockUsecase(c)
			tc.mockBehavior(usecase, tc.model)

			controller := NewController(usecase, sessions.NewCookieStore([]byte("some_key")))

			r := mux.NewRouter()
			r.HandleFunc("/short_forecast", controller.getShortForecast)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/short_forecast", bytes.NewBufferString(tc.input))

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Equal(t, tc.expectedRequestBody, w.Body.String())
		})
	}
}

func TestController_getFullForecast(t *testing.T) {
	type mockBehavior func(s *mock.MockUsecase, city models.FullForecastResponse)

	testCases := []struct {
		name                string
		input               string
		model               models.FullForecastResponse
		mockBehavior        mockBehavior
		expectedCode        int
		expectedRequestBody string
	}{
		{
			name:  "valid",
			input: `{"name":"London","date":"2022-10-29 00:00:00"}`,
			model: models.FullForecastResponse{
				Name: "London",
				Date: "2022-10-29 00:00:00",
			},
			mockBehavior: func(s *mock.MockUsecase, city models.FullForecastResponse) {
				s.EXPECT().GetForecastForTime(city.Name, city.Date).Return(&models.FullForecast{
					Cod:     "0",
					Message: 0,
					Cnt:     0,
					City: struct {
						Name  string `json:"name,omitempty"`
						Coord struct {
							Lat float64 `json:"lat,omitempty"`
							Lon float64 `json:"lon,omitempty"`
						} `json:"coord,omitempty"`
						Country string `json:"country,omitempty"`
					}{
						Name:    "London",
						Country: "GB",
					},
				}, nil)
			},
			expectedCode:        http.StatusOK,
			expectedRequestBody: `{"cod":"0","city":{"name":"London","coord":{},"country":"GB"}}` + "\n",
		},
		{
			name:  "invalid",
			input: `{"name":""}`,
			model: models.FullForecastResponse{},
			mockBehavior: func(s *mock.MockUsecase, city models.FullForecastResponse) {
				s.EXPECT().GetForecastForTime(city.Name, city.Date).Return(nil, errors.New("error"))
			},
			expectedCode:        http.StatusBadRequest,
			expectedRequestBody: `{"error":"incorrect input"}` + "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock.NewMockUsecase(c)
			tc.mockBehavior(usecase, tc.model)

			controller := NewController(usecase, sessions.NewCookieStore([]byte("some_key")))

			r := mux.NewRouter()
			r.HandleFunc("/full_forecast", controller.getFullForecastForDate)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/full_forecast", bytes.NewBufferString(tc.input))

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Equal(t, tc.expectedRequestBody, w.Body.String())
		})
	}
}

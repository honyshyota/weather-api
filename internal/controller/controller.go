package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/honyshyota/weather-api/internal/usecase"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "weather_api"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestId
)

var (
	errIncorrectInput       = errors.New("incorrect input")
	errIncorrectEmailOrPass = errors.New("incorrect email or password")
	errInternalServer       = errors.New("internal error")
)

type ctxKey int8

type Controller struct {
	usecase usecase.Usecase
	logger  *logrus.Logger
	session sessions.Store
}

func NewController(usecase usecase.Usecase, session sessions.Store) *Controller {
	return &Controller{
		usecase: usecase,
		logger:  logrus.New(),
		session: session,
	}
}

func (c *Controller) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (c *Controller) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.logger.Infof("[router] Started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		c.logger.Infof(
			"[router] completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

// @Summary Create User
// @Tags users
// @Description create user
// @Accept json
// @Produce json
// @Param input body models.SwaggerUser true "user info"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /users [post]
func (c *Controller) usersCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	err := c.usecase.CreateUser(&user)
	if err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	userRespond, err := c.usecase.FindUser(user.Name)
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, errInternalServer)
		return
	}

	c.respond(w, r, http.StatusCreated, userRespond)
}

// @Summary Login
// @Tags users
// @Description login user
// @Accept json
// @Produce json
// @Param input body models.SwaggerLoginUser true "user info"
// @Success 202
// @Failure 400
// @Failure 500
// @Router /sessions [post]
func (c *Controller) sessionsCreate(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		c.error(w, r, http.StatusBadRequest, err)
		return
	}

	u, err := c.usecase.FindUser(user.Name)
	if err != nil || user.Password != u.Password {
		c.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPass)
		return
	}

	session, err := c.session.Get(r, sessionName)
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, errInternalServer)
		return
	}

	session.Values["user_name"] = u.Name

	if err := c.session.Save(r, w, session); err != nil {
		c.error(w, r, http.StatusInternalServerError, errInternalServer)
		return
	}

	c.respond(w, r, http.StatusAccepted, w.Header().Values("Set-Cookie"))
}

func (c *Controller) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := c.session.Get(r, sessionName)
		if err != nil {
			c.error(w, r, http.StatusInternalServerError, err)
			return
		}

		userName, ok := session.Values["user_name"]

		if !ok {
			c.error(w, r, http.StatusUnauthorized, errIncorrectInput)
			return
		}

		u, err := c.usecase.FindUser(userName.(string))

		if err != nil {
			c.error(w, r, http.StatusUnauthorized, errIncorrectInput)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

// @Summary Update Favorite City
// @Tags users
// @Description update favorite city
// @Accept json
// @Produce json
// @Param input body models.SwaggerCity true "city info"
// @Success 202
// @Failure 400
// @Router /private/fav_city [post]
func (c *Controller) createFavCity(w http.ResponseWriter, r *http.Request) {
	favCity := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(favCity); err != nil {
		c.error(w, r, http.StatusBadRequest, err)
		return
	}

	user := r.Context().Value(ctxKeyUser).(*models.User)

	err := c.usecase.UpdateFavCity(favCity.FavCity, user.Name)
	if err != nil {
		c.error(w, r, http.StatusBadRequest, err)
		return
	}

	u, err := c.usecase.FindUser(user.Name)
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, errInternalServer)
		return
	}

	u.FavCity = favCity.FavCity

	c.respond(w, r, http.StatusCreated, u)
}

// @Summary Get City List
// @Tags Weather
// @Description list of available cities
// @Produce json
// @Success 200
// @Failure 500
// @Router / [get]
func (c *Controller) getCityList(w http.ResponseWriter, r *http.Request) {
	result := models.CityNamesResponse{}

	cities, err := c.usecase.GetAllCities()
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, errInternalServer)
		return
	}

	for _, city := range cities {
		cityName := models.CityNamesResponse{
			struct {
				Name string `json:"name"`
			}{
				Name: city,
			},
		}

		result = append(result, cityName...)
	}

	c.respond(w, r, http.StatusOK, result)
}

// @Summary Forecast for dates
// @Tags Weather
// @Description getting a forecast for the requested city
// @Accept json
// @Produce json
// @Param input body models.SwaggerShortForecastCity true "city info"
// @Success 200
// @Failure 400
// @Router /short_forecast [post]
func (c *Controller) getShortForecast(w http.ResponseWriter, r *http.Request) {
	city := &models.City{}

	if err := json.NewDecoder(r.Body).Decode(city); err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	shortForecast, err := c.usecase.GetShortForecast(city.Name)
	if err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	c.respond(w, r, http.StatusOK, shortForecast)
}

// @Summary Full forecast for date
// @Tags Weather
// @Description getting a forecast for the requested city and date
// @Accept json
// @Produce json
// @Param input body models.SwaggerFullForecast true "city and date input"
// @Success 200
// @Failure 400
// @Router /full_forecast [post]
func (c *Controller) getFullForecastForDate(w http.ResponseWriter, r *http.Request) {
	weather := &models.FullForecastResponse{}

	if err := json.NewDecoder(r.Body).Decode(weather); err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	forecast, err := c.usecase.GetForecastForTime(weather.Name, weather.Date)
	if err != nil {
		c.error(w, r, http.StatusBadRequest, errIncorrectInput)
		return
	}

	c.respond(w, r, http.StatusOK, forecast)
}

func (c *Controller) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (c *Controller) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	GetCityURI      string
	GetForecastURI  string
	CityList        []string
	UpdateTime      int
	PortApp         string
	ReadTO          time.Duration
	WriteTO         time.Duration
	HostDB          string
	NameDB          string
	User            string
	Password        string
	PortDB          string
	PgMigrationPath string
	PsqlConnect     string
	WeatherToken    string
	SessionKey      string
	DbConn          *sqlx.Conn
	HttpClient      *http.Client
}

func NewConfig() *Config {
	if err := godotenv.Load("config/conf.env"); err != nil {
		logrus.Errorln("[config] Cannot load conf.env, ", err)
		return nil
	}

	readTemp := os.Getenv("READ_TO")
	readTO, err := time.ParseDuration(readTemp)
	if err != nil {
		return nil
	}

	writeTemp := os.Getenv("WRITE_TO")
	writeTO, err := time.ParseDuration(writeTemp)
	if err != nil {
		return nil
	}

	getForecastURI := os.Getenv("GET_FORECAST_URI")
	getCityURI := os.Getenv("GET_CITY_URY")
	pgMigrationPath := os.Getenv("PG_MIGRATIONS_PATH")
	updTime, err := strconv.Atoi(os.Getenv("UPD_TIME"))
	if err != nil {
		logrus.Println("[config] Incorrect data UPD_TIME in .env")
	}
	cityList := strings.Split(os.Getenv("CITY_LIST"), ", ")
	portApp := os.Getenv("PORT_APP")
	hostDB := os.Getenv("HOST_DB")
	nameDB := os.Getenv("NAME_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD")
	portDB := os.Getenv("PORT_DB")
	psqlConnect := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, password, hostDB, portDB, nameDB)
	weatherToken := os.Getenv("WEATHER_TOKEN")
	sessionKey := os.Getenv("SESSION_KEY")
	httpClient := &http.Client{}

	db, err := sqlx.Connect("pgx", psqlConnect)
	if err != nil {
		logrus.Errorln("[config] Cannot connect to DB, ", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Errorln("[config] Failed ping DB, ", err)
		return nil
	}

	conn, err := db.Connx(context.Background())
	if err != nil {
		logrus.Errorln("[config] Cannot connect to DB, ", err)
		return nil
	}

	logrus.Println("[config] DB connection OK")

	m, err := migrate.New(
		pgMigrationPath,
		psqlConnect,
	)
	if err != nil {
		logrus.Errorln("[config] Error create migrate instanse, ", err)
		return nil
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Errorln("[config] Failed migrations, ", err)
		return nil
	}

	logrus.Println("[config] Migrate done")

	config := &Config{
		GetCityURI:      getCityURI,
		GetForecastURI:  getForecastURI,
		UpdateTime:      updTime,
		ReadTO:          readTO,
		WriteTO:         writeTO,
		CityList:        cityList,
		PortApp:         portApp,
		HostDB:          hostDB,
		NameDB:          nameDB,
		User:            user,
		Password:        password,
		PortDB:          portDB,
		PgMigrationPath: pgMigrationPath,
		PsqlConnect:     psqlConnect,
		WeatherToken:    weatherToken,
		SessionKey:      sessionKey,
		DbConn:          conn,
		HttpClient:      httpClient,
	}

	return config
}

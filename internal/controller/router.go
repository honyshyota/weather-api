package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/usecase"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/honyshyota/weather-api/docs"
)

type Server struct {
	*http.Server
	shutdownReq chan bool
	reqCount    uint32
}

func Build(usecase usecase.Usecase, config *config.Config) {
	srv := NewServer(config)

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	controller := NewController(usecase, sessionStore)

	router := mux.NewRouter()
	router.Use(controller.setRequestID)
	router.Use(controller.logRequest)
	router.HandleFunc("/", controller.getCityList).Methods("GET")
	router.HandleFunc("/short_forecast", controller.getShortForecast).Methods("POST")
	router.HandleFunc("/full_forecast", controller.getFullForecastForDate).Methods("POST")
	router.HandleFunc("/users", controller.usersCreate).Methods("POST")
	router.HandleFunc("/sessions", controller.sessionsCreate).Methods("POST")

	private := router.PathPrefix("/private").Subrouter()
	private.Use(controller.authenticateUser)
	private.HandleFunc("/fav_city", controller.createFavCity).Methods("POST")
	private.HandleFunc("/shutdown", srv.ShutdownHandler)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	srv.Handler = router

	done := make(chan bool)
	go func() {
		logrus.Infoln("[server] Server is runnig")
		err := srv.ListenAndServe()
		if err != nil {
			logrus.Errorln("[server] Listen and serve: ", err)
		}
		done <- true
	}()

	//wait shutdown
	srv.WaitShutdown()

	<-done
	logrus.Infoln("[server] Gracefully shutting down")
}

func NewServer(config *config.Config) *Server {
	srv := &Server{
		Server: &http.Server{
			Addr:         config.PortApp,
			ReadTimeout:  config.ReadTO,
			WriteTimeout: config.WriteTO,
		},
		shutdownReq: make(chan bool),
	}
	return srv
}

func (s *Server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		logrus.Infof("[server] Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		logrus.Infof("[server] Shutdown request (/shutdown %v)", sig)
	}

	logrus.Infoln("[server] Stoping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		logrus.Infoln("[server] Shutdown request error: ", err)
	}
}

// @Summary Shutdown Server
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /private/shutdown [post]
func (s *Server) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

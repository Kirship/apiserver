package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Starting APIServer...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.loglevel)

	s.logger.SetLevel(level)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handlehello())
}

func (s *APIServer) handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

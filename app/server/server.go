package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/filter", s.handleFilter()).Methods("POST")
	s.router.HandleFunc("/compare", s.handleCompare()).Methods("POST")
}

func (s *Server) handleFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.Gps{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("(/filter) incorrect body")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var err error
		req.Lat, req.Long, err = service.KalmanFilter(req.Lat, req.Long)
		if err != nil {
			s.logger.Error("(/filter) filter error")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.logger.Info("(/filter) success data filter")
		s.respond(w, r, http.StatusOK, req)
	}
}

func (s *Server) handleCompare() http.HandlerFunc {
	type trajectory struct {
		Gps           model.Gps           `json:"gps"`
		Accelerometer model.Acceletometer `json:"acc"`
	}

	type request struct {
		Perfect  trajectory `json:"perfect"`
		Compared trajectory `json:"compared"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("(/compare) incorrect body")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

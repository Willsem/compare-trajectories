// Package server ...
package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/server/config"
	"github.com/Willsem/compare-trajectories/app/service/comparing"
	"github.com/Willsem/compare-trajectories/app/service/filtering"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *config.Config) *Server {
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

	s.logger.Info("starting server at", s.config.BindAddr)

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
		req.Lat, req.Long, err = filtering.KalmanFilter(req.Lat, req.Long)
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
	type request struct {
		Perfect  model.Trajectory `json:"perfect"`
		Compared model.Trajectory `json:"compared"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("(/compare) incorrect body")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		comparedTrajectory, err := comparing.Compare(req.Perfect, req.Compared)
		if err != nil {
			s.logger.Error("(/compare) error with comparing")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.logger.Info("(/compare) success comparing")
		s.respond(w, r, http.StatusOK, comparedTrajectory)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("status", fmt.Sprint(code))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) makeRoutes() {
	s.mux = mux.NewRouter().StrictSlash(true)

	authMW := AuthMW(s.log, s.cfg, s.uc.Auth)
	s.mux.Use(authMW)

	v1 := s.mux.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/login", s.Login).Methods(http.MethodPost)
	v1.HandleFunc("/logout", s.Logout).Methods(http.MethodPost)
	v1.HandleFunc("/register", s.Register).Methods(http.MethodPost)
	v1.HandleFunc("/me", s.CurrentUser).Methods(http.MethodGet)
}

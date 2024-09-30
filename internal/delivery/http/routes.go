package http

import (
	"github.com/gorilla/mux"
)

func (s *Server) makeRoutes() {
	s.mux = mux.NewRouter().StrictSlash(true)

	authMW := AuthMW(s.log, s.cfg, s.uc.Auth)
	s.mux.Use(authMW)

	v1 := s.mux.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/login", s.Login)
	v1.HandleFunc("/logout", s.Logout)
	v1.HandleFunc("/register", s.Register)
	v1.HandleFunc("/me", s.CurrentUser)
}

package common

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/delivery/http/middleware"
)

func (s *Server) makeRoutes() {
	s.mux = mux.NewRouter().StrictSlash(true)

	hV1 := s.handlers.V1

	corsMW := middleware.CorsMW(s.cfg)
	s.mux.Use(corsMW)

	authMW := middleware.AuthMW(s.log, s.cfg, hV1.UC.Auth)
	s.mux.Use(authMW)

	v1 := s.mux.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/login", hV1.Login).Methods(http.MethodPost)
	v1.HandleFunc("/logout", hV1.Logout).Methods(http.MethodPost)
	v1.HandleFunc("/register", hV1.Register).Methods(http.MethodPost)
	v1.HandleFunc("/me", hV1.CurrentUser).Methods(http.MethodGet)

	v1.HandleFunc("/feed", hV1.Feed).Methods(http.MethodGet)

	v1.HandleFunc("/users/{authorID:[0-9]+}/articles", hV1.UserArticles).Methods(http.MethodGet)

	v1.HandleFunc("/users/{userID:[0-9]+}", hV1.GetUserInfo).Methods(http.MethodGet)
	v1.HandleFunc("/users/{userID:[0-9]+}", hV1.UpdateUserInfo).Methods(http.MethodPut)

	v1.HandleFunc("/users/{userID:[0-9]+}/changepassword", hV1.UpdatePassword).Methods(http.MethodPut)
}

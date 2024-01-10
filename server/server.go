package server

import (
	"log"
	"net/http"
	"web-college/database"
	"web-college/server/handlers"
	"web-college/middleware"
	"github.com/gorilla/mux"
)

type srv struct {
	Server  *http.Server
	Storage *database.Storage
}

func New() *srv {
	
	return &srv{
	}
}

func (s *srv) App() {
	db, err := database.DBLite()
	if err != nil {
		log.Fatal(err)
	}
	s.Storage = db

	router := mux.NewRouter()
	dir := http.Dir("./static")
	fs := http.StripPrefix("/static/", http.FileServer(dir))
	router.PathPrefix("/static/").Handler(fs) //static load
	router.HandleFunc("/login", handlers.RegHandler(s.Storage)).Methods("GET", "POST")
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.Use(middleware.PanicRecovery)

	log.Fatal(http.ListenAndServe(":80", router))

}

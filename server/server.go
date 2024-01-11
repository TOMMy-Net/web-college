package server

import (
	"log"
	"net/http"
	"web-college/database"
	"web-college/logging"
	"web-college/middleware"
	"web-college/server/handlers"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type srv struct {
	Server  *http.Server
	Storage *database.Storage
	Logger  *logrus.Logger
}

func New() *srv {
	log := logging.LoggerNew(logging.ServerError)
	db, err := database.DBLite()
	if err != nil {
		log.Fatal(err)
	}
	return &srv{Storage: db, Logger: log}
}

func (s *srv) App() {

	router := mux.NewRouter()
	router.Use(middleware.PanicRecovery(s.Logger)) //middleware

	dir := http.Dir("./static")
	fs := http.StripPrefix("/static/", http.FileServer(dir))
	router.PathPrefix("/static/").Handler(fs) //static load
	router.HandleFunc("/login", handlers.RegHandler(s.Storage)).Methods("GET", "POST")
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	log.Println("Server start")
	s.Logger.Fatal(http.ListenAndServe(":80", router))

}

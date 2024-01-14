package server

import (
	"log"
	"net/http"
	"os"
	"web-college/database"
	"web-college/logging"
	"web-college/middleware"
	"web-college/server/handlers"

	gorillaH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type srv struct {
	Server       *http.Server
	Storage      *database.Storage
	LoggerDB     *logrus.Logger
	LoggerServer *logrus.Logger
}

func New() *srv {
	logServerErr := logging.LoggerNew()
	logServerErr.SetOutput(logging.WriteLog(logging.ServerError))

	db, err := database.DBLite()
	if err != nil {
		log.Fatal(err)
	}
	return &srv{Storage: db, LoggerServer: logServerErr}
}

func (s *srv) App() {

	router := mux.NewRouter()
	router.Use(middleware.PanicRecovery(s.LoggerServer)) //middleware

	dir := http.Dir("./static")
	fs := http.StripPrefix("/static/", http.FileServer(dir))
	router.PathPrefix("/static/").Handler(fs) //static load
	router.HandleFunc("/signup", handlers.RegHandler(s.Storage)).Methods("GET", "POST")
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	log.Println("Server start")
	s.LoggerServer.Fatal(http.ListenAndServe(":80", gorillaH.LoggingHandler(os.Stdout, router)))
	defer log.Println("Server stop")

}

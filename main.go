package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type RegForm struct {
	FirsName    string
	LastName    string
	PhoneNumber string
}

func RegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		f := RegForm{
			FirsName:    r.FormValue("first_name"),
			LastName:    r.FormValue("last_name"),
			PhoneNumber: r.FormValue("phone_number"),
		}
		fmt.Print(f)
		val := reflect.ValueOf(f)
		if f == (RegForm{}) {
			http.ServeFile(w, r, "static/html/reg.html")
			return
		}
		for i := 0; i < val.NumField(); i++ {
			if (val.Field(i).Interface()) == "" {
				http.ServeFile(w, r, "static/html/reg.html")
				return
			}

		}

	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/html/reg.html")
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RegHandler).Methods("GET", "POST")
	dir := http.Dir("./static")
	fs := http.StripPrefix("/static/", http.FileServer(dir))
	router.PathPrefix("/static/").Handler(fs)


	log.Fatal(http.ListenAndServe(":80", router))

}

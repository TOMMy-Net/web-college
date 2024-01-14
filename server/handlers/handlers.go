package handlers

import (
	"net/http"
	"reflect"
	"web-college/internal"

	//"github.com/sirupsen/logrus"
	"html/template"
)

type RegForm struct {
	FirsName string
	LastName string
	Email    string
	Password string
}

type IsReg struct {
	Yes bool
}

type SaveUser interface {
	SaveUser(fn, ln, email, password string) error
	CheckUser(email string) (bool, error)
}

func RegHandler(db SaveUser) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			f := RegForm{
				FirsName: r.FormValue("firstName"),
				LastName: r.FormValue("lastName"),
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
			}

			val := reflect.ValueOf(f)
			if f == (RegForm{}) {
				http.ServeFile(w, r, "templates/html/reg.html")
				return
			}
			for i := 0; i < val.NumField(); i++ {
				if (val.Field(i).Interface()) == "" {
					http.ServeFile(w, r, "templates/html/reg.html")
					return
				}

			}
			if ok, err := db.CheckUser(f.Email); err != nil {
				panic(err)
			} else if ok {
				tmpl, _ := template.ParseFiles("templates/html/reg.html")
				tmpl.Execute(w, IsReg{Yes: true})
				return
			}

			err := db.SaveUser(f.FirsName, f.LastName, f.Email, internal.SumPassword(f.Password))
			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		} else if r.Method == http.MethodGet {
			tmpl, _ := template.ParseFiles("templates/html/reg.html")
			tmpl.Execute(w, "")
		}
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/html/index.html")
}

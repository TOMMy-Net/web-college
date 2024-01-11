package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func PanicRecovery(log *logrus.Logger)  func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					log.WithFields(logrus.Fields{
						"URL": req.URL,
						"METHOD": req.Method,
					}).Error(r)
					return
				}
			}()
			next.ServeHTTP(w, req)
		})
	}
}

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
	})
}

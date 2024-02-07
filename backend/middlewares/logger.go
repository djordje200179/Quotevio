package middlewares

import (
	"log"
	"net/http"
)

type logger struct {
	*log.Logger

	next http.Handler
}

func Log(next http.Handler, l *log.Logger) http.Handler {
	return logger{next: next, Logger: l}
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.Println(r.Method, r.URL.Path, r.RemoteAddr)

	l.next.ServeHTTP(w, r)
}

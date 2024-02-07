package middlewares

import (
	"log"
	"net/http"
	"os"
	"slices"
)

type logger struct {
	stdOut, stdErr *log.Logger

	next http.Handler
}

type responseRecorder struct {
	http.ResponseWriter

	status int
	body   []byte
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = slices.Clone(b)
	return r.ResponseWriter.Write(b)
}

func Log(next http.Handler) http.Handler {
	return logger{
		stdOut: log.New(os.Stdout, "", 0),
		stdErr: log.New(os.Stderr, "", 0),
		next:   next,
	}
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.stdOut.Println(r.Method, r.URL.Path, r.RemoteAddr)

	response := &responseRecorder{ResponseWriter: w}
	l.next.ServeHTTP(response, r)

	switch response.status % 100 {
	case 4:
		l.stdOut.Println(response.status, string(response.body))
	case 5:
		l.stdErr.Println(response.status, string(response.body))
	}
}

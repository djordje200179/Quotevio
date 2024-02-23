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

type responseInterceptor struct {
	http.ResponseWriter

	status int
	body   []byte
}

func (r *responseInterceptor) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *responseInterceptor) Write(b []byte) (int, error) {
	r.body = slices.Clone(b)
	return r.ResponseWriter.Write(b)
}

func Log(next http.Handler) http.Handler {
	return logger{
		stdOut: log.New(os.Stdout, "", log.LstdFlags),
		stdErr: log.New(os.Stderr, "", log.LstdFlags),

		next: next,
	}
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.stdOut.Println(r.Method, r.URL.Path, r.RemoteAddr)

	response := &responseInterceptor{ResponseWriter: w}
	l.next.ServeHTTP(response, r)

	statusGroup := response.status / 100
	switch statusGroup {
	case 2, 3:
		return
	case 4:
		l.stdOut.Printf("(%d) %s\n", response.status, response.body)
	case 5:
		l.stdErr.Printf("(%d) %s\n", response.status, response.body)
	}
}

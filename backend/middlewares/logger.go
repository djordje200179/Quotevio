package middlewares

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
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

	statusGroup := response.status / 100
	switch statusGroup {
	case 2, 3:
		return
	case 4:
		l.stdOut.Println(makeResponseMessage(response.status, response.body))
	case 5:
		l.stdErr.Println(makeResponseMessage(response.status, response.body))
	}
}

func makeResponseMessage(statusCode int, body []byte) string {
	var sb strings.Builder

	sb.WriteByte('(')
	sb.WriteString(strconv.Itoa(statusCode))
	sb.WriteByte(')')
	sb.WriteByte(' ')
	sb.Write(body)

	return sb.String()
}

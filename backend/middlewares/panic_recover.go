package middlewares

import (
	"log"
	"net/http"
)

type panicRecover struct {
	next http.Handler
}

func PanicRecover(next http.Handler) http.Handler {
	return panicRecover{
		next: next,
	}
}

func (p panicRecover) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}()

	p.next.ServeHTTP(w, req)
}

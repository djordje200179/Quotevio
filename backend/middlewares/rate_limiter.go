package middlewares

import (
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"time"
)

const RateLimit = 10

type rateLimiter struct {
	db *redis.Client

	next http.Handler
}

func Limit(next http.Handler) http.Handler {
	db := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	return &rateLimiter{
		db: db,

		next: next,
	}
}

func (r *rateLimiter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ipAddr := req.RemoteAddr

	ctx := req.Context()

	err := r.db.SetNX(ctx, ipAddr, RateLimit, time.Minute).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, err := r.db.Decr(ctx, ipAddr).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if val < 0 {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	r.next.ServeHTTP(w, req)
}

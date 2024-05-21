package main

import (
	"log"
	"net/http"

	config "github.com/deduardolima/ratelimiter/configs"
	"github.com/deduardolima/ratelimiter/infra/database"
	"github.com/deduardolima/ratelimiter/internal/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Error loading config file, %s", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	store := database.NewRedisStore(rdb)
	r := mux.NewRouter()
	r.Use(middleware.RateLimiterMiddleware(store, cfg))

	r.HandleFunc("/", HomeHandler)

	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

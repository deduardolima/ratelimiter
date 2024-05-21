package internal_test

import (
	"log"
	"testing"
	"time"

	config "github.com/deduardolima/ratelimiter/configs"
	"github.com/deduardolima/ratelimiter/infra/database"
	"github.com/deduardolima/ratelimiter/internal/limiter"
	"github.com/go-redis/redis/v8"
)

func TestRateLimiter(t *testing.T) {
	cfg, err := config.LoadConfig("/app/.env")
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	store := database.NewRedisStore(client)
	store.FlushDB()

	rateLimiter := limiter.NewRateLimiter(store, cfg)

	// Testando limitação por IP
	ip := "192.168.1.1"
	for i := 0; i < cfg.RateLimitIP; i++ {
		if !rateLimiter.AllowRequest(ip, "") {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}
	if rateLimiter.AllowRequest(ip, "") {
		t.Errorf("6th request should be blocked")
	}

	time.Sleep(2 * time.Second)

	store.FlushDB()

	// Testando limitação por Token
	token := "abc123"
	for i := 0; i < cfg.RateLimitToken; i++ {
		if !rateLimiter.AllowRequest("", token) {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}
	if rateLimiter.AllowRequest("", token) {
		t.Errorf("11th request should be blocked")
	}
}

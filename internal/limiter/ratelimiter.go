package limiter

import (
	"log"
	"time"

	config "github.com/deduardolima/ratelimiter/configs"
	"github.com/deduardolima/ratelimiter/infra/database"
	"golang.org/x/net/context"
)

type RateLimiter struct {
	Store     database.Store
	IpRate    int
	TokenRate int
	BlockTime time.Duration
	Ctx       context.Context
}

func NewRateLimiter(store database.Store, cfg *config.Config) *RateLimiter {

	return &RateLimiter{
		Store:     store,
		IpRate:    cfg.RateLimitIP,
		TokenRate: cfg.RateLimitToken,
		BlockTime: time.Duration(cfg.BlockTime) * time.Second,
		Ctx:       context.Background(),
	}
}

func (rl *RateLimiter) AllowRequest(ip string, token string) bool {
	if token != "" {
		allowed := rl.allow(token, rl.TokenRate)
		if !allowed {
			log.Printf("Token %s bloqueado\n", token)
			return false
		}
	}

	if ip != "" {
		allowed := rl.allow(ip, rl.IpRate)
		if !allowed {
			log.Printf("IP %s bloqueado\n", ip)
			return false
		}
	}

	return true
}

func (rl *RateLimiter) IsBlocked(key string) (bool, error) {
	blocked, err := rl.Store.Get(key + "_blocked")
	if err != nil {
		return false, err
	}
	return blocked == "1", nil
}

func (rl *RateLimiter) allow(key string, limit int) bool {
	blocked, err := rl.Store.Get(key + "_blocked")
	if err == nil && blocked == "1" {
		return false
	}

	count, err := rl.Store.Increment(key)
	if err != nil {
		return false
	}

	if count == 1 {
		rl.Store.Expire(key, time.Second)
	}

	if count > int64(limit) {
		log.Printf("Chave %s excedeu o limite de %d. Bloqueando por %s\n", key, limit, rl.BlockTime)
		rl.Store.Set(key+"_blocked", "1", rl.BlockTime)
		return false
	}

	return true
}

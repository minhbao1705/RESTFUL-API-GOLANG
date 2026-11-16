package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exits := clients[ip]
	if !exits {
		limiter := rate.NewLimiter(5, 10) // 5 request/sec, brust 10
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient
		// log.Printf("A client[%s] - {limiter: %+v, lastSeen: %s}", ip, newClient.limiter, newClient.lastSeen)
		return limiter
	}

	// log.Printf("A client[%s] - {limiter: %+v, lastSeen: %s}", ip, client.limiter, client.lastSeen)
	client.lastSeen = time.Now()
	return client.limiter
}

func CLeanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// ab -n 2000 -c 1 -H "X-API-KEY:419c7a3a-3642-4678-b9b5-0a6572a191cb" localhost:8080/api/v1/categories/python

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		// log.Println("IP Address: ", ip) // ::1 <=> 127.0.0.1

		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many reques",
				"message": "Bạn đã gửi quá nhiều request. Hãy thử lại sau",
			})
			return
		}
		ctx.Next()
	}
}

package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/gddo/log"
	"github.com/joho/godotenv"
)

type Redis struct {
	Rds *redis.Client
}

func (r *Redis) ConnectRedis() *redis.Client {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Error(ctx, "Error loading .env file.", nil)
	}

	// Get declared name in env
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisDbStr := os.Getenv("REDIS_DB")
	redisPoolSizeStr := os.Getenv("REDIS_POOL_SIZE")
	redisIdleTimeoutStr := os.Getenv("REDIS_IDLE_TIMEOUT_MINUTES")

	// Konversi string menjadi tipe yang sesuai
	redisDB, _ := strconv.Atoi(redisDbStr)
	redisPoolSize, _ := strconv.Atoi(redisPoolSizeStr)
	redisIdleTimeout, _ := strconv.Atoi(redisIdleTimeoutStr)

	// Inisialisasi koneksi ke Redis dengan menggunakan connection pool
	client := redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    redisPass,
		DB:          redisDB,
		PoolSize:    redisPoolSize,
		IdleTimeout: time.Duration(redisIdleTimeout) * time.Minute,
	})

	// Set koneksi Redis pada struct
	r.Rds = client
	fmt.Println("Success connect to Redis.")

	return client
}

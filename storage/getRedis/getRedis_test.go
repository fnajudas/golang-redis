package getredis

import (
	"context"
	dtredis "golangredis/models/dtRedis"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestGetDataRedis(t *testing.T) {
	ctx := context.Background()
	miniRedis, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}
	defer miniRedis.Close()

	options := &redis.Options{
		Addr: miniRedis.Addr(),
	}
	client := redis.NewClient(options)

	// Set mock redis server
	key := "myKey"
	value := "Value"
	miniRedis.Set(key, value)

	request := dtredis.DataSet{
		Key: key,
	}

	// Initialize get data
	g := NewGetData(ctx, client)

	// Panggil metode GetDataRedis
	resp, err := g.GetDataRedis(request)
	if err != nil {
		t.Errorf("Error calling GetDataRedis: %v", err)
	}

	// Check response
	assert.Equal(t, value, resp.Key)
}

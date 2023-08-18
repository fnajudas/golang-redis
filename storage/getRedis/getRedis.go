package getredis

import (
	"context"
	dtredis "golangredis/models/dtRedis"

	"github.com/go-redis/redis"
)

type getData struct {
	ctx    context.Context
	client *redis.Client
}

func NewGetData(ctx context.Context, client *redis.Client) *getData {
	return &getData{
		ctx:    ctx,
		client: client,
	}
}

func (g *getData) GetDataRedis(req dtredis.DataSet) (resp dtredis.RespGetData, err error) {
	value, err := g.client.Get(req.Key).Result()
	if err != nil {
		return resp, err
	}

	resp = dtredis.RespGetData{
		Key: value,
	}

	return resp, nil
}

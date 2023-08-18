package setredis

import (
	"context"

	dtredis "golangredis/models/dtRedis"

	"github.com/go-redis/redis"
)

type setRedis struct {
	ctx    context.Context
	client *redis.Client
}

func NewSetRedis(ctx context.Context, client *redis.Client) *setRedis {
	return &setRedis{
		ctx:    ctx,
		client: client,
	}
}

func (s *setRedis) SetData(request dtredis.DataSet) error {
	err := s.client.Set(request.Key, request.Value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

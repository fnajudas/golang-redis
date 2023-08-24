package setredis

import (
	"context"
	"errors"
	"log"

	dtredis "golangredis/models/dtRedis"

	"github.com/go-redis/redis"
)

type saveToDatabase interface {
	SaveMysql(req dtredis.DataSet) error
}

type setRedis struct {
	ctx            context.Context
	client         *redis.Client
	saveToDatabase saveToDatabase
}

func NewSetRedis(ctx context.Context, client *redis.Client, saveToDatabase saveToDatabase) *setRedis {
	return &setRedis{
		ctx:            ctx,
		client:         client,
		saveToDatabase: saveToDatabase,
	}
}

func (s *setRedis) SetData(request dtredis.DataSet) error {
	err := s.client.Set(request.Key, request.Value, 0).Err()
	if err != nil {
		return err
	}

	saveDb := s.saveToDatabase.SaveMysql(request)
	if saveDb != nil {
		log.Println("Error to save database")
		return errors.New("Error to save database")
	}

	return nil
}

package main

import (
	"context"
	"golangredis/config/mysql"
	"golangredis/config/redis"
	"golangredis/routes"
	"os"

	setDatabase "golangredis/storage/setRedis"

	setDataHandler "golangredis/controller/setRedis"
	setDataStorage "golangredis/storage/setRedis"

	getDataHandler "golangredis/controller/getRedis"
	getDataStorage "golangredis/storage/getRedis"

	"github.com/thedevsaddam/renderer"
)

// test push

func main() {
	ctx := context.Background()

	// Connect MySQL
	var Database = mysql.Db{}
	Database.DatabaseConnection()
	defer Database.Database.Close()
	db := Database.Database

	// Membuat instance Redis
	var redisInstance = redis.Redis{}
	redisInstance.ConnectRedis()
	defer redisInstance.Rds.Close()

	// Menggunakan client Redis dari struct Redis
	redisClient := redisInstance.Rds
	render := renderer.New()

	getDataStorage := getDataStorage.NewGetData(ctx, redisClient)
	getDataHandler := getDataHandler.NewHandler(getDataStorage, render)

	setDatabase := setDatabase.NewSetMysql(db)
	setDataStorage := setDataStorage.NewSetRedis(ctx, redisClient, setDatabase)
	setDataHandler := setDataHandler.NewSetRedis(setDataStorage, getDataStorage)

	r := routes.Routes{
		Redis:   redisClient,
		SetData: setDataHandler,
		GetData: getDataHandler,
	}

	r.Run(os.Getenv("APP_PORT"))
}

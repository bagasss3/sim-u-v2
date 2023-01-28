package database

import (
	"context"
	"sim-u/config"

	log "github.com/sirupsen/logrus"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	conn, err := clientRedis(config.RedisHost(), config.RedisDB(), config.RedisPoolSize(), config.RedisMaxIdleConns())
	if err != nil {
		log.WithField("redisHost", config.RedisHost()).Fatal("Failed to connect redis:", err)
	}
	log.Info("Success connect redis")
	return conn
}

func clientRedis(host string, db, poolSize, MaxIdleConns int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     "", // no password set
		DB:           db, // use default DB
		PoolSize:     poolSize,
		MaxIdleConns: MaxIdleConns,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}

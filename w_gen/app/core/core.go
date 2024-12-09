package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/m-wilk/w_gen/repository"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/jackc/pgx/v5/stdlib"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Core struct {
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
	Repository  repository.Repository
	RedisClient *redis.Client
}

func New() *Core {
	return &Core{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *Core) initDB(db string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db))

	if err != nil {
		c.ErrorLog.Fatalf("initDB: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		c.ErrorLog.Fatalf("initDB: %v", err)
		return nil, err
	}

	return client, nil
}

func (c *Core) InitRepository(databaseName string) error {
	client, err := c.initDB(
		os.Getenv("SUS_DB"),
	)

	if err != nil {
		return err
	}

	c.Repository = *repository.New(client, databaseName)

	return nil
}

func (c *Core) InitRedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		c.ErrorLog.Fatalf("redis connection problem %v", err)
	}

	c.RedisClient = client
}

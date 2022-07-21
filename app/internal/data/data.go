package data

import (
	"context"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq" // db driver
	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent/migrate"
	"github.com/zYros90/pkg/logger"
)

const dbConnTimeout = 10 * time.Second

type Data struct {
	ent   *ent.Client
	redis *redis.Client
}

func New(logger *logger.Log, conf *config.Config) (*Data, error) {
	entClient, err := newEntClient(logger, conf)
	if err != nil {
		return nil, errors.Wrap(err, "error creating ent client")
	}
	redisClient, err := newRedisClient(logger, conf)
	if err != nil {
		return nil, errors.Wrap(err, "error creating redis client")
	}
	return &Data{
		ent:   entClient,
		redis: redisClient,
	}, nil
}

// New initializes ent client.
func newEntClient(logger *logger.Log, conf *config.Config) (*ent.Client, error) {
	source := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		conf.PG.Host,
		conf.PG.Port,
		conf.PG.Auth.User,
		conf.PG.DBName,
		conf.PG.Auth.Pass,
		conf.PG.SSL,
	)

	db, err := entsql.Open("postgres", source)
	if err != nil {
		return nil, errors.Wrap(err, "error open connection to db")
	}

	opts := []ent.Option{
		ent.Log(logger.Sugar().Debug),
		ent.Driver(db),
	}

	// Open the database connection.
	client := ent.NewClient(opts...)

	ctx, cancel := context.WithTimeout(context.Background(), dbConnTimeout)
	defer cancel()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func newRedisClient(logger *logger.Log, conf *config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Auth.Pass, // no password set
		DB:       conf.Redis.DB,        // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func ExampleClient() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

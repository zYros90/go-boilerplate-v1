package data

import (
	"context"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq" // db driver
	"github.com/pkg/errors"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data/ent/migrate"
	"github.com/zYros90/pkg/logger"
)

const dbConnTimeout = 10 * time.Second

// New initializes ent client.
func New(logger *logger.Log, conf *config.Config) (*ent.Client, error) {
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

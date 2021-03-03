package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/database"
	"moocss.com/gaea/pkg/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewArticleRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)
	config := conf.File("config")

	// database
	dbClient, err := database.Init(logger)
	if err != nil {
		dialect := config.Get("database.dialect")
		log.Errorf("failed opening connection to %s: %v", dialect, err)
		return nil, err
	}

	// redis cache
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Get("cache.redis.addr"),
		DB:       config.GetInt("cache.redis.addr"),
		Password: config.Get("cache.redis.password"),
	})

	return &Data{
		db:  dbClient,
		rdb: rdb,
	}, nil
}

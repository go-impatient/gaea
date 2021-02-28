package data

import (
	"github.com/google/wire"

	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/database"
	"moocss.com/gaea/pkg/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewPostRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)
	client, err := database.Init(logger)
	if err != nil {
		dialect := conf.Get("database.dialect")
		log.Errorf("failed opening connection to %s: %v", dialect, err)
		return nil, err
	}

	return &Data{db: client}, err
}

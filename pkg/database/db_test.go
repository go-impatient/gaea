package database

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/log"
)

var dbClient *ent.Client
var err error
var logger = log.DefaultLogger

func TestMain(m *testing.M) {
	dbClient, err = Init(logger)

	code := m.Run()
	if dbClient != nil {
		_ = dbClient.Close() // cleanup
	}

	os.Exit(code)
}

func TestGetConnection(t *testing.T) {
	log := log.NewHelper("conf", logger)
	conf.Load("../../config", log)

	dbClient, err = Init(logger)
	assert.Nil(t, err)
	assert.NotNil(t, dbClient)
	assert.NotNil(t, GetDatabase())
	assert.Equal(t, GetDatabase(), dbClient)
	assert.Nil(t, dbClient.Close())
}

func TestDatabaseClient(t *testing.T) {
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}

	total, err := dbClient.User.
		Query().
		Count(context.TODO())

	if err != nil {
		t.Fatalf("Database Query failed, %v!", err)
	}
	if total != 0 {
		t.Fatalf("Database Count failed, expected: %d, actual: %d", 0, total)
	}

	t.Logf("Total Users: %d", total)
}

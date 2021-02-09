package database

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"moocss.com/gaea/internal/data/ent"
)

var dbClient *ent.Client
var err error

func TestMain(m *testing.M) {
	dbClient, err = InitDatabase()

	code := m.Run()
	if dbClient != nil {
		_ = dbClient.Close() // cleanup
	}

	os.Exit(code)
}

func TestGetConnection(t *testing.T) {
	dbClient, err = InitDatabase()
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

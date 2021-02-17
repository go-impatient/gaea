package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	Load("../../config")

	logPath := Get("LOG_PATH")
	assert.NotNil(t, logPath)
	//t.Logf("LOG_PATH: %s", logPath)
	logLevel := File("gaea").Get("LOG_LEVEL")
	//t.Logf("LOG_LEVEL: %s", logLevel)
	assert.NotNil(t, logLevel)
	dsn := File("config").Get("database.dsn")
	assert.NotNil(t, dsn)
	//t.Logf("DSN: %s", dsn)
}

package config_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/nguyenhoaibao/logsvc/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromEnv(t *testing.T) {
	const (
		serverAddr = ":8080"
		dbAddr     = "127.0.0.1:2345"
		dbUser     = "postgres"
		dbPass     = "mypostgrespw"
		dbName     = "test"
	)

	os.Setenv("LOGSVC_SERVER_ADDR", serverAddr)
	os.Setenv("LOGSVC_DATABASE_ADDR", dbAddr)
	os.Setenv("LOGSVC_DATABASE_USER", dbUser)
	os.Setenv("LOGSVC_DATABASE_PASSWORD", dbPass)
	os.Setenv("LOGSVC_DATABASE_NAME", dbName)
	defer func() {
		os.Unsetenv("LOGSVC_SERVER_ADDR")
		os.Unsetenv("LOGSVC_DATABASE_ADDR")
		os.Unsetenv("LOGSVC_DATABASE_USER")
		os.Unsetenv("LOGSVC_DATABASE_PASSWORD")
		os.Unsetenv("LOGSVC_DATABASE_NAME")
	}()

	var b []byte
	c, err := config.Load(bytes.NewReader(b))

	assert.NoError(t, err)
	assert.Equal(t, serverAddr, c.Server.Addr)
	assert.Equal(t, dbAddr, c.Database.Addr)
	assert.Equal(t, dbUser, c.Database.User)
	assert.Equal(t, dbPass, c.Database.Password)
	assert.Equal(t, dbName, c.Database.Name)
}

func TestLoadConfigFromReader(t *testing.T) {
	b := []byte(`
server:
  addr: ":8080"
database:
  addr: "127.0.0.1:5432"
  user: "postgres"
  password: "mypostgrespw"
  name: "test"
`)
	r := bytes.NewReader(b)
	c, err := config.Load(r)

	assert.NoError(t, err)
	assert.Equal(t, ":8080", c.Server.Addr)
	assert.Equal(t, "127.0.0.1:5432", c.Database.Addr)
	assert.Equal(t, "postgres", c.Database.User)
	assert.Equal(t, "mypostgrespw", c.Database.Password)
	assert.Equal(t, "test", c.Database.Name)
}

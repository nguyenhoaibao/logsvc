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

	f := []byte{}
	c, err := config.Load(bytes.NewReader(f))

	assert.NoError(t, err)
	assert.Equal(t, &config.Config{
		Server: struct {
			Addr string
		}{
			Addr: serverAddr,
		},
		Database: struct {
			Addr     string
			User     string
			Password string
			Name     string
		}{
			Addr:     dbAddr,
			User:     dbUser,
			Password: dbPass,
			Name:     dbName,
		},
	}, c)
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
	assert.Equal(t, &config.Config{
		Server: struct {
			Addr string
		}{
			Addr: ":8080",
		},
		Database: struct {
			Addr     string
			User     string
			Password string
			Name     string
		}{
			Addr:     "127.0.0.1:5432",
			User:     "postgres",
			Password: "mypostgrespw",
			Name:     "test",
		},
	}, c)
}

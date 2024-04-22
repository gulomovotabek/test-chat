package repository

import (
	"cae_chat/config"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type connectionConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Dbname   string `env:"POSTGRES_DB"`
	Port     int    `env:"POSTGRES_PORT"`
	SslMode  bool   `env:"POSTGRES_SSL_MODE"`
}

func (c *connectionConfig) ssl() string {
	if !c.SslMode {
		return "sslmode=disable"
	}
	return ""
}

func (c *connectionConfig) build() string {
	return fmt.Sprintf(
		`host=%s
     port=%d 
     user=%s 
     password=%s 
     dbname=%s 
     %s`,
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Dbname,
		c.ssl(),
	)
}

func New() (*sqlx.DB, error) {
	configInstance := config.Load(&connectionConfig{})
	return sqlx.Connect(
		"postgres",
		configInstance.build())
}

package config

import "fmt"

type DBConfig struct {
	PostgresDatabase string `env:"POSTGRES_DATABASE,required"`
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	PostgresPort     uint16 `env:"POSTGRES_PORT,required"`
	PostgresHost     string `env:"POSTGRES_HOST,required"`
}

func (c *DBConfig) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDatabase,
	)
}

package config

import (
	"github.com/spf13/pflag"
	"strconv"
)

const (
	dbConfigFile = "config/db.env"
)

type DbConfig struct {
	Dialect            string
	Debug              bool
	MaxOpenConnections int
	DbName             string
	User               string
	Password           string
	Host               string
	Port               int
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Dialect:            "postgres",
		Debug:              false,
		MaxOpenConnections: 50,
		DbName:             "postres",
		User:               "postgres",
		Password:           "",
		Host:               "",
		Port:               5432,
	}
}

func (c *DbConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.DbName, "db-name", c.DbName, "Name of the Database")
	fs.StringVar(&c.User, "db-user", c.User, "Username of the user to access the DB")
	fs.StringVar(&c.Password, "db-password", c.Password, "Password to access the DB")
	fs.StringVar(&c.Host, "db-host", c.Host, "Host at which database is being serviced")
	fs.IntVar(&c.Port, "db-port", c.Port, "Port at which the database host is exposed")
}

func (c *DbConfig) ReadFromFile() error {
	return nil
}

func (c *DbConfig) ReadFromEnv() error {
	c.DbName = getEnvDefault("DB_NAME", c.DbName)
	c.User = getEnvDefault("DB_USER", c.User)
	c.Password = getEnvDefault("DB_PASSWORD", c.Password)
	c.Host = getEnvDefault("DB_HOST", c.Host)

	portStr := getEnvDefault("DB_PORT", strconv.Itoa(c.Port))
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}
	c.Port = portInt
	return nil
}

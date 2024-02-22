package config

import (
	"github.com/spf13/pflag"
	"strconv"
)

const (
	serverConfigFile = "config/server.env"
)

type ServerConfig struct {
	JwtSecret     string
	ListenPort    int
	ListenAddress string
	EnvName       string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		JwtSecret:     "",
		ListenPort:    8000,
		ListenAddress: "0.0.0.0",
		EnvName:       "dev",
	}
}

func (c *ServerConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.JwtSecret, "server-jwt-secret", c.JwtSecret, "Jwt Secret")
	fs.StringVar(&c.ListenAddress, "server-listen-address", c.ListenAddress, "The IP address of the network interface at which the server should listen and be exposed")
	fs.StringVar(&c.EnvName, "server-env-name", c.EnvName, "Name of the environment in which the server is running: dev/stage/prod")
	fs.IntVar(&c.ListenPort, "server-listen-port", c.ListenPort, "Port at which the server should run")
}

func (c *ServerConfig) ReadFromFile() error {
	return nil
}

func (c *ServerConfig) ReadFromEnv() error {
	c.JwtSecret = getEnvDefault("SERVER_JWT_SECRET", c.JwtSecret)
	c.ListenAddress = getEnvDefault("SERVER_LISTEN_ADDRESS", c.ListenAddress)
	c.EnvName = getEnvDefault("SERVER_ENV_NAME", c.EnvName)

	portStr := getEnvDefault("SERVER_LISTEN_PORT", strconv.Itoa(c.ListenPort))
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}
	c.ListenPort = portInt
	return nil
}

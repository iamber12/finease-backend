package config

import (
	"flag"
	"github.com/spf13/pflag"
	"sync"
)

type Config interface {
	ReadFromEnv() error
	ReadFromFile() error
	AddFlags(set *pflag.FlagSet)
}

type ApplicationConfig struct {
	ServerConfig *ServerConfig
	DbConfig     *DbConfig
	AwsConfig    *AwsConfig
}

var Conf = ApplicationConfig{}
var initOnce = sync.Once{}

func Setup(flagset *pflag.FlagSet) (err error) {
	initOnce.Do(func() {
		Conf.ServerConfig = NewServerConfig()
		Conf.DbConfig = NewDbConfig()
		Conf.AwsConfig = NewAwsConfig()

		flagset.AddGoFlagSet(flag.CommandLine)

		// Server config setup
		if err = Conf.ServerConfig.ReadFromEnv(); err != nil {
			return
		}

		if err = Conf.ServerConfig.ReadFromFile(); err != nil {
			return
		}
		Conf.ServerConfig.AddFlags(flagset)

		// DB Config setup
		if err = Conf.DbConfig.ReadFromEnv(); err != nil {
			return
		}
		if err = Conf.DbConfig.ReadFromFile(); err != nil {
			return
		}
		Conf.DbConfig.AddFlags(flagset)

		// AWS Config Setup
		if err = Conf.AwsConfig.ReadFromEnv(); err != nil {
			return
		}
		if err = Conf.AwsConfig.ReadFromFile(); err != nil {
			return
		}
		Conf.AwsConfig.AddFlags(flagset)
	})

	return err
}

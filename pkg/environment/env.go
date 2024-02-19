package environment

import (
	"flag"
	"github.com/spf13/pflag"
	"sync"
)

type environment struct {
	ServerConfig *serverConfig
	DbConfig     *dbConfig
	AwsConfig    *awsConfig
}

var env = environment{}
var initOnce = sync.Once{}

func Environment() environment {
	return env
}

func (e environment) Setup(flagset *pflag.FlagSet) error {
	var err error
	initOnce.Do(func() {
		env.ServerConfig = NewServerConfig()
		env.DbConfig = NewDbConfig()
		env.AwsConfig = NewAwsConfig()

		flagset.AddGoFlagSet(flag.CommandLine)

		// Server config setup
		if err = env.ServerConfig.ReadFromEnv(); err != nil {
			return
		}

		if err = env.ServerConfig.ReadFromFile(); err != nil {
			return
		}
		env.ServerConfig.AddFlags(flagset)

		// DB Config setup
		if err = env.DbConfig.ReadFromEnv(); err != nil {
			return
		}
		if err = env.DbConfig.ReadFromFile(); err != nil {
			return
		}
		env.DbConfig.AddFlags(flagset)

		// AWS Config Setup
		if err = env.AwsConfig.ReadFromEnv(); err != nil {
			return
		}
		if err = env.AwsConfig.ReadFromFile(); err != nil {
			return
		}
		env.AwsConfig.AddFlags(flagset)
	})

	return err
}

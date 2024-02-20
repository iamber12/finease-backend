package environment

import (
	"bitbucket.com/finease/backend/pkg/environment/config"
)

type Environment struct {
	ApplicationConfig *config.ApplicationConfig
}

var Env = Environment{}

func Initialize(config *config.ApplicationConfig) {
	Env.ApplicationConfig = config
}

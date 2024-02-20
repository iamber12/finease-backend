package environment

import (
	"bitbucket.com/finease/backend/pkg/db"
	dbSession "bitbucket.com/finease/backend/pkg/db/db_session"
	"bitbucket.com/finease/backend/pkg/environment/config"
)

type Environment struct {
	ApplicationConfig *config.ApplicationConfig
	Database          Database
}

type Database struct {
	SessionFactory db.SessionFactory
}

var Env = Environment{}

func Initialize(config *config.ApplicationConfig) {
	Env.ApplicationConfig = config
	Env.Database = Database{
		SessionFactory: dbSession.NewSessionFactory(config.DbConfig),
	}
}

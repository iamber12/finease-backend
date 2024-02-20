package migrate

import (
	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/environment/config"
	"context"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Perform DB migrations",
		Long:  "Perform DB migrations",
		Run:   runMigrate,
	}

	if err := config.Setup(cmd.PersistentFlags()); err != nil {
		glog.Fatalf("Unable to setup the application config: '%v'", err.Error())
	}

	return cmd
}

func runMigrate(cmd *cobra.Command, args []string) {
	environment.Initialize(&config.Conf)
	sessionFactory := environment.Env.Database.SessionFactory
	defer sessionFactory.Close()
	connection := sessionFactory.New(context.Background())

	db.Migrate(connection)
}

package main

import (
	"bitbucket.com/finease/backend/cmd/migrate"
	"bitbucket.com/finease/backend/cmd/serve"
	"bitbucket.com/finease/backend/pkg/environment/config"
	"flag"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func main() {
	// This is needed to make `glog` believe that the flags have already been parsed, otherwise
	// every log messages is prefixed by an error message stating the the flags haven't been
	// parsed.
	_ = flag.CommandLine.Parse([]string{})

	// Always log to stderr by default
	if err := flag.Set("logtostderr", "true"); err != nil {
		glog.Infof("Unable to set logtostderr to true")
	}

	rootCmd := &cobra.Command{
		Use:  "finease-backend",
		Long: "finease-backend serves as a service for the backend of the core services associated with Finease",
	}

	if err := config.Setup(rootCmd.PersistentFlags()); err != nil {
		glog.Fatalf("Unable to setup the application config: '%v'", err.Error())
	}

	// All subcommands under root
	serveCmd := serve.NewServeCommand()
	migrateCmd := migrate.NewMigrateCommand()

	rootCmd.AddCommand(serveCmd, migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
	}
}

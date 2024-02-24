package serve

import (
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/environment/config"
	"bitbucket.com/finease/backend/pkg/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"log"
)

func NewServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the finease-backend",
		Long:  "Serve the finease-backend.",
		Run:   runServe,
	}

	return cmd
}

func runServe(cmd *cobra.Command, args []string) {
	environment.Initialize(&config.Conf)

	applicationConfig := environment.Env.ApplicationConfig
	server := gin.New()

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "Not found", "message": "Page not found", "env": applicationConfig.ServerConfig.EnvName})
	})

	server.Use(gin.Recovery())
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})

	glog.Infof("server running in %s mode", applicationConfig.ServerConfig.EnvName)

	routers.SetupRouter(server)

	host, port := applicationConfig.ServerConfig.ListenAddress, applicationConfig.ServerConfig.ListenPort
	err := server.Run(fmt.Sprintf("%s:%d", host, port))
	log.Fatal(err)
}

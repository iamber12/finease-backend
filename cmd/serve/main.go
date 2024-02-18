import (  
	"bitbucket.com/finease/backend/pkg/environment"  
	"bitbucket.com/finease/backend/pkg/routers"  
	"fmt"  
	"github.com/gin-gonic/gin"  
	"github.com/golang/glog"  
	"github.com/spf13/cobra"  
	"log"  
	)  
	  
	func NewServeCommand() *cobra.Command {  
		cmd := &cobra.Command{  
			Use: "serve",  
			Short: "Serve the finease-backend",  
			Long: "Serve the finease-backend.",  
			Run: runServe,  
		}  
		
		return cmd  
	}  
	  
	func runServe(cmd *cobra.Command, args []string) {  
		server := gin.New()  
		
		server.NoRoute(func(c *gin.Context) {  
			c.JSON(404, gin.H{"code": "Not found", "message": "Page not found"})  
		})  
		
		server.Use(gin.Recovery())  
		
		glog.Infof("server running")  
		
		host, port := "localhost", 8000
		err := server.Run(fmt.Sprintf("%s:%d", host, port))  
		log.Fatal(err)  
	}
package login

import (
	"bitbucket.com/finease/backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller(c *gin.Context) {
	response := utils.ResponseRenderer("Welcome to Finease Login!")
	c.AbortWithStatusJSON(http.StatusOK, response)
}

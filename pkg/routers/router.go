package routers

import (
	"bitbucket.com/finease/backend/pkg/routers/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(parentRouter *gin.Engine) {
	v1Router := parentRouter.Group("/v1")

	v1.SetupAuthRouter(v1Router)
}

package routers

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/routers/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(parentRouter *gin.Engine) {
	v1Router := parentRouter.Group("/v1")

	authService := services.NewAuthService(dao.NewSqlUserDao(environment.Env.Database.SessionFactory))

	v1.SetupAuthRouter(v1Router, handlers.NewAuthHandler(authService))
}

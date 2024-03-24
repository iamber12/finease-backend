package routers

import (
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/middlewares"
	v1 "bitbucket.com/finease/backend/pkg/routers/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(parentRouter *gin.Engine) {
	v1Router := parentRouter.Group("/v1")
	dbSessionFactory := environment.Env.Database.SessionFactory

	authService := services.NewAuthService(
		dao.NewSqlUserDao(dbSessionFactory),
	)
	loanProposalService := services.NewLoanProposalService(
		dao.NewSqlLoanProposalDao(dbSessionFactory),
		dao.NewSqlUserDao(dbSessionFactory),
	)
	loanRequestService := services.NewLoanRequestService(
		dao.NewSqlLoanRequestDao(dbSessionFactory),
		dao.NewSqlUserDao(dbSessionFactory),
	)

	jwtAuthzMiddleware := middlewares.IsJwtAuthorized(dao.NewSqlUserDao(dbSessionFactory))

	v1.SetupAuthRouter(v1Router, authService)
	v1.SetupLoanProposalRouter(v1Router, loanProposalService, jwtAuthzMiddleware)
	v1.SetupLoanRequestsRouter(v1Router, loanRequestService, jwtAuthzMiddleware)
}

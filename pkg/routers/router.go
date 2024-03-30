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

	loanRequestDao := dao.NewSqlLoanRequestDao(dbSessionFactory)
	loanProposalDao := dao.NewSqlLoanProposalDao(dbSessionFactory)
	loanAgreementDao := dao.NewSqlLoanAgreementDao(dbSessionFactory)
	userDao := dao.NewSqlUserDao(dbSessionFactory)

	authService := services.NewAuthService(
		dao.NewSqlUserDao(dbSessionFactory),
	)
	loanProposalService := services.NewLoanProposalService(
		loanProposalDao,
		loanRequestDao,
		loanAgreementDao,
		userDao,
	)
	loanRequestService := services.NewLoanRequestService(
		loanRequestDao,
		loanProposalDao,
		loanAgreementDao,
		userDao,
	)
	userService := services.NewUserService(
		userDao,
	)

	jwtAuthzMiddleware := middlewares.IsJwtAuthorized(dao.NewSqlUserDao(dbSessionFactory))

	v1.SetupAuthRouter(v1Router, authService)
	v1.SetupLoanProposalRouter(v1Router, loanProposalService, jwtAuthzMiddleware)
	v1.SetupLoanRequestsRouter(v1Router, loanRequestService, jwtAuthzMiddleware)
	v1.SetupUserRouter(v1Router, userService, jwtAuthzMiddleware)
}

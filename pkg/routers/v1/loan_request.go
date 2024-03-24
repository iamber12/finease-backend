package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupLoanRequestsRouter(parentRouter *gin.RouterGroup, loanRequestsService services.LoanRequest, additionalMiddlewares ...gin.HandlerFunc) {
	authRouter := parentRouter.Group("/loan/request")
	loanRequestsHandler := handlers.NewLoanRequestHandler(loanRequestsService)

	authRouter.Use(additionalMiddlewares...)

	authRouter.POST("/", loanRequestsHandler.Create)
	authRouter.PUT("/:loan_request_uuid", loanRequestsHandler.Update)
	authRouter.DELETE("/:loan_request_uuid", loanRequestsHandler.Delete)
	authRouter.GET("/", loanRequestsHandler.FindByUserId)
	authRouter.GET("/loan-requests", loanRequestsHandler.FindAll)
	authRouter.GET("/:loan_request_uuid", loanRequestsHandler.FindById)
}

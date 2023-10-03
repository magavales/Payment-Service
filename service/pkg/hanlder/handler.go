package handler

import (
	"github.com/gin-gonic/gin"
	"service/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			accounts := v1.Group("/accounts")
			{
				accounts.GET("/:user_id", h.getBalance)
				accounts.POST("/", h.createAccount)
				accounts.PUT("/:operations_id", h.updateBalance)
			}
			history := v1.Group("/history")
			{
				history.GET("/:user_id", h.getHistory)
			}
		}
	}

	return router
}

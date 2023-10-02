package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			operations := v1.Group("/operations")
			{
				operations.PUT("/:operations_id", h.updateBalance)
			}
			accounts := v1.Group("/accounts")
			{
				accounts.GET("/:user_id", h.getBalance)
				accounts.POST("/", h.createAccount)
			}
			history := v1.Group("/history")
			{
				history.GET("/:user_id", h.getHistory)
			}
		}
	}

	return router
}

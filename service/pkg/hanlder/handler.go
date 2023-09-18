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
				operations.POST("/:operations_id", h.createAccount)
				operations.PUT("/:operations_id", h.updateBalance)
			}
			accounts := v1.Group("/accounts")
			{
				accounts.GET("/:user_id", h.getBalance)
			}
		}
	}

	return router
}

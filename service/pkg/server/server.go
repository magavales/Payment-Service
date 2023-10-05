package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) InitServer(port string, router *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return s.httpServer.ListenAndServe()
}

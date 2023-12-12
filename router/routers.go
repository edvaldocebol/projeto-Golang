package router

import (
	"github.com/edvaldocebol/projeto-Golang/handler"
	"github.com/gin-gonic/gin"
)

// InitializeRouter inicializa as rotas do aplicativo
func InitializeRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.DeleteOpeningHandler)
		v1.GET("/opening", handler.UpdateOpeningHandler)
		v1.GET("/opening", handler.ListOpeningHandler)
	}
}

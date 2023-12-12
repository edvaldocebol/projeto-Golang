package router

import "github.com/gin-gonic/gin"

func Inisialize() {
	router := gin.Default()
	InitializeRouter(router)
	router.Run(":8080")
}

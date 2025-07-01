package main

import(
	"LiyerNortorsAIpart/internal/handler"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)
	r.POST("/save", handler.SaveHandler)
	r.GET("/search", handler.AnalyzeHandler)

	user := r.Group("user")
	user.Use(handler.AuthMiddlerWare)
	user.POST("/save", handler.SaveHandler)
	user.GET("/search", handler.AnalyzeHandler)
	r.Run("localhost:8080")
}
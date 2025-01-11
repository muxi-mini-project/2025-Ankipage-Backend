package routes

import (
	"Ankipage/controllers"
	"Ankipage/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 笔记相关路由
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/notes", controllers.GetNotes)
	authorized.GET("/recentnotes", controllers.ListRecentNotes)
	authorized.POST("/notes", controllers.CreateNote)
	authorized.PUT("/notes/:id", controllers.UpdateNote)
	authorized.DELETE("/notes/:id", controllers.DeleteNote)

	return r
}

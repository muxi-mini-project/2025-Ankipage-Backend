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
	authorized.Use(middleware.CORSMiddleware())
	authorized.GET("/getnote/:id", controllers.GetNote)
	authorized.GET("/recentnotes/:userid", controllers.ListRecentNotes)
	authorized.GET("/getallnotes/:userid", controllers.GetNotes)
	authorized.POST("/createnote/:userid", controllers.CreateNote)
	authorized.PUT("/updatenote/:id", controllers.UpdateNote)
	authorized.DELETE("/deletenote/:id", controllers.DeleteNote)

	return r
}

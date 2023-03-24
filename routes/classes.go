package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/middleware"
)

func addClassesRoutes(router *gin.Engine) {
	classes := router.Group("/classes")
	classes.Use(middleware.Auth)
	classes.GET("/", controllers.ListClasses)
	classes.GET("/:id", controllers.RetrieveClass)

	classes.POST("/", controllers.CreateClass)

	classes.DELETE("/:id", controllers.DeleteClass)

	classes.POST("/:id/students/:studentId", controllers.AddStudentToClass)
}

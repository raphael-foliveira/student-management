package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/middleware"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addClassesRoutes(router *gin.Engine, db *gorm.DB) {
	classesRepository := repository.NewClassRepository(db)
	classesController := controllers.NewClassesController(classesRepository)
	classes := router.Group("/classes")
	classes.Use(middleware.Auth)
	classes.GET("/", classesController.ListClasses)
	classes.GET("/:id", classesController.RetrieveClass)

	classes.POST("/", classesController.CreateClass)

	classes.DELETE("/:id", classesController.DeleteClass)

	// classes.POST("/:id/students/:studentId", classesController.AddStudentToClass)
}

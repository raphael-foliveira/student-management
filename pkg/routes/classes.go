package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/middleware"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addClassesRoutes(router *gin.RouterGroup, db *gorm.DB) {
	classesRepository := repository.NewClassRepository(db)
	classesController := controllers.NewClassesController(classesRepository)
	router.Use(middleware.Auth)
	router.GET("/", classesController.ListClasses)
	router.GET("/:id", classesController.RetrieveClass)

	router.POST("/", classesController.CreateClass)

	router.DELETE("/:id", classesController.DeleteClass)

	// router.POST("/:id/students/:studentId", classesController.AddStudentToClass)
}

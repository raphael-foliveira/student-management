package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addTeachersRoutes(router *gin.Engine, db *gorm.DB) {
	teachersRepository := repository.NewTeacherRepository(db)
	teachersController := controllers.NewTeachersController(teachersRepository)

	teachers := router.Group("/teachers")
	teachers.GET("/", teachersController.List)
	teachers.GET("/:id", teachersController.Retrieve)

	teachers.POST("/", teachersController.Create)
	// teachers.PUT("/:id", teachersController.UpdateTeacher)

	teachers.DELETE("/:id", teachersController.Delete)
}

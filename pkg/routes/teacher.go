package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addTeachersRoutes(router *gin.RouterGroup, db *gorm.DB) {
	teachersRepository := repository.NewTeacherRepository(db)
	teachersController := controllers.NewTeachersController(teachersRepository)

	router.GET("/", teachersController.ListTeachers)
	router.GET("/:id", teachersController.RetrieveTeacher)

	router.POST("/", teachersController.CreateTeacher)
	// router.PUT("/:id", teachersController.UpdateTeacher)

	router.DELETE("/:id", teachersController.DeleteTeacher)
}

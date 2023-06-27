package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addStudentsRoutes(router *gin.RouterGroup, db *gorm.DB) {
	studentsRepository := repository.NewStudentRepository(db)
	studentsController := controllers.NewStudentsController(studentsRepository)
	// students.Use(middleware.Auth())

	router.GET("/", studentsController.ListStudents)
	router.GET("/:id", studentsController.RetrieveStudent)

	router.POST("/", studentsController.CreateStudent)
	// router.PUT("/:id", studentsController.UpdateStudent)

	router.DELETE("/:id", studentsController.DeleteStudent)
}

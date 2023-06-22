package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/controllers"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
	"gorm.io/gorm"
)

func addStudentsRoutes(router *gin.Engine, db *gorm.DB) {
	studentsRepository := repository.NewStudentRepository(db)
	studentsController := controllers.NewStudentsController(studentsRepository)
	students := router.Group("/students")
	// students.Use(middleware.Auth())

	students.GET("/", studentsController.ListStudents)
	students.GET("/:id", studentsController.RetrieveStudent)

	students.POST("/", studentsController.CreateStudent)
	// students.PUT("/:id", studentsController.UpdateStudent)

	students.DELETE("/:id", studentsController.DeleteStudent)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/db"
)

func Start() {
	database, err := db.GetDb()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	addStudentsRoutes(router.Group("/students"), database)
	addClassesRoutes(router.Group("/classes"), database)
	addTeachersRoutes(router.Group("/teachers"), database)
	router.Run(":8000")
}

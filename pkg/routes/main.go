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
	addStudentsRoutes(router, database)
	addClassesRoutes(router, database)
	addTeachersRoutes(router, database)

	router.Run(":8000")
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
)

type StudentsController struct {
	repo repository.Repository[models.Student]
}

func NewStudentsController(repo repository.Repository[models.Student]) *StudentsController {
	return &StudentsController{repo: repo}
}

func (sc *StudentsController) List(c *gin.Context) {
	students, err := sc.repo.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, students)
}

func (sc *StudentsController) Retrieve(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	student, err := sc.repo.Find(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (sc *StudentsController) Create(c *gin.Context) {
	var newStudent models.Student
	fmt.Println("creating student")
	if err := (c.BindJSON(&newStudent)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := sc.repo.Create(&newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newStudent)
}

// func (sc *StudentsController) UpdateStudent(c *gin.Context) {
// 	uintId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	student, err := sc.repo.Find(uint(uintId))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	if student.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
// 		return
// 	}
// 	if err := c.BindJSON(&student); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	fmt.Println("about to run Update()")
// 	student.Update(c.Param("id"))
// 	c.JSON(http.StatusOK, student)
// }

func (sc *StudentsController) Delete(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	student, err := sc.repo.Find(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = sc.repo.Delete(student.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListStudents(c *gin.Context) {
	manager := models.Student{}
	var students = manager.All()
	c.JSON(http.StatusOK, students)
}

func RetrieveStudent(c *gin.Context) {
	student := models.Student{}
	student.Find(c.Param("id"))
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	fmt.Println("creating student")
	if err := (c.BindJSON(&newStudent)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newStudent.Create()
	c.JSON(http.StatusCreated, newStudent)
}

func UpdateStudent(c *gin.Context) {
	student := models.Student{}
	student.Find(c.Param("id"))
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	if err := c.BindJSON(&student); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("about to run Update()")
	student.Update(c.Param("id"))
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	student := models.Student{}
	student.Delete(c.Param("id"))
	c.JSON(http.StatusNoContent, nil)
}
